package smitego

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SessionID is the current assigned session id given by the Smite API server.
var SessionID = ""

// DevID is the unique developer id assigned by HiRez when requesting API
// access.
var DevID = ""

// AuthKey is the unique authentication key assigned by HiRez when requestion
// API access.
var AuthKey = ""

type session struct {
	ReturnMessage string `json:"ret_msg"`
	SessionID     string `json:"session_id"`
	Timestamp     string `json:"timestamp"`
}

// GetSessionID populates the SessionID with a current and valid session id
// assigned by the Smite API server.
func GetSessionID() {
	hash := getMD5Hash(DevID + "createsession" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/createsessionJson/" + DevID + "/" + hash + "/" + getTimestamp()

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var session session
			json.Unmarshal(contents, &session)
			SessionID = session.SessionID
		}
	}
}
