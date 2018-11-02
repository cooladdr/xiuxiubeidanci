package repositories

import (
	"errors"
	"database/sql"

	"github.com/cooladdr/xiuxiubeidanci/datamodels"
)


type WordRepository interface {
	Select(where string)(w datamodels.Word,found bool)
	SelectMany(where, group_by, order_by string, page, size int) (results []datamodels.Word)
	InsertOrUpdate(wrod datamodels.Word) (updatedWord datamodels.Word, err error)
	Delete(where string) (deleted bool)
}


func NewWordRepository(db *sql.DB) *WordRepository {
	return &wordMysqlRepository{db: db}
}

type wordMysqlRepository struct {
	db *sql.DB
}


func (r *wordMysqlRepository) Select(where string) (w datamodels.Word, found bool) {

	sql_str :="SELECT id,spelling,in_usa,in_uk,w_type FROM word WHERE " + where

	row, err := r.db.QueryRow(sql_str)
	if err != nil {
		found = false
		return
	}
	defer row.Close()

	err = row.Scan(&w.ID, &w.spelling, &w.in_usa, &w.in_uk, &w.w_type)
	if err != nil {
		found = false
		return 
	}
	
	found = true
	return 
}
