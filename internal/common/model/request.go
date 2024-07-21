package model

const (
	DefaultLimit = 10
	DefaultSort  = "created_at"
	DefaultOrder = "desc"
)

// RequestPagination godoc
// @Description Represents a pagination request.
type RequestPagination struct {
	// Token is a base64 string that can be used to retrieve the next page of results.
	// If not present, retrieve the first page.
	Token string `default:"" example:"Vk4gV29ya2RheSBFeGFtcGxl" format:"base64" json:"token" validate:"optional"`
	// Limit is the maximum number of items to return in the response.
	// The default value is 10.
	Limit uint `default:"10" example:"10" json:"limit" maximum:"100" minimum:"1" validate:"required"`
}
