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

// DataUsed is a struct that represents a how many resources you have used
// so far in the HiRez Smite API server.
type DataUsed struct {
	ActiveSession      int    `json:"Active_Sessions"`
	ConcurrentSessions int    `json:"Concurrent_Sessions"`
	RequestLimitDaily  int    `json:"Request_Limit_Daily"`
	SessionCap         int    `json:"Session_Cap"`
	SessionTimeLimit   int    `json:"Session_Time_Limit"`
	TotalRequestsToday int    `json:"Total_Requests_Today"`
	TotalSessionsToday int    `json:"Total_Sessions_Today"`
	ReturnMessage      string `json:"ret_msg"`
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

// GetDataUsed returns a DataUsed instance loaded with API usage information
// on the HiRez Smite API server.
func GetDataUsed() DataUsed {
	hash := getMD5Hash(DevID + "getdataused" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getdatausedJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp()

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var dataUsed []DataUsed
			json.Unmarshal(contents, &dataUsed)
			return dataUsed[0]
		}
	}
	dataUsed := DataUsed{ReturnMessage: "Not found"}
	return dataUsed
}
