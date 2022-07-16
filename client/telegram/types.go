package telegram

type UpdatesResponse struct {
	Ok     bool      `json:"ok"`
	Result []Updates `json:"result"`
}

type Updates struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}
