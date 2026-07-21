package database

import (
	"go.etcd.io/bbolt"
	"time"
)

func InitDB(filepath string) (*bbolt.DB, error) {
	db, err := bbolt.Open(filepath, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return db, nil
}
