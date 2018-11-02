package repositories

import (
	"errors"
	"sync"

	"github.com/cooladdr/xiuxiubeidanci/datamodels"
)


type WordRepository interface {
	Select(where string)datamodels.Word,found bool
	SelectMany(where, group_by, order_by string, page, size int) (results []datamodels.Word)
	InsertOrUpdate(wrod datamodels.Word) (updatedWord datamodels.Word, err error)
	Delete(where string) (deleted bool)
}


func NewWordRepository(db *sql.DB) WordRepository {
	return &wordMysqlRepository{db: db}
}

type wordMysqlRepository struct {
	db *sql.DB
}


func (r *wordMysqlRepository) Select(where string) word datamodels.Word, found bool {

	sql_str="SELECT id,spelling,in_usa,in_uk,w_type FROM word WHERE " + where

	row, err := r.db.QueryRow(sql_str)
	if err != nil {
		found = false
		return
	}
	defer row.Close()

	err = row.Scan(&word.ID, &wrod.spelling, &word.in_usa, &word.in_uk, &word.w_type)
	if err != nil {
		found = false
		return 
	}
	
	found = true
	return 
}

// SelectMany same as Select but returns one or more datamodels.User as a slice.
// If limit <=0 then it returns everything.
func (r *userMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.User) {
	r.Exec(query, func(m datamodels.User) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

// InsertOrUpdate adds or updates a user to the (memory) storage.
//
// Returns the new user and an error if any.
func (r *userMemoryRepository) InsertOrUpdate(user datamodels.User) (datamodels.User, error) {
	id := user.ID

	if id == 0 { // Create new action
		var lastID int64
		// find the biggest ID in order to not have duplications
		// in productions apps you can use a third-party
		// library to generate a UUID as string.
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		user.ID = id

		// map-specific thing
		r.mu.Lock()
		r.source[id] = user
		r.mu.Unlock()

		return user, nil
	}

	// Update action based on the user.ID,
	// here we will allow updating the poster and genre if not empty.
	// Alternatively we could do pure replace instead:
	// r.source[id] = user
	// and comment the code below;
	current, exists := r.Select(func(m datamodels.User) bool {
		return m.ID == id
	})

	if !exists { // ID is not a real one, return an error.
		return datamodels.User{}, errors.New("failed to update a nonexistent user")
	}

	// or comment these and r.source[id] = user for pure replace
	if user.Username != "" {
		current.Username = user.Username
	}

	if user.Firstname != "" {
		current.Firstname = user.Firstname
	}

	// map-specific thing
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()

	return user, nil
}

func (r *userMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.User) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
