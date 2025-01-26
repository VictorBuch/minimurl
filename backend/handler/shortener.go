package handler

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

type ShortLink struct {
	OriginalURL, ShortenedURL string
}

func (sl *ShortLink) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := r.Form.Get("url")
	if u == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Url field is empty"))
	}
	short := GenerateShortenedURL()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("http://www.short.victorbuch.com/" + short))

}
func (sl *ShortLink) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List method called on short link")
}
func (sl *ShortLink) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetByID method called on short link")
}
func (sl *ShortLink) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateByID method called on short link")
}
func (sl *ShortLink) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteByID method called on short link")
}

func uniqid(prefix string) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000

	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
}

func GenerateShortenedURL() string {
	var (
		randomChars   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
		randIntLength = 27
		stringLength  = 32
	)

	str := make([]rune, stringLength)

	for char := range str {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(randIntLength)))
		if err != nil {
			panic(err)
		}

		str[char] = randomChars[nBig.Int64()]
	}

	hash := sha256.Sum256([]byte(uniqid(string(str))))
	encodedString := base64.StdEncoding.EncodeToString(hash[:])

	return encodedString[0:9]
}
