package response

type CreateFlightResponse struct {
	ID             string `json:"id`
	Name           string `json:"name"`
	From           string `json:"from"`
	To             string `json:"to"`
	Date           string `json:"date"`
	Status         string `json:"status"`
	Available_slot int    `json:"available_slot"`
}

type UpdateFlightResponse struct {
	ID             string `json:"id`
	Name           string `json:"name"`
	From           string `json:"from"`
	To             string `json:"to"`
	Date           string `json:"date"`
	Status         string `json:"status"`
	Available_slot int    `json:"available_slot"`
}
type SearchFlightResponse struct {
	ID             string `json:"id`
	Name           string `json:"name"`
	From           string `json:"from"`
	To             string `json:"to"`
	Date           string `json:"date"`
	Status         string `json:"status"`
	Available_slot int    `json:"available_slot"`
}
