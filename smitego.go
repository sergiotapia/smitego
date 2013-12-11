package smitego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var SessionId = ""
var DevId = ""
var AuthKey = ""

// Queue ID's
const (
	Conquest5v5               = 423
	Novice                    = 424
	Conquest                  = 426
	Practice                  = 427
	Challenge                 = 429
	Joust                     = 431
	Domination                = 433
	MatchOfTheDay             = 434
	Arena                     = 435
	ArenaChallenge            = 438
	DominationChallenge       = 439
	JoustQueued               = 440
	RenaissanceJoustChallenge = 441
	ConquestTeamRanked        = 442
	RankedConquest            = 430
	RankedJoust1v1            = 449
	RankedJoust3v3            = 450
	MasteryConquest           = 451
	LeagueArena               = 451
	LeagueConquest            = 452
	LeagueBronze              = 1
	LeagueSilver              = 2
	LeagueGold                = 3
	LeaguePlatinum            = 4
	LeagueDiamond             = 5
)

type Player struct {
	Created_datetime    string
	Last_login_datetime string
	Leaves              int
	Level               int
	Losses              int
	Masterylevel        int
	Name                string
	Rank_confidence     int
	Rank_stat           float32
	Teamid              int
	Team_name           string
	Wins                int
	Ret_msg             string
}

func GetPlayer(playerName string) Player {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := GetMD5Hash(DevId + "getplayer" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getplayerJson/" + DevId + "/" + hash + "/" + SessionId + "/" + timestamp + "/" + playerName

	response, err := http.Get(url)
	if err != nil {
		Perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			Perror(err)
		} else {
			var players []Player
			json.Unmarshal(contents, &players)
			fmt.Println(string(contents))
			return players[0]
		}
	}

	player := Player{Ret_msg: "Not found"}
	return player
}

//TO-DO: Service Endpoint not found.
func GetGods() string {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := GetMD5Hash(DevId + "getgods" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getgodsJson/" + DevId + "/" + hash + "/" + SessionId + "/" + timestamp

	response, err := http.Get(url)
	if err != nil {
		Perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			Perror(err)
		} else {
			return string(contents)
		}
	}
	return ""
}

//TO-DO: The server encountered an error processing the request. See server logs for more details.
func GetItems() string {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := GetMD5Hash(DevId + "getitems" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getitemsJson/" + DevId + "/" + hash + "/" + SessionId + "/" + timestamp + "/en"

	response, err := http.Get(url)
	if err != nil {
		Perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			Perror(err)
		} else {
			return string(contents)
		}
	}

	return ""
}

//TO-DO: Service Endpoint not found.
func GetTopRanked(queueId int) string {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := GetMD5Hash(DevId + "gettopranked" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/gettoprankedJson/" + DevId + "/" + hash + "/" + SessionId + "/" + timestamp + "/" + string(queueId)

	response, err := http.Get(url)
	if err != nil {
		Perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			Perror(err)
		} else {
			return string(contents)
		}
	}

	return ""
}

//TO-DO: Parse into actual struct type.
func GetMatchDetails(matchId string) string {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := GetMD5Hash(DevId + "getmatchdetails" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getmatchdetailsJson/" + DevId + "/" + hash + "/" + SessionId + "/" + timestamp + "/" + matchId

	response, err := http.Get(url)
	if err != nil {
		Perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			Perror(err)
		} else {
			return string(contents)
		}
	}

	return ""
}
