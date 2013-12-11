package smitego

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Session struct {
	Ret_msg    string
	Session_id string
	Timestramp string
}

func GetSessionId() {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := GetMD5Hash(DevId + "createsession" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/createsessionJson/" + DevId + "/" + hash + "/" + timestamp

	response, err := http.Get(url)
	if err != nil {
		Perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			Perror(err)
		} else {
			var session Session
			json.Unmarshal(contents, &session)
			SessionId = session.Session_id
		}
	}
}

func GetJson(url string) string {
	response, err := http.Get(url)
	Perror(err)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	Perror(err)

	return string(body)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Perror(err error) {
	if err != nil {
		panic(err)
	}
}
