package bolt

import (
	"errors"
	"strconv"

	"github.com/HAtherlolz/tg-pocket-gobot/pkg/repository"
	"github.com/boltdb/bolt"
)

type TokenRepository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

// Save Dict data in boltBD

func (r *TokenRepository) Save(chatID int64, token string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

// Get Token form boltDB

func (r *TokenRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var token string

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToBytes(chatID))
		token = string(data)
		return nil
	})

	if err != nil {
		return "", err
	}

	if token == "" {
		return "", errors.New("token not found")
	}

	return token, nil
}

// Convert Int64 to string and String to slice of bytes

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
