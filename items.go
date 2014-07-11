package smitego

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Item is a struct that represents an Item in Smite.
type Item struct {
	ChilditemID      int             `json:"ChildItemId"`
	DeviceName       string          `json:"DeviceName"`
	IconID           int             `json:"IconId"`
	ItemDescription  ItemDescription `json:"itemDescription"`
	ItemID           int             `json:"ItemId"`
	ItemTier         int             `json:"ItemTier"`
	Price            int             `json:"Price"`
	RootItemID       int             `json:"RootItemId"`
	ShortDescription string          `json:"ShortDesc"`
	StartingItem     bool            `json:"StartingItem"`
	Type             string          `json:"Type"`
	RetMsg           string          `json:"ret_msg"`
}

// GetItems returns a new Item array instance with all the Items and their
// information from Smite.
func GetItems() []Item {
	hash := getMD5Hash(DevID + "getitems" + AuthKey + getTimestamp())
	url := "http://api.smitegame.com/smiteapi.svc/getitemsJson/" + DevID + "/" + hash + "/" + SessionID + "/" + getTimestamp() + "/1"

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
