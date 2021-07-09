// The package urlShortener provides basic functionality to encode Url strings within a given alphabet
// by making use of permutation with repetition and positional notation.
// See https://en.wikipedia.org/wiki/Permutation#Permutations_with_repetition and https://en.wikipedia.org/wiki/Positional_notation
// Decoding Functionality is not provided.
package urlShortener

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Imports:

import (
	"errors"
	"hash/fnv"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Constants:

// Encoding with a number of encoding possibilities: p = Base^CodeLength
// For this particular exercise the Base for the encoding should be about length of the alphabet.
const (
  // The desired outputs encoded string length
  maxCodeLength    = 11
  // URL alphabet according to: https://www.ietf.org/rfc/rfc3986.txt
  urlAlphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789:./-_~?#[]@!$&'()*+,;%="
  urlALength    = uint64(len(urlAlphabet))
  // Define the alphabet for encoding strings by permutation with repetition.
  encAlphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  encALength    = uint64(len(encAlphabet))
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exported:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Implementation:

// GetUrlEncoding, permutate an URL string into it's coded form as a string based on a pre-defined alphabet.
// Returns the Encoded Url and its positional Code.
func GetUrlEncoding(inputUrl string) (urlCode string, positionalCode uint64, err error) {
  positionalCode, err = GeneratePositionalCode(inputUrl)
  if err != nil {
    urlCode = "FailedToEncodeUrl"
  } else {
    urlCode, err = Encode(positionalCode)
  }
  return urlCode, positionalCode, err
}

////////////////////////////////////////////////////////////

// Convert an input string (e.g. a valid URL) by the mapped input alphabet into a positional number.
func GeneratePositionalCode(inputStr string) (uint64, error) {
  for _, symbol := range inputStr {
    charIndex := strings.IndexRune(urlAlphabet, symbol)
    if (charIndex < 0) {
      return uint64(charIndex), errors.New("Character: " + string(symbol) + " not valid for provided URL-alphabet.")
    }
  }

  hashing := fnv.New64a()
  _, err := hashing.Write([]byte(inputStr))
  return hashing.Sum64(), err
}

// Encode, convert a provided positional number as uint64 into it's coded form as a string.
// The encoding is based on a pre-defined alphabet, its length and outputs the encoding string with a max-code-length.
func Encode(pNumber uint64) (string, error) {
  var encodedBuilder strings.Builder
  var err error
  encodedBuilder.Grow(maxCodeLength)
  if (pNumber == 0) {
  	encodedBuilder.WriteByte(encAlphabet[(pNumber % encALength)])
  } else {
    for ; pNumber > 0; pNumber = (pNumber / encALength) {
      encodedBuilder.WriteByte(encAlphabet[(pNumber % encALength)])
    }
  }
  return encodedBuilder.String(), err
}
