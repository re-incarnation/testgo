package stalcraftgo

type FullCharacterInfo struct {
	Information Information `json:"information"`
	Clan        Clan        `json:"clan"`
}

type Information struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	CreationTime string `json:"creation_time"`
}
