package db

import (
	"time"

	"github.com/boltdb/bolt"
)

type DBConfig struct {
	DBPath string
}

type DB struct {
	db     *bolt.DB
	config *DBConfig
}

func NewDB(config *DBConfig) (*DB, error) {
	bdb, err := bolt.Open(
		config.DBPath,
		0600,
		&bolt.Options{Timeout: 1 * time.Second},
	)
	if err != nil {
		return nil, err
	}

	return &DB{
		config: config,
		db:     bdb,
	}, nil
}
