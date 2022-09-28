package internal

import (
	"awesomeProject/database"
	"awesomeProject/gen/postgres/public/model"
	"awesomeProject/utils"
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"time"
)

var count = 1

func printLog(log model.Logs) {
	printString := "[" + strconv.Itoa(count) + "] "

	printString += time.Now().Format("02.01.06 15:04:05") + " "

	switch log.Type {
	case database.LogTypeInfo:
		printString += "[LOG] "
	case database.LogTypeWarning:
		printString += color.YellowString("[WARNING] ")
	case database.LogTypeError:
		printString += color.RedString("[ERROR] ")
	default:

	}

	if log.Text != "" {
		printString += fmt.Sprint(log.Text, " ")
	}
	if log.RequestURI != nil {
		printString += fmt.Sprint("RequestURI:", utils.StringValue(log.RequestURI), " ")
	}
	if log.RemoteAddr != nil {
		printString += fmt.Sprint("RemoteAddr:", utils.StringValue(log.RemoteAddr), " ")
	}
	if log.ErrorKey != nil {
		printString += fmt.Sprint("ErrorKey:", utils.IntValue(log.ErrorKey), " ")
	}

	fmt.Fprintf(color.Output, printString+"\n")
	err := database.LogInsert(database.JetDB, &log)
	if err != nil {
		fmt.Println(err.Error())
	}

	count += 1
}

func LogCreate(Text string) {

	log := database.LogInit(database.LogTypeInfo, Text, nil, nil, nil)
	printLog(log)
}

func LogRequest(RemoteAddr string, RequestURI string) {
	log := database.LogInit(database.LogTypeInfo, "request", &RemoteAddr, &RequestURI, nil)
	printLog(log)
}

func LogError(Text string, key int64) map[string]interface{} {
	log := database.LogInit(database.LogTypeError, Text, nil, nil, &key)
	printLog(log)
	return map[string]interface{}{"Error": utils.ErrorInit(0, Text, key)}
}

func IfLogError(err error, key int64) {
	if err != nil {
		LogError(err.Error(), key)
	}
}
