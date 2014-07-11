package smitego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Queue ID's
// const (
// 	Conquest5v5               = 423
// 	Novice                    = 424
// 	Conquest                  = 426
// 	Practice                  = 427
// 	Challenge                 = 429
// 	Joust                     = 431
// 	Domination                = 433
// 	MatchOfTheDay             = 434
// 	Arena                     = 435
// 	ArenaChallenge            = 438
// 	DominationChallenge       = 439
// 	JoustQueued               = 440
// 	RenaissanceJoustChallenge = 441
// 	ConquestTeamRanked        = 442
// 	RankedConquest            = 430
// 	RankedJoust1v1            = 449
// 	RankedJoust3v3            = 450
// 	MasteryConquest           = 451
// 	LeagueArena               = 451
// 	LeagueConquest            = 452
// 	LeagueBronze              = 1
// 	LeagueSilver              = 2
// 	LeagueGold                = 3
// 	LeaguePlatinum            = 4
// 	LeagueDiamond             = 5
// )

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

func GetItems() []Item {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := getMD5Hash(DevID + "getitems" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getitemsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + timestamp + "/1"

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var items []Item
			json.Unmarshal(contents, &items)
			return items
		}
	}
	items := []Item{}
	return items
}

// TODO: Create struct and find out why Request Error is
// happening on correct API call.
func GetLeagueSeasons(queueId int) string {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := getMD5Hash(DevID + "getleagueseasons" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getleagueseasonsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + timestamp + "/" + string(queueId)
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

func GetMatchDetails(matchId string) []MatchPlayer {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := getMD5Hash(DevID + "getmatchdetails" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getmatchdetailsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + timestamp + "/" + matchId

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			fmt.Println("Test")
			var matchDetail []MatchPlayer
			json.Unmarshal(contents, &matchDetail)
			return matchDetail
		}
	}

	matchDetail := []MatchPlayer{}
	return matchDetail
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
