package stalcraftgo

type EmissionResponse struct {
	CurrentStart  string `json:"current_start"`
	PreviousStart string `json:"previous_start"`
	PreviousEnd   string `json:"previous_end"`
}
