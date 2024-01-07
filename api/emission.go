package stalcraftgo

import stalcraftgo "github.com/re-incarnation/stalcraftgo/object"

func (exbo *EXBO) EmissionStatus(Region string, params Params) (response stalcraftgo.AuthResponse, err error) {
	//Used with access_token
	//exp : Bearer {AcessToken}
	err = exbo.RequestUnmarshalEmission(Region, params, &response)
	return
}
