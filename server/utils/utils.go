package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/openlyinc/pointy"
	"log"

	"time"

	"net/http"

)

//const ApiUrl = "http://10.0.2.2:8080"

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{} {"Success" : status, "ApiMessage" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	//pp.Print(data)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(data)
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func RemoveDuplicateValues(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	var list []int64

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Contains(str int64, s []int64) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func CompareSlice(old []int64, new []int64) ([]int64, []int64) {
	var forAdd []int64
	var forDel []int64

	for _, element := range new {
		if !Contains(element, old) {
			forAdd = append(forAdd, element)
		}
	}

	for _, element := range old {
		if !Contains(element, new) {
			forDel = append(forDel, element)
		}
	}

	return forAdd, forDel
}

func ToInt(myBool bool) int {
	if myBool {
		return 1
	}
	return 0
}

func TimeMax() time.Time {
	return time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC)
}

//func GetUrlFromId(id string) string {
//	return fmt.Sprintf("%s/file/download/%s", ApiUrl, id)
//}

func StringEmptyToNil(a *string) *string {
	if StringValue(a) == "" {
		return nil
	}
	return a
}

func IntEmptyToNil(a *int64) *int64 {
	if IntValue(a) == 0 {
		return nil
	}
	return a
}

func StringValue(data *string) string {
	return pointy.StringValue(data, "")
}

func IntValue(data *int64) int64 {
	return pointy.Int64Value(data, 0)
}

func SendPush(title string, body string, userIDs []int64) {
	values := map[string]interface{}{"title": title, "body": body, "userIDs": userIDs}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://"+ Settings.PushServer +"/push/send", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}