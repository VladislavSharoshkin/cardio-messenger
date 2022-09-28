package app

import (
	"awesomeProject/database"
	. "awesomeProject/gen/postgres/public/table"
	"awesomeProject/internal"
	"awesomeProject/utils"
	"context"
	. "github.com/go-jet/jet/v2/postgres"
	"strings"

	"net/http"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log := queries.LogInit(nil, &r.RemoteAddr, &r.RequestURI)
		//err := queries.LogInsert(database.JetDB, &log)
		//if err != nil {
		//	utils.Print(err.Error(), utils.PrintTypeWarning)
		//}
		internal.LogRequest(r.RemoteAddr, r.RequestURI)

		notAuth := []string{"/misc/info", "/user/reg", "/user/login", "/file/add", "/sync/start", "/push/add", "/push/send"} //Список эндпоинтов, для которых не требуется авторизация
		requestPath := r.URL.Path                                                                                            //текущий путь запроса

		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}
		if strings.Contains(requestPath, "/file/download/") {
			next.ServeHTTP(w, r)
			return
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Получение токена

		if tokenHeader == "" { //Токен отсутствует, возвращаем  403 http-код Unauthorized
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			utils.Respond(w, response)
			return
		}

		//splitted := strings.Split(tokenHeader, " ") //Токен обычно поставляется в формате `Bearer {token-body}`, мы проверяем, соответствует ли полученный токен этому требованию
		//if len(splitted) != 2 {
		//	response = utils.Message(false, "Invalid/Malformed auth token")
		//	w.WriteHeader(http.StatusForbidden)
		//	w.Header().Add("Content-Type", "application/json; charset=utf-8")
		//	utils.Respond(w, response)
		//	return
		//}
		//
		//tokenPart := splitted[1] //Получаем вторую часть токена
		//tk := &models.Token{}
		//
		//token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		//	return []byte(os.Getenv("token_password")), nil
		//})

		stmt := SELECT(Tokens.AllColumns).
			FROM(Tokens).
			WHERE(Tokens.Token.EQ(String(tokenHeader)).
				AND(database.NotDeleted(Tokens.DeletedAt))).
			LIMIT(1)

		var token database.Token
		err := stmt.Query(database.JetDB, &token)
		if err != nil {
			internal.SocketSend("logout", []int64{token.UserID}, "")
			response = internal.LogError("Неверный токен", 513368)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			utils.Respond(w, response)
			return
		}

		//if !token.Valid { //токен недействителен, возможно, не подписан на этом сервере
		//	response = utils.Message(false, "Token is not valid.")
		//	w.WriteHeader(http.StatusForbidden)
		//	w.Header().Add("Content-Type", "application/json; charset=utf-8")
		//	utils.Respond(w, response)
		//	return
		//}

		//Всё прошло хорошо, продолжаем выполнение запроса

		ctx := context.WithValue(r.Context(), "user", token.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //передать управление следующему обработчику!
	})
}
