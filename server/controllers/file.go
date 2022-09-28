package controllers

import (
	"awesomeProject/database"
	"awesomeProject/services"
	u "awesomeProject/utils"
	"crypto/sha256"
	"encoding/hex"

	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var UploadFile = func(w http.ResponseWriter, r *http.Request) {

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	filesForm := formdata.File["multipleFiles"] // grab the filenames
	//var files []models.File
	for _, fileForm := range filesForm { // loop through the files one by one
		fileFormStream, err := fileForm.Open()
		defer fileFormStream.Close()
		if err != nil {
			u.Respond(w, u.Message(false, err.Error()))
			return
		}

		// CreateChat a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("files", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(fileFormStream)
		if err != nil {
			fmt.Println(err)
		}

		h := sha256.New()
		h.Write(fileBytes)
		var hash = hex.EncodeToString(h.Sum(nil))

		// write this byte array to our temporary file
		_, err = tempFile.Write(fileBytes)
		if err != nil {
			u.Respond(w, u.Message(false, err.Error()))
			return
		}
		tempFile.Close()
		e := os.Rename(tempFile.Name(), "files/"+hash)
		if e != nil {
			log.Fatal(e)
		}
		var contentType = http.DetectContentType(fileBytes)

		//fileExtension := filepath.Ext(fileForm.Filename)

		var fileType int64
		switch {
		case strings.Contains(contentType, "image"):
			fileType = database.FileTypeImage
		default:
			fileType = database.FileTypeFile
		}

		var file = database.FileInit(r.Context().Value("user").(int64), fileForm.Filename, fileType, hash, fileForm.Size)
		resp := services.FileUpload(file)
		u.Respond(w, resp)
		return
	}

	//resp := models.CreateFiles(files)
	//u.Respond(w, resp)
}

var DownloadFile = func(w http.ResponseWriter, r *http.Request) {
	//parts := strings.Split(r.URL.Path, "/")
	//fileId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	//if err != nil {
	//	u.Respond(w, u.Message(false, err.Error()))
	//	return
	//}
	fileToken := mux.Vars(r)["token"]
	_, fileBytes := services.FileDownload(fileToken)
	//if !ok { // if the map contains the image name
	//	u.Respond(w, resp)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}
