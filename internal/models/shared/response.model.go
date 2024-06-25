package shared

// ResponsePagination godoc
// @Description Represents a pagination response.
type ResponsePagination struct {
	// NextToken is a token that can be used to retrieve the next page of results.
	// If this field is not present, it means that there are no more results to retrieve.
	NextToken string `example:"Vk4gV29ya2RheSBFeGFtcGxl" format:"base64" json:"nextToken" validate:"optional"`
	// PreviousToken a token that can be used to retrieve the previous page of results.
	// If this field is not present, it means that there are no more results to retrieve.
	PreviousToken string `example:"Vk4gV29ya2RheSBFeGFtcGxl" format:"base64" json:"previousToken" validate:"optional"`
	// Total is the total number of items in the list.
	Total uint `example:"100" json:"total" minimum:"0" validate:"required"`
	// TotalPages is the total number of pages in the list.
	TotalPages uint `example:"10" json:"totalPages" minimum:"0" validate:"required"`
}
