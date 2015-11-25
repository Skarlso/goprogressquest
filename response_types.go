package main

//Message Basic Message format for greeting
type Message struct {
	Message string `json:"welcome"`
}

//NewCharacter The ID of a newly created character
type NewCharacter struct {
	CharacterID string `json:"id"`
}

//ErrorResponse A wrapped error response with proper message
type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}
