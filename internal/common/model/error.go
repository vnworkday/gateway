package model

// Error model
// @Description Represents a client-facing API error.
type Error struct {
	// Code should be unique and identifiable. It is used to determine the type of error.
	Code Code `example:"1000" json:"code" validate:"required"`
	// Title is a short, human-readable title of the error. It is used to determine the type of error.
	Title string `example:"Internal Server Error" json:"title" validate:"optional"`
	// Message is a human-readable description of the error. Should not be used to display to the user.
	Message string `example:"An unexpected error occurred. Please try again later." json:"message" validate:"optional"`
}

func NewError(code Code, message ...string) Error {
	msg := "Something went wrong. Please try again later."

	if len(message) > 0 {
		msg = message[0]
	}

	return Error{
		Code:    code,
		Title:   CodeToMessage[code],
		Message: msg,
	}
}
