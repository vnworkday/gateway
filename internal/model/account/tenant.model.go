package account

import "github.com/vnworkday/gateway/internal/model/shared"

// Tenant godoc
// @Description represents a tenant.
type Tenant struct {
	// ID of the tenant
	ID string `example:"abcxyz" json:"id" validate:"required"`
	// Name of the tenant
	Name string `example:"Tenant Name" json:"name" validate:"required"`
}

// ListTenantsResponse godoc
// @Description represents a list of tenants with pagination information in the response.
type ListTenantsResponse struct {
	shared.ResponsePagination
	// Items is a list of tenants
	Items []Tenant `json:"items" minItems:"0" validate:"required"`
}
