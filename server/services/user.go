package services

import (
	"awesomeProject/api"
	"awesomeProject/crypto"
	"awesomeProject/database"
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"awesomeProject/internal"
	"strings"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/openlyinc/pointy"

	"awesomeProject/utils"
)

func Reg(reg api.Reg) map[string]interface{} {
	if !utils.Settings.Registration {
		internal.LogError("Регистрация отключена", 910673)
		return internal.LogError("Регистрация отключена", 910673)
	}
	err := reg.Validate()
	if err != nil {
		return internal.LogError(err.Error(), 106146)
	}

	hashPass := crypto.PasswordHash(reg.Pass, crypto.RandomBytes256())
	user := database.UserInit(reg.Login, &hashPass, reg.FirstName, reg.LastName, nil, nil, nil)

	err = database.InsertUser(database.JetDB, &user)
	if err != nil {
		return internal.LogError(err.Error(), 225482)
	}

	res := utils.Message(true, "")
	res["User"] = user
	return res
}

func Login(login api.Login) map[string]interface{} {
	if ok, res := login.Validate(); !ok {
		return res
	}
	login.Login = strings.ReplaceAll(strings.ToLower(login.Login), " ", "")

	var user database.User
	err := SELECT(database.UserSelect()).
		FROM(database.UserFrom()).
		WHERE(Users.Login.EQ(String(login.Login))).
		LIMIT(1).Query(database.JetDB, &user)
	if err != nil && err != qrm.ErrNoRows {
		return internal.LogError(err.Error(), 871497)
	}
	if user.ID == 0 {
		return internal.LogError("Неверный логин или пароль", 812293)
	}

	if user.HashPass == nil { // у пользователя не установлен пароль
		if ok := utils.WindowsLogin(login.Login, login.Pass); !ok {
			return internal.LogError("Неверный пароль windows", 824195)
		}
		hashPass := crypto.PasswordHash(login.Pass, crypto.RandomBytes256())
		user.HashPass = &hashPass

		err = Users.
			UPDATE(Users.HashPass).
			MODEL(&user).
			WHERE(Users.ID.EQ(Int(user.ID))).
			RETURNING(Users.AllColumns).
			Query(database.JetDB, &user)
		if err != nil {
			return internal.LogError(err.Error(), 835478)
		}
	}

	if ok := crypto.PasswordCheck(login.Pass, pointy.StringValue(user.HashPass, "")); !ok {
		return internal.LogError("Неверный пароль", 139657)
	}

	token := database.TokenInit(user.ID, crypto.NewToken(), login.Push)
	tx, _ := database.JetDB.Begin()
	err = database.CreateUserToken(tx, &token)
	if err != nil {
		return internal.LogError(err.Error(), 941846)
	}

	tx.Commit()

	resp := utils.Message(true, "")
	resp["Token"] = token
	return resp
}

func Logout(byId api.ById) map[string]interface{} {
	stmt := SELECT(database.TokenSelect()).
		FROM(database.TokenFrom()).
		WHERE(Tokens.UserID.EQ(Int(byId.SenderId)).
			AND(database.NotDeleted(Tokens.DeletedAt)))

	var tokens []database.Token
	err := stmt.Query(database.JetDB, &tokens)
	if err != nil {
		return internal.LogError(err.Error(), 484282)
	}

	_, err = database.DeleteTokens(database.JetDB, database.TokenGetIDs(tokens))
	if err != nil {
		return internal.LogError(err.Error(), 561966)
	}

	internal.SocketSend("logout", []int64{byId.SenderId}, "")

	resp := utils.Message(true, "")
	return resp
}

func UserEdit(userEdit api.UserEdit) map[string]interface{} {

	var user database.User
	err := SELECT(
		database.UserSelect(),
	).FROM(
		database.UserFrom(),
	).WHERE(
		Users.ID.EQ(Int(userEdit.SenderId)),
	).Query(database.JetDB, &user)
	if err != nil {
		return internal.LogError(err.Error(), 107466)
	}

	if userEdit.FirstName != nil {
		user.FirstName = utils.StringValue(userEdit.FirstName)
	}
	if userEdit.MiddleName != nil {
		user.MiddleName = utils.StringEmptyToNil(userEdit.MiddleName)
	}
	if userEdit.LastName != nil {
		user.LastName = utils.StringValue(userEdit.LastName)
	}
	if userEdit.AvatarID != nil {
		user.AvatarID = utils.IntEmptyToNil(userEdit.AvatarID)
	}

	err = database.UpdateUser(database.JetDB, &user.Users)
	if err != nil {
		return internal.LogError(err.Error(), 490369)
	}

	resp := utils.Message(true, "")
	resp["User"] = user
	return resp
}

func GetUser(getById api.ById) map[string]interface{} {

	stmt := SELECT(database.UserSelect()).
		FROM(database.UserFrom()).
		WHERE(Users.ID.EQ(Int(getById.Id)).
			AND(Users.DeletedAt.GT_EQ(TimestampzT(time.Now()))),
		)

	var user database.User
	err := stmt.Query(database.JetDB, &user)
	if err != nil {
		return internal.LogError(err.Error(), 293563)
	}

	resp := utils.Message(true, "")
	resp["User"] = user
	return resp
}

func SearchUsers(searchUser api.SearchUsers) map[string]interface{} {
	if ok, res := searchUser.Validate(); !ok {
		return res
	}
	search := "%" + strings.ToLower(searchUser.SearchString) + "%"
	//var searchArray []string = strings.Split(searchUser.SearchString, " ")
	//var users []UserGorm
	//if err := database.Db.Where("LOWER(first_name) LIKE LOWER(?) OR LOWER(last_name) LIKE LOWER(?)", searchUser.SearchString, searchUser.SearchString).Find(&users).Error; err != nil {
	//	return u.ErrorInit(u.ErrorCodeDatabase, err.Error(), 621622).ToResponse()
	//}

	stmt := SELECT(database.UserSelect()).
		FROM(database.UserFrom()).
		WHERE(LOWER(Users.FirstName).
			LIKE(String(search)).
			OR(LOWER(Users.LastName).
				LIKE(String(search))))

	var users []database.User
	err := stmt.Query(database.JetDB, &users)
	if err != nil {
		return internal.LogError(err.Error(), 195385)
	}

	resp := utils.Message(true, "")
	resp["Users"] = users
	return resp
}

func SetOnline(online model.Onlines) map[string]interface{} {

	err := database.InsertOnline(database.JetDB, &online)
	if err != nil {
		return internal.LogError(err.Error(), 359835)
	}

	resp := utils.Message(true, "")
	resp["Online"] = online
	return resp
}
