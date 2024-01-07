package stalcraftgo

type ClansListResponse struct {
	TotalClans int32  `json:"total_clans"`
	Data       []Data `json:"data"`
}

type Data struct {
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
