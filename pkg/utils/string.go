package utils

import (
    "time"
    "math/rand"
)

// RandomString return the random string
func RandomString(length int) string {
    rand.Seed(time.Now().UnixNano())

    letters      := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
    randomValues := make([]rune, length)

    for i := range randomValues {
        randomValues[i] = letters[rand.Intn(len(letters))]
    }

    return string(randomValues)
}
