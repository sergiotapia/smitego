package smitego

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// LeagueConquest is a struct that represents a player's status within
// the Conquest league in Smite.
type LeagueConquest struct {
	Leaves        int    `json:"Leaves"`
	Losses        int    `json:"Losses"`
	Name          string `json:"Name"`
	Points        int    `json:"Points"`
	PreviousRank  int    `json:"PrevRank"`
	Rank          int    `json:"Rank"`
	Season        int    `json:"Season"`
	Tier          int    `json:"Tier"`
	Trend         int    `json:"Trend"`
	VictoryPoints int    `json:"VictoryPoints"`
	Wins          int    `json:"Wins"`
	ReturnMessage string `json:"ret_msg"`
}

// LeagueJoust is a struct that represents a player's status within
// the Joust league in Smite.
type LeagueJoust struct {
	Leaves        int    `json:"Leaves"`
	Losses        int    `json:"Losses"`
	Name          string `json:"Name"`
	Points        int    `json:"Points"`
	PreviousRank  int    `json:"PrevRank"`
	Rank          int    `json:"Rank"`
	Season        int    `json:"Season"`
	Tier          int    `json:"Tier"`
	Trend         int    `json:"Trend"`
	VictoryPoints int    `json:"VictoryPoints"`
	Wins          int    `json:"Wins"`
	ReturnMessage string `json:"ret_msg"`
}

// Player is a struct that represents a Smite game player.
type Player struct {
	ID                int            `json:"Id"`
	CreatedDateTime   string         `json:"Created_Datetime"`
	LastLoginDateTime string         `json:"Last_Login_Datetime"`
	Leaves            int            `json:"Leaves"`
	Level             int            `json:"Level"`
	LeagueConquest    LeagueConquest `json:"LeagueConquest"`
	LeagueJoust       LeagueJoust    `json:"LeagueJoust"`
	Losses            int            `json:"Losses"`
	MasteryLevel      int            `json:"MasteryLevel"`
	Name              string         `json:"Name"`
	RankStat          float32        `json:"Rank_Stat"`
	RankStatJoust     float32        `json:"Rank_Stat_Joust"`
	TeamID            int            `json:"TeamId"`
	TeamName          string         `json:"Team_Name"`
	TierConquest      int            `json:"Tier_Conquest"`
	TierJoust         int            `json:"Tier_Joust"`
	Wins              int            `json:"Wins"`
	ReturnMessage     string         `json:"ret_msg"`
}

// GodRank is a struct that represents a player's god rank for the Gods
// in Smite.
type GodRank struct {
	Rank          int    `json:"rank"`
	Worshippers   int    `json:"worshippers"`
	God           string `json:"god"`
	ReturnMessage string `json:"ret_msg"`
}

// Friend is a struct that represents a Smite player's friend.
type Friend struct {
	Name          string `json:"name"`
	ReturnMessage string `json:"ret_msg"`
}

// GetPlayer returns a new Player instance with the fields filled with
// data from the json response.
func GetPlayer(playerName string) Player {
	hash := getMD5Hash(DevID + "getplayer" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getplayerJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp() + "/" + playerName

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var players []Player
			json.Unmarshal(contents, &players)
			return players[0]
		}
	}

	player := Player{ReturnMessage: "Not found"}
	return player
}

// GetGodRanks returns a new GodRank array instance with a collection of a
// player's god rank for each God in Smite.
func GetGodRanks(playerName string) []GodRank {
	hash := getMD5Hash(DevID + "getgodranks" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getgodranksJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp() + "/" + playerName

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var godRanks []GodRank
			json.Unmarshal(contents, &godRanks)
			return godRanks
		}
	}
	godRanks := []GodRank{}
	return godRanks
}

// GetFriends returns a new Friend array instance with a collection of a
// player's friends.
func GetFriends(playerName string) []Friend {
	hash := getMD5Hash(DevID + "getfriends" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getfriendsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp() + "/" + playerName

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var friends []Friend
			json.Unmarshal(contents, &friends)
			return friends
		}
	}
	friends := []Friend{}
	return friends
}
