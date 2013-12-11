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

func GetGods() string {
	//getgods[ResponseFormat]/{devId}/{signature}/{sessionId}/{timestamp}/{languageCode}
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
			/*var session Session
			json.Unmarshal(contents, &session)
			SessionId = session.Session_id*/
			return string(contents)
		}
	}
	return ""
}
