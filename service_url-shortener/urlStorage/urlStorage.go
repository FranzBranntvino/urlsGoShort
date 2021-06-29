// Package Url Storage is a simple client structure for storing and reading Url Records for the Url shortener from a DB.
// For convenience wrapper functions to read and write comple Url-Records are provided.
// It does not implement any logic for Data consistency.
// It does not implement any logic for DB consistency.
// It does not implement any logic for Data extraction.
package urlStorage

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Imports:

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	// not standard
	redis "github.com/go-redis/redis/v8"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

type storeClient struct {
	redisClient *redis.Client
}

var (
	storeService = &storeClient{}
	ctx          = context.Background()
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Constants:

// Time when the Url-Record in the DB will expire.
const recordCacheDuration = (24 * time.Hour)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exported:

// Url-Item as the definition of a URL record stored in the URL storage DB.
type UrlItem struct {
	Id         uint64 `json:"id" redis:"id"`
	UrlBase    string `json:"urlBase" redis:"urlBase"`
	UrlCode    string `json:"urlCode" redis:"urlCode"`
	VisitCount uint64 `json:"visits" redis:"visits"`
}

// NewUrlItem constructor for this type, but with default values.
func NewUrlItem() *UrlItem {
	instance := new(UrlItem)

	instance.Id = 0
	instance.UrlBase = "https://www.urlitem.com"
	instance.UrlCode = "abcde:12345"
	instance.VisitCount = 0

	return instance
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Implementation:

// Initialize the DB-client and return a pointer to access the Storage-Service.
func Initialize(storeIp string, storePort string) (*storeClient, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     storeIp + ":" + storePort,
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	storeService.redisClient = redisClient

	return storeService, err
}

// Close the connection to the DB
func CloseStore() error {
	return storeService.redisClient.Close()
}

// Wrapper Function to check the existence of a key in the DB.
func UrlRecordIsExisting(urlRecord UrlItem) bool {
	exists, _ := isExistingKey(urlRecord.UrlCode)
	return (exists != 0)
}

// Wrapper Function to Save a URL Dataset into the DB.
// The short URL from the UrlRecord is used as key.
func UrlRecordWrite(urlRecord UrlItem) error {
	return writeRecord(urlRecord.UrlCode, urlRecord)
}

// Wrapper Function to Load a URL Dataset from the DB, using the short URL as key.
func UrlRecordRead(urlCode string) (UrlItem, error) {
	urlRecord := NewUrlItem()
	err := readRecord(urlCode, urlRecord)
	return *urlRecord, err
}

// Wrapper Function to delete a URL Dataset from the DB, using the short URL as key.
func UrlRecordDelete(urlCode string) error {
	_, err := deleteRecord(urlCode)
	return err
}

////////////////////////////////////////////////////////////

func isExistingKey(key string) (int64, error) {
	return storeService.redisClient.Exists(ctx, key).Result()
}

func writeRecord(key string, record interface{}) error {
	if recordBytes, err := json.Marshal(record); err == nil {
		return storeService.redisClient.Set(ctx, key, recordBytes, recordCacheDuration).Err()
	} else {
		return err
	}
}

func readRecord(key string, record interface{}) error {
	if recordBytesAsString, err := storeService.redisClient.Get(ctx, key).Result(); err == nil {
		recordBytes := []byte(recordBytesAsString)
		return json.Unmarshal(recordBytes, record)
	} else {
		return err
	}
}

func deleteRecord(key string) (int64, error) {
	return storeService.redisClient.Del(ctx, key).Result()
}
