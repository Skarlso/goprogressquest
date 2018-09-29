package responsetypes

//Message Basic Message format for greeting
type Message struct {
	Message string `json:"message"`
}

//ErrorResponse A wrapped error response with proper message
type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

//LookingForID encloses the ID which is searched for when looking for a Player
type LookingForID struct {
	ID string `json:"id"`
}
