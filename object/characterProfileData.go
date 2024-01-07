package stalcraftgo

type CharacterProfileData struct {
	Username              string   `json:"username"`
	Uuid                  string   `json:"uuid"`
	Status                string   `json:"status"`
	Alliance              string   `json:"alliance"`
	LastLogin             string   `json:"last_login"`
	DisplayedAchievements []string `json:"displayed_achievements"`
	Clan                  Clan     `json:"clan"`
	Stats                 Stats    `json:"stats"`
}

type Clan struct {
	Info   Info   `json:"info"`
	Member Member `json:"member"`
}

type Stats struct {
	Id    string      `json:"id"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
