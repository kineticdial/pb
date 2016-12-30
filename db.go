package pb

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type DBClient struct {
	file  string
	perms int
}

func NewDBClient() *DBClient {
	return &DBClient{"./.pb", 0644}
}

func (c *DBClient) PutObject(b []byte, k []byte, v []byte) {
	db := c.db()
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(b)
		if err != nil {
			return err
		}

		err = bucket.Put(k, v)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (c *DBClient) GetObject(b []byte, k []byte) []byte {
	var v []byte
	db := c.db()
	defer db.Close()

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", b)
		}

		v = bucket.Get(k)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return v
}

func (c *DBClient) db() *bolt.DB {
	db, err := bolt.Open("./.pb", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
