package stalcraftgo

type CharacterClanInfo struct {
	Info   Info   `json:"info"`
	Member Member `json:"member"`
}

type Info struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Tag              string `json:"tag"`
	Level            int32  `json:"level"`
	LevelPoints      int32  `json:"level_points"`
	RegistrationTime string `json:"registration_time"`
	Alliance         string `json:"alliance"`
	Description      string `json:"description"`
	Leader           string `json:"leader"`
	MemberCount      int32  `json:"member_count"`
}

type Member struct {
	Name     string `json:"name"`
	Rank     string `json:"rank"`
	JoinTime string `json:"join_time"`
}
