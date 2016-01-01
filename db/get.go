package db

import (
	"github.com/boltdb/bolt"
)

func (d *DB) Get(bucket, key string) ([]byte, error) {
	var val []byte
	var getErr error

	d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}

		v := b.Get([]byte(key))

		val = v
		return nil
	})

	return val, getErr
}
