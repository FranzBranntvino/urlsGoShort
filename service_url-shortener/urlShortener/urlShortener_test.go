package urlShortener

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Imports:

import (
	"fmt"
	"math"
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

func TestConversionSuccess(t *testing.T) {
	assert := assert.New(t)

	testLink := "https://github.com/FranzBranntvino/urlsGoShort"
	number, err := GeneratePositionalCode(testLink)
	var expect uint64 = 12685634646217438788
	fmt.Println("Returned: " + fmt.Sprint(number) + " with Error: " + fmt.Sprint(err) + " for: " + fmt.Sprint(testLink))
	assert.Equal(expect, number)
	assert.Equal(nil, err)

	// expected success ...
	longLinks := [...]string{
		"https://www.google.com",
		"https://google.com",
		"www.google.com",
		"https://www.google.net",
		"https://www.gmx.net",
		"https://www.gmx.de",
	}
	for _, longLink := range longLinks {
		number, err = GeneratePositionalCode(longLink)
		fmt.Println("Returned: " + fmt.Sprint(number) + " with Error: " + fmt.Sprint(err) + " for: " + fmt.Sprint(longLink))
		assert.Equal(nil, err)
	}
}

func TestConversionFail(t *testing.T) {
	assert := assert.New(t)
	// expected error ...
	longLinks := [...]string{
		"https://google.com``",
		"https://gooßle.com",
		"www.gooÖgle.com",
	}
	for _, longLink := range longLinks {
		number, err := GeneratePositionalCode(longLink)
		var expect uint64 = math.MaxUint64
		fmt.Println("Returned: " + fmt.Sprint(number) + " with Error: " + fmt.Sprint(err))
		assert.Equal(expect, number)
		assert.NotEqual(nil, err)
	}
}

func TestEncoding(t *testing.T) {
	rString := Encode(6833151723267039701)
	expectS := "9GkkU9p4Vii"
	assert.Equal(t, rString, expectS)
}

func TestUrlEncoding(t *testing.T) {
	rString, rNumber, err := GetUrlEncoding("https://github.com/FranzBranntvino/urlsGoShort")
	expectN := uint64(12685634646217438788)
	expectS := "6FEC3Yjvghp"
	assert := assert.New(t)
	assert.Equal(expectN, rNumber)
	assert.Equal(expectS, rString)
	assert.Equal(nil, err)

	longLinks := [...]string{
		"https://github.com/FranzBranntvino/urlsGoShort",
		"github.com/FranzBranntvino/urlsGoShort/",
		"https://www.google.com",
		"https://google.com",
		"www.google.com",
	}
	for _, longLink := range longLinks {
		rString, _, err := GetUrlEncoding(longLink)
		fmt.Println("Returned: " + fmt.Sprint(rString) + " with Error: " + fmt.Sprint(err) + " for: " + fmt.Sprint(longLink))
		assert.Equal(nil, err)
	}
}
