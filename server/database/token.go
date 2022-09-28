package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"github.com/ahmetb/go-linq/v3"
	"github.com/go-jet/jet/v2/qrm"

	. "github.com/go-jet/jet/v2/postgres"
)

type Token struct {
	model.Tokens
}

var tokenInsertedColumns = Tokens.MutableColumns.Except(Tokens.CreatedAt, Tokens.DeletedAt, Tokens.UpdatedAt)

func TokenInit(UserID int64, Token string, Push *string) model.Tokens {
	return model.Tokens{
		Token:  Token,
		UserID: UserID,
		Push:   Push,
	}
}

func TokenSelect() ProjectionList {
	return ProjectionList{
		Tokens.AllColumns,
	}
}

func TokenFrom() ReadableTable {
	return Tokens
}

func CreateUserToken(db qrm.DB, token *model.Tokens) error {
	stmt := Tokens.INSERT(tokenInsertedColumns).
		MODEL(token).
		RETURNING(Tokens.AllColumns)
	if err := stmt.QueryContext(Ctx, db, token); err != nil {
		return err
	}

	return nil
}

func DeleteTokens(db qrm.DB, ids []int64) ([]model.Tokens, error) {
	stmt := Tokens.UPDATE(Tokens.DeletedAt).
		SET("now()").
		WHERE(Tokens.ID.IN(InInt(ids)...)).
		RETURNING(Tokens.AllColumns)

	var tokens []model.Tokens
	if err := stmt.QueryContext(Ctx, db, &tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}

func TokenGetIDs(tokens []Token) []int64 {
	var IDs []int64
	linq.From(tokens).Select(func(c interface{}) interface{} {
		return c.(Token).ID
	}).ToSlice(&IDs)
	return IDs
}
