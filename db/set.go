package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

func (d *DB) Set(bucket, key string, val []byte) error {
	log.Debugf("db.set: bucket=%s key=%s data.size=%d", bucket, key, len(val))

	return d.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		if err := b.Put([]byte(key), val); err != nil {
			return err
		}

		return nil
	})
}
