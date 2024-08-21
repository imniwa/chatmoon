package model

type ChatSocketPayload struct {
	Type    string `json:"type,omitempty"`
	From    string `json:"from,omitempty"`
	RoomID  string `json:"room_id,omitempty"`
	Message string `json:"message,omitempty"`
}

type ChatSocketResponse struct {
	Type    string `json:"type,omitempty"`
	From    string `json:"from,omitempty"`
	RoomID  string `json:"room_id,omitempty"`
	Message string `json:"message,omitempty"`
}
