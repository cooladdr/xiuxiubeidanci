package repositories

import (
	"fmt"

	"database/sql"

	"github.com/cooladdr/xiuxiubeidanci/datamodels"
)


type WordRepository interface {
	Find(word string)datamodels.Word
	Select(where string)(datamodels.Word, bool)
	SelectMany(where, group_by, order_by string, page, size int) []datamodels.Word
	InsertOrUpdate(wrod datamodels.Word) ( datamodels.Word,  error)
	Delete(where string) bool
}


func NewWordRepository(db *sql.DB) WordRepository {
	return &wordMysqlRepository{db: db}
}

type wordMysqlRepository struct {
	db *sql.DB
}

func (r *wordMysqlRepository) Find(word string) (w datamodels.Word ){
	sql := "SELECT id,spelling,in_usa,in_uk,w_type FROM word WHERE spelling='" + word + "'"
	row := r.db.QueryRow(sql)

	fmt.Printf("%t",row)


	err := row.Scan(&w.ID, &w.Spelling, &w.InUSA, &w.InUK, &w.WType)
	if err != nil {
		return 
	}
	
	return 
}


func (r *wordMysqlRepository) Select(where string) (w datamodels.Word, found bool) {

	sql_str :="SELECT id,spelling,in_usa,in_uk,w_type FROM word WHERE " + where

	row := r.db.QueryRow(sql_str)

	err := row.Scan(&w.ID, &w.Spelling, &w.InUSA, &w.InUK, &w.WType)
	if err != nil {
		found = false
		return 
	}
	
	found = true
	return 
}


func (r *wordMysqlRepository) SelectMany(where, group_by, order_by string, page, size int) (wordlist []datamodels.Word) {
	return nil
}



func (r *wordMysqlRepository) InsertOrUpdate(wrod datamodels.Word) (w datamodels.Word, success error) {
	return w,nil
}


func (r *wordMysqlRepository) Delete(where string) (success bool) {
	return true
}

