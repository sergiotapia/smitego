package smitego

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Returns ping information for the Smite API endpoint server.
func Ping() string {
	url := "http://api.smitegame.com/smiteapi.svc/pingJson"

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			return string(contents[:])
		}
	}
	return "Ping service not available."
}

// TODO: Create struct and find out why Request Error is
// happening on correct API call.
func GetQueueStats(playerName string, queueId int) string {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := getMD5Hash(DevID + "getqueuestats" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getqueuestatsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + timestamp + "/" + playerName + "/" + string(queueId)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			return string(contents)
		}
	}
	return ""
}

//TO-DO: Service Endpoint not found.
func GetTopRanked(queueId int) string {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := getMD5Hash(DevID + "gettopranked" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/gettoprankedJson/" + DevID + "/" + hash + "/" + SessionID + "/" + timestamp + "/" + string(queueId)

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			return string(contents)
		}
	}

	return ""
}
