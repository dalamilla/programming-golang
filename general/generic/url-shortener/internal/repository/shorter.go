package repository

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/dalamilla/programming-golang/general/generic/url-shortener/internal/types"
	"go.etcd.io/bbolt"
)

type ShorterRepository struct {
	db *bbolt.DB
}

var bucketName = []byte("URLShorter")

func NewShorterRepository(db *bbolt.DB) (*ShorterRepository, error) {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &ShorterRepository{db: db}, nil
}

func (r *ShorterRepository) Create(originalURL string) (*types.Shorter, error) {
	var shorter types.Shorter

	err := r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)

		id, err := b.NextSequence()

		shorter = types.Shorter{
			ShortURL:    id,
			OriginalURL: originalURL,
		}

		data, err := json.Marshal(shorter)
		if err != nil {
			return err
		}

		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, id)

		return b.Put(key, data)
	})

	if err != nil {
		return nil, err
	}

	return &shorter, nil
}

func (r *ShorterRepository) Get(shortURL uint64) (*types.Shorter, error) {
	var shorter types.Shorter
	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)

		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, shortURL)

		data := b.Get(key)
		if data == nil {
			return fmt.Errorf("url not found")
		}
		return json.Unmarshal(data, &shorter)
	})
	if err != nil {
		return nil, err
	}
	return &shorter, nil
}
