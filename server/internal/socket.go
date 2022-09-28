package internal

import (
	"awesomeProject/database"
	. "awesomeProject/gen/postgres/public/table"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"time"
)

var SocketServer *socketio.Server
var onlineObserve map[string][]socketio.Conn
var connectedUserIds = make(map[string]int64)

func init() {

}

func SocketServerStart() {
	//onlineObserve = make(map[string][]string)
	SocketServer = socketio.NewServer(nil)

	SocketServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		LogCreate(fmt.Sprint("socket connected ", s.ID()))
		return nil
	})

	SocketServer.OnEvent("/", "token", func(s socketio.Conn, msg string) {

		stmt := SELECT(database.TokenSelect()).
			FROM(database.TokenFrom()).
			WHERE(Tokens.Token.EQ(String(msg)).
				AND(Tokens.DeletedAt.GT_EQ(TimestampzT(time.Now()))))
		var token database.Token
		if err := stmt.Query(database.JetDB, &token); err != nil {
			return
		}
		online := database.OnlineInit(token.UserID, database.OnlineTypeOnline)
		err := database.InsertOnline(database.JetDB, &online)
		if err != nil {
			fmt.Println(err.Error())
		}
		LogCreate(fmt.Sprint("socket token ", msg))
		s.Join(fmt.Sprint(token.UserID))
		connectedUserIds[s.ID()] = token.UserID
		//OnlineAdd(OnlineInit(token.UserID, OnlineTypeOnline))
	})

	SocketServer.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	SocketServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
		//var userId = connectedUserIds[s.ID()]

		// OnlineAdd(OnlineInit(userId, OnlineTypeOffline))
	})

	go SocketServer.Serve()
	defer SocketServer.Close()
	http.Handle("/socket.io/", SocketServer)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	LogCreate("socket server started")
	log.Fatal(http.ListenAndServeTLS(":"+"27992", "cert.pem", "key.pem", nil))
}

func SocketSend(event string, userIds []int64, data interface{}) {
	for _, element := range userIds {
		SocketServer.BroadcastToRoom("/", fmt.Sprint(element), event, data)
	}
}

func SocketSendUpdate(userIds []int64) {
	SocketSend("update", userIds, "")
}
