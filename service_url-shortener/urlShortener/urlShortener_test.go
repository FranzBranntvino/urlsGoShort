package urlShortener

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

var longLinks = [...]string{
    "https://www.google.com",
    "https://google.com",
    "www.google.com",
    "https://www.google.net",
    "https://www.gmx.net",
    "https://www.gmx.de",
    "https://www.google.de/maps/dir/Pescara,+Province+of+Pescara/Split,+Split,+Croatia/@43.0577057,14.4071143,9z/data=!3m1!4b1!4m13!4m12!1m5!1m1!1s0x1331a60db9286477:0xa0b89e89b22cbfe2!2m2!1d14.2160898!2d42.4617902!1m5!1m1!1s0x133560aa837e8095:0xad4b199c6d9344b0!2m2!1d16.4429578!2d43.5047301",
    "https://wego.here.com/directions/mix/Berlin-Hauptbahnhof,-Europaplatz-1,-Moabit,-10557-Berlin:276u33db-c5224e41937f46c48c4a342b79a1a716/Colosseum,-Piazza-del-Colosseo,-00184-Rome:380sr2yk-1ad6e01f4dcf495080cddf487062482b?map=47.47676,12.07512,6,normal",
    "https://godbolt.org/z/qrTT8GTbK",
    "https://godbolt.org/#z:OYLghAFBqd5QCxAYwPYBMCmBRdBLAF1QCcAaPECAM1QDsCBlZAQwBtMQBGAFlJvoCqAZ0wAFAB4gA5AAYppAFZdSrZrVDIApACYAQjt2kR7ZATx1KmWugDCqVgFcAtrWVX0AGTy1MAOWcARpjEIADMAJykAA6oQoTmtHaOLsoxcWZ0Xj7%2BTkEhEUaYJhm0DATMxARJzq6chcUJZRUEWX6BwWGRQuWV1Sl13c2tOXmdAJRGqA7EyBxSOqHeyI5YANSaoTYVxMwAnhvYmjIAggtLK5jrm%2BbdxJjMTgdHp9qLtMsOaxs2yLfewE8Ts9ngB6ABUqzKxAcplWWCo3nidCEqzBIOeBF2UUw8NWtxhBFWAH0hKgnJgicwHEQicArMFmAQcUS0F8AOz6E6rbmrAKoeyrKLEVDYyq7GQbTnHHm8/msQXC0WYziS54yvkCoUi4KY7Sqrk8jXyrVK3ahfXSw1yhXasXcC3q60mnW7ACsFs0bIAIniyRSqTS6T4dkz0CyMJgPScQTHVgBxekhnGrNh4ZgomjEVYAFSxkZODji6lWokVOrwmCEACUivcRFcfaTyZTqahaYnGczWZHQlLnoMzMgS2XKhXq7X05cAG5sByYXQOKhUYIN9YctU8qhsESkDfcresHd71YHo8G/fbzC788ny/Xy0Xw9X4HeqPHJlOKKqJnfZbplHZoCxxTqgeDoKsqgAF67HYn4VJgUIEhAaC0N0OY6AAbKsM6sGMa5SrGPIDngQ7Zths7zouy7EA6PJ4FQqwQIx3ToCAIDkk4yCfhAdzeEyxBCpgBAsumBC/ggFSrCh3RggcECYThc4LkuwQTKsvH0MEgnCSw3TiZJ0kELJoTYPJ2gYThalxJBmCoFQEDZmMeFgGAGw%2BjIeF4Z6UoyjKLFsWg1JXDY3zrNo2irL4ADyZGYAAjg4bBhRF3yhf5IDuKwtG%2BeRjiUSpWZubl2Vrj6RT1t5x45dy6WBYSqWbMlqzxYl8o6ClmxpQQrEZdYWW9senpei%2Bw1AmNxzEUOhmYOIQq%2Bs2AZtkGDKhuGawANalraZiVjW7CTpwq6VTep7Pidd7Hqd94yldl0XedT73kNb79uUg5SciTKzVmTb%2Bq27bBp2Ybdqsm0jjt477SIHU%2BsdD63o9d2Iw9Z7wwQ0JnfDt0o5jz0DeNzx8asTjMN4EBeeuN61VM9WdY11m2fZin5dRXl0zYeLdWxmVvjKUEwWSUTwYhpgQGD21jntdaYJwYwlfzsFC3cIsEGLW1KpLE7Q3L%2BOnDedwENMtCrBKuvPeNUgTKw0iuvIriyPIqDSCFegGL60yzGFoScPIBDSHITmkOtIDcDIAB0oQyAAHNobJR2y3A8JwrrhOZKjSNw8hOFwMgyKQ9tyKQTtSPIQggHnfsOxMcCwDAiAgDTUTUuQlBoJ%2BeDsCEdQ4oQJDKPwjAsOwXC8APwhiJIDukAA7jsUTSD7Vs23b/uO9IUXUk3hJ2asxxVgAsqswDIEO4Rh9oYeHRAtB0JgeEQIrHcrgssu%2B6vgfB9oUdhzIoRsqEUdvZsjZK6NkudU7pykJnUg2dOBn1zpwaOoC2QYVCJHGQKD86ryLtIUu5dSCVwDqQGu9cZqYGQK2MgFAIAVGAEIUQagigMAQKgae9sfakDblEDujIEgMJ8KwZhrCC7yC4U/EIwBEGhE4YLcRG9GAsLYdgshyBjjEDobg0gKiyj4HtvIAeTA2AcB4HwOgBBx4SGwUoOoqh1AoAMAYFQeAAhl0gBMEUJQy4l0mB7YxRhureH4UwxR7D5Cz2YPPKQi9IG2ywVPYu4go4YQALQYW4EfE%2BqxEER1WNgcQ5DKGMVwL3QqrxZarBdvoPQb8q4TAQPcLAIRyaQOgdnKO4dwjhCjuEHg4VQgYQwp06RIicHeLLhXd%2BExP7cDDq6GQczwqJyTvHTgUdIGhBXvEzRhDLbELrhAJAYjO4t2QrIzuIAKhOFTnUBErB%2BJlwgAEbBARvAVF2AvURZJyT0CirQVg7yp5YBJuoYegK8B3FMHgKclZlH5IoUyD55BNLWynqwZxOxiACywNg9GeBs5RN2QYoexjR5mIsZPQu1iVBqA0A4vQTiXHwHcVETx0hklRQ2fUchJRLDWD6LULR1hhjtC7tEWISJEj2BqKkcVJRhW5A6AMWskK6BNF6FK/oXKVWlB6C0bwbQFWisGOq5IArjV6uyCKrgExSS%2BOtTEzZhdi570PsfU%2B59L6MRvj4e%2BxSiClO9qQCpZzn5lLwpUgwNSA51IaR0ZpwdQiukvkkwZUco7cDZJwDCrpQjcAwi0x1a8xn4J2YHFF2hC2jKjZbCY0LiBxAsNwIAA%3D%3D",
}

var wrongLinks = [...]string{
    "https://google.com``",
    "https://gooßle.com",
    "www.gooÖgle.com",
}

var pNumStringMap = map[uint64]string {
    6833151723267039701 : "9GkkU9p4Vii",
    0 : "a",
    1 : "b",
    math.MaxUint64 : "pIrkgbKrQ8v",
}

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
    for _, longLink := range longLinks {
        number, err = GeneratePositionalCode(longLink)
        fmt.Println("Returned: " + fmt.Sprint(number) + " with Error: " + fmt.Sprint(err) + " for: " + fmt.Sprint(longLink))
        assert.Equal(nil, err)
    }
}

func TestConversionFail(t *testing.T) {
    assert := assert.New(t)
    // expected error ...
    for _, longLink := range wrongLinks {
        number, err := GeneratePositionalCode(longLink)
        var expect uint64 = math.MaxUint64
        fmt.Println("Returned: " + fmt.Sprint(number) + " with Error: " + fmt.Sprint(err))
        assert.Equal(expect, number)
        assert.NotEqual(nil, err)
    }
}

func TestEncoding(t *testing.T) {
    for pInput, expectS := range pNumStringMap {
        rString, err := Encode(pInput)
        assert.GreaterOrEqual(t, maxCodeLength, len(rString))
        assert.Equal(t, expectS, rString)
        assert.Equal(t, nil, err)
    }
}

func TestGetUrlEncoding(t *testing.T) {
    rString, rNumber, err := GetUrlEncoding("https://github.com/FranzBranntvino/urlsGoShort")
    expectN := uint64(12685634646217438788)
    expectS := "6FEC3Yjvghp"
    assert := assert.New(t)
    assert.Equal(expectN, rNumber)
    assert.Equal(expectS, rString)
    assert.Equal(nil, err)

    for _, longLink := range longLinks {
        rString, _, err := GetUrlEncoding(longLink)
        fmt.Println("Returned: " + fmt.Sprint(rString) + " with Error: " + fmt.Sprint(err) + " for: " + fmt.Sprint(longLink))
        assert.Equal(nil, err)
    }

    for _, wrongLink := range wrongLinks {
        rString, _, err := GetUrlEncoding(wrongLink)
        fmt.Println("Returned: " + fmt.Sprint(rString) + " with Error: " + fmt.Sprint(err) + " for: " + fmt.Sprint(wrongLink))
        assert.NotEqual(nil, err)
    }
}
