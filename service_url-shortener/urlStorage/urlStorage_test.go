package urlStorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Constants:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exported:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Implementation:

func init() {
	Initialize("172.17.0.1", "6379")
}

func TestStoreInit(t *testing.T) {
	// fail init attempt
	assert.NotEqual(t, nil, Initialize("0.0.0.0", "0815"))

	// real init attempt
	assert.Equal(t, nil, Initialize("172.17.0.1", "6379"))
	assert.True(t, storeService.redisClient != nil)
}

func TestInsertRecord(t *testing.T) {
	urlDataset := UrlItem{
		Id:         12685634646217438788,
		UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
		UrlCode:    "6FEC3Yjvghp",
		VisitCount: 0}

	err := UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)
}

func TestInsertRecordTwice(t *testing.T) {
	urlDataset := UrlItem{
		Id:         12685634646217438788,
		UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
		UrlCode:    "6FEC3Yjvghp",
		VisitCount: 0}

	err := UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)

	err = UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)
}

func TestDeleteUrlRecordTwice(t *testing.T) {
	urlDataset := UrlItem{
		Id:         12685634646217438788,
		UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
		UrlCode:    "6FEC3Yjvghp",
		VisitCount: 0}

	err := UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)

	err = UrlRecordDelete(urlDataset.UrlCode)
	assert.Equal(t, nil, err)

	err = UrlRecordDelete(urlDataset.UrlCode)
	assert.Equal(t, nil, err)
}

func TestInsertRecordAndRetrieveRecord(t *testing.T) {
	urlDataset := UrlItem{
		Id:         12685634646217438788,
		UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
		UrlCode:    "6FEC3Yjvghp",
		VisitCount: 0}

	err := UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)

	loadedRecord := UrlItem{}
	loadedRecord, err = UrlRecordRead(urlDataset.UrlCode)
	assert.Equal(t, nil, err)
	assert.Equal(t, loadedRecord.UrlBase, urlDataset.UrlBase)

	urlDataset = UrlItem{
		Id:         1234567890,
		UrlBase:    "www.google.de",
		UrlCode:    "SDOLwlryDhi",
		VisitCount: 0}

	err = UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)

	loadedRecord = UrlItem{}
	loadedRecord, err = UrlRecordRead(urlDataset.UrlCode)
	assert.Equal(t, nil, err)
	assert.Equal(t, loadedRecord.UrlBase, urlDataset.UrlBase)

	urlDataset = UrlItem{
		Id:         13707503500505479484,
		UrlBase:    "https://godbolt.org/#z:OYLghAFBqd5QCxAYwPYBMCmBRdBLAF1QCcAaPECAM1QDsCBlZAQwBtMQBGAFlJvoCqAZ0wAFAB4gA5AAYppAFZdSrZrVDIApACYAQjt2kR7ZATx1KmWugDCqVgFcAtrWVX0AGTy1MAOWcARpjEIADMAJykAA6oQoTmtHaOLsoxcWZ0Xj7%2BTkEhEUaYJhm0DATMxARJzq6chcUJZRUEWX6BwWGRQuWV1Sl13c2tOXmdAJRGqA7EyBxSOqHeyI5YANSaoTYVxMwAnhvYmjIAggtLK5jrm%2BbdxJjMTgdHp9qLtMsOaxs2yLfewE8Ts9ngB6ABUqzKxAcplWWCo3nidCEqzBIOeBF2UUw8NWtxhBFWAH0hKgnJgicwHEQicArMFmAQcUS0F8AOz6E6rbmrAKoeyrKLEVDYyq7GQbTnHHm8/msQXC0WYziS54yvkCoUi4KY7Sqrk8jXyrVK3ahfXSw1yhXasXcC3q60mnW7ACsFs0bIAIniyRSqTS6T4dkz0CyMJgPScQTHVgBxekhnGrNh4ZgomjEVYAFSxkZODji6lWokVOrwmCEACUivcRFcfaTyZTqahaYnGczWZHQlLnoMzMgS2XKhXq7X05cAG5sByYXQOKhUYIN9YctU8qhsESkDfcresHd71YHo8G/fbzC788ny/Xy0Xw9X4HeqPHJlOKKqJnfZbplHZoCxxTqgeDoKsqgAF67HYn4VJgUIEhAaC0N0OY6AAbKsM6sGMa5SrGPIDngQ7Zths7zouy7EA6PJ4FQqwQIx3ToCAIDkk4yCfhAdzeEyxBCpgBAsumBC/ggFSrCh3RggcECYThc4LkuwQTKsvH0MEgnCSw3TiZJ0kELJoTYPJ2gYThalxJBmCoFQEDZmMeFgGAGw%2BjIeF4Z6UoyjKLFsWg1JXDY3zrNo2irL4ADyZGYAAjg4bBhRF3yhf5IDuKwtG%2BeRjiUSpWZubl2Vrj6RT1t5x45dy6WBYSqWbMlqzxYl8o6ClmxpQQrEZdYWW9senpei%2Bw1AmNxzEUOhmYOIQq%2Bs2AZtkGDKhuGawANalraZiVjW7CTpwq6VTep7Pidd7Hqd94yldl0XedT73kNb79uUg5SciTKzVmTb%2Bq27bBp2Ybdqsm0jjt477SIHU%2BsdD63o9d2Iw9Z7wwQ0JnfDt0o5jz0DeNzx8asTjMN4EBeeuN61VM9WdY11m2fZin5dRXl0zYeLdWxmVvjKUEwWSUTwYhpgQGD21jntdaYJwYwlfzsFC3cIsEGLW1KpLE7Q3L%2BOnDedwENMtCrBKuvPeNUgTKw0iuvIriyPIqDSCFegGL60yzGFoScPIBDSHITmkOtIDcDIAB0oQyAAHNobJR2y3A8JwrrhOZKjSNw8hOFwMgyKQ9tyKQTtSPIQggHnfsOxMcCwDAiAgDTUTUuQlBoJ%2BeDsCEdQ4oQJDKPwjAsOwXC8APwhiJIDukAA7jsUTSD7Vs23b/uO9IUXUk3hJ2asxxVgAsqswDIEO4Rh9oYeHRAtB0JgeEQIrHcrgssu%2B6vgfB9oUdhzIoRsqEUdvZsjZK6NkudU7pykJnUg2dOBn1zpwaOoC2QYVCJHGQKD86ryLtIUu5dSCVwDqQGu9cZqYGQK2MgFAIAVGAEIUQagigMAQKgae9sfakDblEDujIEgMJ8KwZhrCC7yC4U/EIwBEGhE4YLcRG9GAsLYdgshyBjjEDobg0gKiyj4HtvIAeTA2AcB4HwOgBBx4SGwUoOoqh1AoAMAYFQeAAhl0gBMEUJQy4l0mB7YxRhureH4UwxR7D5Cz2YPPKQi9IG2ywVPYu4go4YQALQYW4EfE%2BqxEER1WNgcQ5DKGMVwL3QqrxZarBdvoPQb8q4TAQPcLAIRyaQOgdnKO4dwjhCjuEHg4VQgYQwp06RIicHeLLhXd%2BExP7cDDq6GQczwqJyTvHTgUdIGhBXvEzRhDLbELrhAJAYjO4t2QrIzuIAKhOFTnUBErB%2BJlwgAEbBARvAVF2AvURZJyT0CirQVg7yp5YBJuoYegK8B3FMHgKclZlH5IoUyD55BNLWynqwZxOxiACywNg9GeBs5RN2QYoexjR5mIsZPQu1iVBqA0A4vQTiXHwHcVETx0hklRQ2fUchJRLDWD6LULR1hhjtC7tEWISJEj2BqKkcVJRhW5A6AMWskK6BNF6FK/oXKVWlB6C0bwbQFWisGOq5IArjV6uyCKrgExSS%2BOtTEzZhdi570PsfU%2B59L6MRvj4e%2BxSiClO9qQCpZzn5lLwpUgwNSA51IaR0ZpwdQiukvkkwZUco7cDZJwDCrpQjcAwi0x1a8xn4J2YHFF2hC2jKjZbCY0LiBxAsNwIAA%3D%3D",
		UrlCode:    "8k7D1cSFKuq",
		VisitCount: 0}

	err = UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)

	loadedRecord = UrlItem{}
	loadedRecord, err = UrlRecordRead(urlDataset.UrlCode)
	assert.Equal(t, nil, err)
	assert.Equal(t, loadedRecord.UrlBase, urlDataset.UrlBase)
}

func TestInsertRecordAndExistenceOfRecord(t *testing.T) {
	urlDataset := UrlItem{
		Id:         12685634646217438788,
		UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
		UrlCode:    "6FEC3Yjvghp",
		VisitCount: 0}

	err := UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)
	assert.Equal(t, UrlRecordIsExisting(urlDataset), true)

	urlDataset = *NewUrlItem()
	urlDataset.UrlCode = "543210uvwxyz"
	assert.Equal(t, UrlRecordIsExisting(urlDataset), false)
}

func Test_readRecord(t *testing.T) {
	var err error = nil
	urlDataset := UrlItem{
		Id:         12685634646217438788,
		UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
		UrlCode:    "6FEC3Yjvghp",
		VisitCount: 0}

	// fail read attempt
	err = readRecord(urlDataset.UrlCode, urlDataset)
	assert.NotEqual(t, nil, err)

	// write
	err = UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)
	// successful read attempt
	urlRecord := NewUrlItem()
	err = readRecord(urlDataset.UrlCode, urlRecord)
	assert.Equal(t, nil, err)
}

func Test_deleteRecord(t *testing.T) {
	urlDataset := UrlItem{
		Id:         12685634646217438788,
		UrlBase:    "https://github.com/FranzBranntvino/urlsGoShort",
		UrlCode:    "6FEC3Yjvghp",
		VisitCount: 0}

	err := UrlRecordWrite(urlDataset)
	assert.Equal(t, nil, err)

	var returnCode int64 = -1
	returnCode, err = deleteRecord(urlDataset.UrlCode)
	assert.Equal(t, nil, err)
	assert.Equal(t, returnCode, int64(1))

	returnCode = -1
	returnCode, err = deleteRecord(urlDataset.UrlCode)
	assert.Equal(t, nil, err)
	assert.Equal(t, returnCode, int64(0))
}

func TestCloseStore(t *testing.T) {
	assert.True(t, (CloseStore() == nil))
}
