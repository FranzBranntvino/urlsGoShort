package urlStorage

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Imports:

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

var testStoreService = &storeClient{}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Constants:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exported:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Implementation:

func init() {
    _store, _ := Initialize("172.17.0.1","6379")
    testStoreService = _store
}

func TestStoreInit(t *testing.T) {
    assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertRecord(t *testing.T) {
    urlDataset := UrlItem{
        Id:         12685634646217438788,
        UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
        UrlCode:    "6FEC3Yjvghp",
        VisitCount: 0}

    err := UrlRecordWrite(urlDataset)
    assert.Equal(t, err, nil)
}

func TestInsertRecordTwice(t *testing.T) {
    urlDataset := UrlItem{
        Id:         12685634646217438788,
        UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
        UrlCode:    "6FEC3Yjvghp",
        VisitCount: 0}

    err := UrlRecordWrite(urlDataset)
    assert.Equal(t, err, nil)

    err = UrlRecordWrite(urlDataset)
    assert.Equal(t, err, nil)
}

func TestDeleteUrlRecordTwice(t *testing.T) {
    urlDataset := UrlItem{
        Id:         12685634646217438788,
        UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
        UrlCode:    "6FEC3Yjvghp",
        VisitCount: 0}

    err := UrlRecordWrite(urlDataset)
    assert.Equal(t, err, nil)

    err = UrlRecordDelete(urlDataset.UrlCode)
    assert.Equal(t, err, nil)

    err = UrlRecordDelete(urlDataset.UrlCode)
    assert.Equal(t, err, nil)
}

func TestDeleteRecordTwice(t *testing.T) {
    urlDataset := UrlItem{
        Id:         12685634646217438788,
        UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
        UrlCode:    "6FEC3Yjvghp",
        VisitCount: 0}

    err := UrlRecordWrite(urlDataset)
    assert.Equal(t, err, nil)

    var returnCode int64 = -1
	returnCode, err = deleteRecord(urlDataset.UrlCode)
    assert.Equal(t, err, nil)
	assert.Equal(t, returnCode, int64(1))

    returnCode = -1
	returnCode, err = deleteRecord(urlDataset.UrlCode)
    assert.Equal(t, err, nil)
	assert.Equal(t, returnCode, int64(0))
}

func TestInsertRecordAndRetrieveRecord(t *testing.T) {
    urlDataset := UrlItem{
        Id:         12685634646217438788,
        UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
        UrlCode:    "6FEC3Yjvghp",
        VisitCount: 0}

    err := UrlRecordWrite(urlDataset)
    assert.Equal(t, err, nil)

    loadedRecord := UrlItem{}
    loadedRecord, err = UrlRecordRead(urlDataset.UrlCode)
    assert.Equal(t, err, nil)
    assert.Equal(t, loadedRecord.UrlBase, urlDataset.UrlBase)

    urlDataset = UrlItem{
        Id:         1234567890,
        UrlBase:    "www.google.de",
        UrlCode:    "SDOLwlryDhi",
        VisitCount: 0}

    err = UrlRecordWrite(urlDataset)
    assert.Equal(t, err, nil)

    loadedRecord = UrlItem{}
    loadedRecord, err = UrlRecordRead(urlDataset.UrlCode)
    assert.Equal(t, err, nil)
    assert.Equal(t, loadedRecord.UrlBase, urlDataset.UrlBase)
}


func TestInsertRecordAndExistenceOfRecord(t *testing.T) {
    urlDataset := UrlItem{
        Id:         12685634646217438788,
        UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
        UrlCode:    "6FEC3Yjvghp",
        VisitCount: 0}

    err := UrlRecordWrite(urlDataset)
    assert.Equal(t, err, nil)
    assert.Equal(t, UrlRecordIsExisting(urlDataset), true)

    urlDataset = *NewUrlItem()
    urlDataset.UrlCode = "543210uvwxyz"
    assert.Equal(t, UrlRecordIsExisting(urlDataset), false)
}



