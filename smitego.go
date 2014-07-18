package smitego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type MatchHistory struct {
	Activeid1      int
	Activeid2      int
	Active_1       string
	Active_2       string
	Active_3       string
	Assists        int
	Creeps         int
	Damage         int
	Damage_taken   int
	Deaths         int
	Itemid1        int
	Itemid2        int
	Itemid3        int
	Itemid4        int
	Itemid5        int
	Itemid6        int
	Killing_spree  int
	Kills          int
	Level          int
	Match          int
	Match_time     string
	Minutes        int
	Multi_kill_max int
	Queue          string
	Skin           string
	Skinid         int
	Surrendered    string
	Team1score     int
	Team2score     int
	Win_status     string
	Playername     string
	Ret_msg        string
}

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

func GetMatchHistory(playerName string) []MatchHistory {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := getMD5Hash(DevID + "getmatchhistory" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getmatchhistoryJson/" + DevID + "/" + hash + "/" + SessionID + "/" + timestamp + "/" + playerName

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var matchHistory []MatchHistory
			json.Unmarshal(contents, &matchHistory)
			return matchHistory
		}
	}

	matchHistory := []MatchHistory{}
	return matchHistory
}
