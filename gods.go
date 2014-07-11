package smitego

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// God is a struct that represents a God character in Smite.
type God struct {
	Ability1                   string      `json:"Ability1"`
	Ability2                   string      `json:"Ability2"`
	Ability3                   string      `json:"Ability3"`
	Ability4                   string      `json:"Ability4"`
	Ability5                   string      `json:"Ability5"`
	Abilityid1                 int         `json:"AbilityId1"`
	Abilityid2                 int         `json:"AbilityId2"`
	Abilityid3                 int         `json:"AbilityId3"`
	Abilityid4                 int         `json:"AbilityId4"`
	Abilityid5                 int         `json:"AbilityId5"`
	Attackspeed                float32     `json:"AttackSpeed"`
	Attackspeedperlevel        float32     `json:"AttackSpeedPerLevel"`
	Cons                       string      `json:"Cons"`
	Hp5perlevel                float32     `json:"HP5PerLevel"`
	Health                     int         `json:"Health"`
	Healthperfive              int         `json:"HealthPerFive"`
	Healthperlevel             int         `json:"HealthPerLevel"`
	Item1                      string      `json:"Item1"`
	Item2                      string      `json:"Item2"`
	Item3                      string      `json:"Item3"`
	Item4                      string      `json:"Item4"`
	Item5                      string      `json:"Item5"`
	Item6                      string      `json:"Item6"`
	Item7                      string      `json:"Item7"`
	Item8                      string      `json:"Item8"`
	Item9                      string      `json:"Item9"`
	Itemid1                    int         `json:"ItemId1"`
	Itemid2                    int         `json:"ItemId2"`
	Itemid3                    int         `json:"ItemId3"`
	Itemid4                    int         `json:"ItemId4"`
	Itemid5                    int         `json:"ItemId5"`
	Itemid6                    int         `json:"ItemId6"`
	Itemid7                    int         `json:"ItemId7"`
	Itemid8                    int         `json:"ItemId8"`
	Itemid9                    int         `json:"ItemId9"`
	Lore                       string      `json:"Lore"`
	Mp5perlevel                float32     `json:"MP5PerLevel"`
	Magicprotection            int         `json:"MagicProtection"`
	Magicprotectionperlevel    int         `json:"MagicProtectionPerLevel"`
	Mana                       int         `json:"Mana"`
	Manaperfive                float32     `json:"ManaPerFive"`
	Manaperlevel               int         `json:"ManaPerLevel"`
	Name                       string      `json:"Name"`
	Onfreerotation             string      `json:"OnFreeRotation"`
	Pantheon                   string      `json:"Pantheon"`
	Physicalpower              int         `json:"PhysicalPower"`
	Physicalpowerperlevel      int         `json:"PhysicalPowerPerLevel"`
	Physicalprotection         int         `json:"PhysicalProtection"`
	Physicalprotectionperlevel float32     `json:"PhysicalProtectionPerLevel"`
	Pros                       string      `json:"Pros"`
	Roles                      string      `json:"Roles"`
	Speed                      int         `json:"Speed"`
	Title                      string      `json:"Title"`
	Type                       string      `json:"Type"`
	AbilityDescription1        Ability     `json:"abilityDescription1"`
	AbilityDescription2        Ability     `json:"abilityDescription2"`
	AbilityDescription3        Ability     `json:"abilityDescription3"`
	AbilityDescription4        Ability     `json:"abilityDescription4"`
	AbilityDescription5        Ability     `json:"abilityDescription5"`
	BasickAttack               BasicAttack `json:"basicAttack"`
	ReturnMessage              string      `json:"ret_msg"`
	ID                         int         `json:"id"`
}

// GodRecommendedItem is a struct that represents a God's recommended
// item to purchase.
type GodRecommendedItem struct {
	Category        string `json:"Category"`
	CategoryValueID int    `json:"category_value_id"`
	Name            string `json:"Item"`
	ItemID          int    `json:"item_id"`
	Role            string `json:"Role"`
	RoleValueID     int    `json:"role_value_id"`
	GodID           int    `json:"god_id"`
	GodName         string `json:"god_name"`
	ReturnMessage   string `json:"ret_msg"`
}

// BasicAttack is a struct that represents a God's basic attack information.
type BasicAttack struct {
	AttackDescription ItemDescription `json:"itemDescription"`
}

// Ability is a struct that represents a God's spell ability information.
type Ability struct {
	AbilityDescription ItemDescription `json:"itemDescription"`
}

// GetGods returns a new God array instance with all the God's and their
// information from Smite.
func GetGods() []God {
	hash := getMD5Hash(DevID + "getgods" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getgodsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp() + "/1"

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var gods []God
			json.Unmarshal(contents, &gods)
			return gods
		}
	}
	gods := []God{}
	return gods
}

// GetGodRecommendedItems returns a new GodRecommendedItem array instance
// with all the God's recommended items and their information from Smite.
func GetGodRecommendedItems(godID string) []GodRecommendedItem {
	hash := getMD5Hash(DevID + "getgodrecommendeditems" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getgodrecommendeditemsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp() + "/" + godID + "/1"

	response, err := http.Get(url)
	if err != nil {
		perror(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			perror(err)
		} else {
			var godRecommendedItems []GodRecommendedItem
			json.Unmarshal(contents, &godRecommendedItems)
			return godRecommendedItems
		}
	}
	godRecommendedItems := []GodRecommendedItem{}
	return godRecommendedItems
}
