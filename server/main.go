package main

import (
	"awesomeProject/app"
	"awesomeProject/controllers"
	"awesomeProject/controllers/main_server"
	"awesomeProject/database"
	"awesomeProject/internal"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	database.ConnectDatabase()

	err := os.MkdirAll("files", os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create the file for writing. Check your write access privilege")
		return
	}

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена
	router.HandleFunc("/user/reg",
		controllers.Reg).Methods("POST")
	router.HandleFunc("/user/search",
		controllers.Search).Methods("POST")
	router.HandleFunc("/user/get",
		controllers.UserGet).Methods("POST")
	router.HandleFunc("/user/edit",
		controllers.EditUser).Methods("POST")
	router.HandleFunc("/user/login",
		controllers.Login).Methods("POST")
	router.HandleFunc("/user/logout",
		controllers.Logout).Methods("POST")
	router.HandleFunc("/user/setOnline",
		controllers.SetOnline).Methods("POST")

	router.HandleFunc("/message/send",
		controllers.Send).Methods("POST")
	router.HandleFunc("/message/del",
		controllers.DeleteMessagesByIds).Methods("POST")
	router.HandleFunc("/message/edit",
		controllers.EditMessage).Methods("POST")
	router.HandleFunc("/message/read",
		controllers.ReadMessage).Methods("POST")

	router.HandleFunc("/chat/create",
		controllers.ChatCreate).Methods("POST")
	router.HandleFunc("/chat/my",
		controllers.GetMyChats).Methods("POST")
	router.HandleFunc("/chat/get",
		controllers.GetChat).Methods("POST")
	router.HandleFunc("/chat/getMessages",
		controllers.GetMessages).Methods("POST")
	router.HandleFunc("/chat/edit",
		controllers.EditChat).Methods("POST")
	router.HandleFunc("/chat/kick", // Удалить участников беседы по ид участников
		controllers.ParticipantDel).Methods("POST")
	router.HandleFunc("/chat/invite",
		controllers.ParticipantAdd).Methods("POST")
	router.HandleFunc("/chat/getParticipants",
		controllers.GetParticipants).Methods("POST")
	router.HandleFunc("/chat/del",
		controllers.DeleteChats).Methods("POST")
	router.HandleFunc("/chat/setActivity",
		controllers.SetActivity).Methods("POST")

	router.HandleFunc("/file/upload",
		controllers.UploadFile).Methods("POST")
	router.HandleFunc("/file/download/{token}/{name}",
		controllers.DownloadFile).Methods("GET")
	router.HandleFunc("/file/download/{token}",
		controllers.DownloadFile).Methods("GET")

	router.HandleFunc("/misc/info",
		controllers.Info).Methods("POST")

	router.HandleFunc("/admin/userAdd",
		controllers.UserAdd).Methods("POST")

	router.HandleFunc("/sync/start",
		controllers.SyncStart).Methods("POST")

	router.HandleFunc("/database/getBranches",
		controllers.GetBranches).Methods("POST")
	router.HandleFunc("/database/getStaffs",
		controllers.GetStaffs).Methods("POST")

	// MAIN SERVER
	router.HandleFunc("/push/add",
		controllers.PushAdd).Methods("POST")
	router.HandleFunc("/push/send",
		controllers.PushSend).Methods("POST")
	// получает список серверов
	router.HandleFunc("/server/get",
		main_server.ServersGet).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "27991"
	}

	internal.LogCreate("http server started")

	go func() {
		log.Fatal(http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", router))
	}()

	internal.SocketServerStart()

	select {}
}
