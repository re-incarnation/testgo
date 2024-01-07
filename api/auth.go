package stalcraftgo

import stalcraftgo "github.com/re-incarnation/stalcraftgo/object"

func (exbo *EXBO) Auth(params Params) (response stalcraftgo.AuthResponse, err error) {
	//Used Client Id and Client Secret
	err = exbo.RequestUnmarshalAuth(params, &response)
	return
}
