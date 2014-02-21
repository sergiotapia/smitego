package smitego

type Item struct {
	Childitemid     int             `json:"ChildItemId"`
	Devicename      string          `json:"DeviceName"`
	Iconid          int             `json:"IconId"`
	ItemDescription ItemDescription `json:"itemDescription"`
	Itemid          int             `json:"ItemId"`
	Itemtier        int             `json:"ItemTier"`
	Price           int             `json:"Price"`
	Rootitemid      int             `json:"RootItemId"`
	Shortdesc       string          `json:"ShortDesc"`
	Startingitem    bool            `json:"StartingItem"`
	Type            string          `json:"Type"`
	RetMsg          string          `json:"ret_msg"`
}
