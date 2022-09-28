package database

import (
	"awesomeProject/crypto"
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"awesomeProject/utils"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

//type FileType int64

const (
	FileTypeFile int64 = iota + 1
	FileTypeImage
)

type File struct {
	model.Files
}

var fileInsertedColumns = Files.MutableColumns.Except(Files.CreatedAt, Files.DeletedAt, Files.UpdatedAt)

func FileInit(CreatorID int64, Name string, Type int64, Hash string, Size int64) model.Files {
	return model.Files{
		CreatorID: CreatorID,
		Name:      Name,
		Type:      Type,
		Hash:      Hash,
		Size:      Size,
		Token:     crypto.NewToken(),
	}
}

func FileSelect() ProjectionList {
	return ProjectionList{
		Files.AllColumns,
	}
}

func FileFrom() ReadableTable {
	return Files
}

func InsertFile(db qrm.DB, file *model.Files) error {
	stmt := Files.INSERT(fileInsertedColumns).
		MODEL(file).RETURNING(Files.AllColumns)

	if err := stmt.QueryContext(Ctx, db, file); err != nil {
		return err
	}

	return nil
}

func GetFileUrl(fileID uint64) string {
	return fmt.Sprintf("https://%s:27991/file/download/%d", utils.Settings.Ip, fileID)
}
