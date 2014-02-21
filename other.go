package smitego

type ItemDescription struct {
	Cooldown             string     `json:"cooldown"`
	Cost                 string     `json:"cost"`
	Description          string     `json:"description"`
	MenuItems            []MenuItem `json:"menuitems"`
	RankItems            []RankItem `json:"rankitems"`
	SecondaryDescription string     `json:"secondaryDescription"`
}

type MenuItem struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

type RankItem struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}
