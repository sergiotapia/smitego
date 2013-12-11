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

type MatchPlayer struct {
	Account_level        int
	Activeid1            int
	Activeid2            int
	Assists              int
	Ban1                 int
	Ban1Id               int
	Ban2                 int
	Ban2Id               int
	Damage_bot           int
	Damage_done_magical  int
	Damage_done_physical int
	Damage_player        int
	Damage_taken         int
	Deaths               int
	Entry_datetime       string
	Final_match_level    int
	Godid                int
	Gold_earned          int
	Gold_per_minute      int
	Healing              int
	Itemid1              int
	Itemid2              int
	Itemid3              int
	Itemid4              int
	Itemid5              int
	Itemid6              int
	Item_active_1        string
	Item_active_2        string
	Item_active_3        string
	Item_purch_1         string
	Item_purch_2         string
	Item_purch_3         string
	Item_purch_4         string
	Item_purch_5         string
	Item_purch_6         string
	Killing_spree        int
	Kills_bot            int
	Kills_player         int
	Master_level         int
	Match                int
	Minutes              int
	Multi_kill_max       int
	Partyid              int
	Reference_name       string
	Skin                 string
	Skinid               int
	Structure_damage     int
	Surrendered          int
	Team1score           int
	Team2score           int
	Teamid               int
	Team_name            string
	Win_status           string
	Name                 string
	Playername           string
	Ret_msg              string
}

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

func GetMatchDetails(matchId string) []MatchPlayer {
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
			var matchPlayers []MatchPlayer
			json.Unmarshal(contents, &matchPlayers)
			return matchPlayers
		}
	}

	matchPlayers := []MatchPlayer{}
	return matchPlayers
}

func GetMatchHistory(playerName string) []MatchHistory {
	timestamp := time.Now().UTC().Format("20060102150405")
	hash := GetMD5Hash(DevId + "getmatchhistory" + AuthKey + timestamp)
	url := "http://api.smitegame.com/smiteapi.svc/getmatchhistoryJson/" + DevId + "/" + hash + "/" + SessionId + "/" + timestamp + "/" + playerName

	response, err := http.Get(url)
	if err != nil {
		Perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			Perror(err)
		} else {
			var matchHistory []MatchHistory
			json.Unmarshal(contents, &matchHistory)
			return matchHistory
		}
	}

	matchHistory := []MatchHistory{}
	return matchHistory
}
