package smitego

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// MatchPlayer is a struct that represents a player that participated in a Smite
// match.
type MatchPlayer struct {
	AccountLevel       int    `json:"Account_Level"`
	Activeid1          int    `json:"ActiveId1"`
	Activeid2          int    `json:"ActiveId2"`
	Assists            int    `json:"Assists"`
	Ban1               string `json:"Ban1"`
	Ban1id             int    `json:"Ban1Id"`
	Ban2               string `json:"Ban2"`
	Ban2id             int    `json:"Ban2Id"`
	Ban3               string `json:"Ban3"`
	Ban3id             int    `json:"Ban3Id"`
	Ban4               string `json:"Ban4"`
	Ban4id             int    `json:"Ban4Id"`
	DamageBot          int    `json:"Damage_Bot"`
	DamageDoneMagical  int    `json:"Damage_Done_Magical"`
	DamageDonePhysical int    `json:"Damage_Done_Physical"`
	DamagePlayer       int    `json:"Damage_Player"`
	DamageTaken        int    `json:"Damage_Taken"`
	Deaths             int    `json:"Deaths"`
	EntryDatetime      string `json:"Entry_Datetime"`
	FinalMatchLevel    int    `json:"Final_Match_Level"`
	Godid              int    `json:"GodId"`
	GoldEarned         int    `json:"Gold_Earned"`
	GoldPerMinute      int    `json:"Gold_Per_Minute"`
	Healing            int    `json:"Healing"`
	Itemid1            int    `json:"ItemId1"`
	Itemid2            int    `json:"ItemId2"`
	Itemid3            int    `json:"ItemId3"`
	Itemid4            int    `json:"ItemId4"`
	Itemid5            int    `json:"ItemId5"`
	Itemid6            int    `json:"ItemId6"`
	ItemActive1        string `json:"Item_Active_1"`
	ItemActive2        string `json:"Item_Active_2"`
	ItemActive3        string `json:"Item_Active_3"`
	ItemPurch1         string `json:"Item_Purch_1"`
	ItemPurch2         string `json:"Item_Purch_2"`
	ItemPurch3         string `json:"Item_Purch_3"`
	ItemPurch4         string `json:"Item_Purch_4"`
	ItemPurch5         string `json:"Item_Purch_5"`
	ItemPurch6         string `json:"Item_Purch_6"`
	KillingSpree       int    `json:"Killing_Spree"`
	KillsBot           int    `json:"Kills_Bot"`
	KillsPlayer        int    `json:"Kills_Player"`
	MasteryLevel       int    `json:"Mastery_Level"`
	Match              int    `json:"Match"`
	Minutes            int    `json:"Minutes"`
	MultiKillMax       int    `json:"Multi_kill_Max"`
	Partyid            int    `json:"PartyId"`
	ReferenceName      string `json:"Reference_Name"`
	Skin               string `json:"Skin"`
	Skinid             int    `json:"SkinId"`
	StructureDamage    int    `json:"Structure_Damage"`
	Surrendered        string `json:"Surrendered"`
	Team1score         int    `json:"Team1Score"`
	Team2score         int    `json:"Team2Score"`
	Teamid             int    `json:"TeamId"`
	TeamName           string `json:"Team_Name"`
	WinStatus          string `json:"Win_Status"`
	Name               string `json:"name"`
	Playername         string `json:"playerName"`
	RetMsg             string `json:"ret_msg"`
}

// ModeDetail is a struct that represents the mode information of a given match.
type ModeDetail struct {
	Ban1               string `json:"Ban1"`
	Ban2               string `json:"Ban2"`
	EntryDatetime      string `json:"Entry_Datetime"`
	Match              int    `json:"Match"`
	MatchTime          int    `json:"Match_Time"`
	OfflineSpectators  int    `json:"Offline_Spectators"`
	RealtimeSpectators int    `json:"Realtime_Spectators"`
	RecordingEnded     string `json:"Recording_Ended"`
	RecordingStarted   string `json:"Recording_Started"`
	Team1Avglevel      int    `json:"Team1_AvgLevel"`
	Team1Gold          int    `json:"Team1_Gold"`
	Team1Kills         int    `json:"Team1_Kills"`
	Team1Score         int    `json:"Team1_Score"`
	Team2Avglevel      int    `json:"Team2_AvgLevel"`
	Team2Gold          int    `json:"Team2_Gold"`
	Team2Kills         int    `json:"Team2_Kills"`
	Team2Score         int    `json:"Team2_Score"`
	WinningTeam        int    `json:"Winning_Team"`
	ReturnMessage      string `json:"ret_msg"`
}

// GetModeDetails returns a ModeDetail instance loaded with mode information
// for a specific match ID.
func GetModeDetails(matchID string) ModeDetail {
	hash := getMD5Hash(DevID + "getdemodetails" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getdemodetailsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp() + "/" + matchID

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var modeDetail []ModeDetail
			json.Unmarshal(contents, &modeDetail)
			return modeDetail[0]
		}
	}
	modeDetail := ModeDetail{ReturnMessage: "Not found"}
	return modeDetail
}

// GetMatchDetails returns a MatchPlayer array instance loaded with each
// player's information for the match.
func GetMatchDetails(matchID string) []MatchPlayer {
	hash := getMD5Hash(DevID + "getmatchdetails" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getmatchdetailsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp() + "/" + matchID

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var matchDetail []MatchPlayer
			json.Unmarshal(contents, &matchDetail)
			return matchDetail
		}
	}

	matchDetail := []MatchPlayer{}
	return matchDetail
}
