package smitego

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"time"
)

func getTimestamp() string {
	return time.Now().UTC().Format("20060102150405")
}

func getJSON(url string) string {
	response, err := http.Get(url)
	perror(err)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	perror(err)

	return string(body)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func perror(err error) {
	if err != nil {
		panic(err)
	}
}
