package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

/* Example Usage

db, err := OpenDB("./bolt.db")

if err != nil {
	log.Fatal(err)
}

defer db.Close()

err = PutData(db, "users", "name", "Nithin")

if err != nil {
	log.Fatal(err)
}

v, err := GetData(db, "users", "name")

fmt.Println(v)

if err != nil {
	log.Fatal(err)
}

*/

//OpenDB opens the database in the given path
func OpenDB(p string) (*bolt.DB, error) {
	return bolt.Open(p, 0644, nil)
}

//KeyValueGenerator generates byte slices key-value pairs
func KeyValueGenerator(k, v string) ([]byte, []byte) {
	return []byte(k), []byte(v)
}

//PutData is used to update database it creates new bucket if the bucket doesn't exist
func PutData(db *bolt.DB, bucketName, key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}
		return nil
	})
}

//GetData return the value for the given key from the bucket
//if there is an error the value returned is empty string
func GetData(db *bolt.DB, bucketName, key string) (string, error) {

	var value string
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket: " + bucketName + " not found")
		}
		value = string(bucket.Get([]byte(key)))
		return nil
	})
	if err != nil {
		return "", err
	}
	{
		return value, nil
	}

}
