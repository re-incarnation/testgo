package stalcraftgo // import "github.com/re-incarnation/stalcraftgo/object"

// Error struct.
type Error struct {
	Errors []struct {
		Code    string            `json:"code"`
		Message string            `json:"message"`
		Params  map[string]string `json:"params,omitempty"`
	} `json:"errors"`
}
