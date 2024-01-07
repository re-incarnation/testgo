package stalcraftgo

import stalcraftgo "github.com/re-incarnation/stalcraftgo/object"

func (exbo *EXBO) ListOfRegions(params Params) (response stalcraftgo.RegionsResponse, err error) {
	//Used without token
	err = exbo.RequestUnmarshalRegions(params, &response)
	return
}
