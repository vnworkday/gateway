package tenant

import (
	"github.com/vnworkday/gateway/internal/common/model"
	"github.com/vnworkday/gateway/internal/common/util"
)

type Tenant struct {
	ID                      string `json:"id"`
	Name                    string `json:"name"`
	Status                  int    `json:"status"`
	Domain                  string `json:"domain"`
	Timezone                string `json:"timezone"`
	ProductionType          int    `json:"production_type"`
	SubscriptionType        int    `json:"subscription_type"`
	SelfRegistrationEnabled bool   `json:"self_registration_enabled"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
}

type CreateTenantRequest struct {
	Name                    string `json:"name"                      validate:"required,min=3,max=255"`
	Domain                  string `json:"domain"                    validate:"required,min=3,max=255"`
	Timezone                string `json:"timezone"                  validate:"required,timezone"`
	SubscriptionType        int    `json:"subscription_type"         validate:"required,gte=1"`
	SelfRegistrationEnabled bool   `json:"self_registration_enabled"`
}

type CreateTenantResponse struct {
	Code int     `json:"code"`
	Item *Tenant `json:"item"`
}

type UpdateTenantRequest struct {
	ID                      string `json:"id"                        validate:"required,uuidv4"`
	Name                    string `json:"name"                      validate:"required,min=3,max=255"`
	SubscriptionType        int    `json:"subscription_type"         validate:"required,gte=1"`
	SelfRegistrationEnabled bool   `json:"self_registration_enabled"`
	Status                  int    `json:"status"                    validate:"required,gte=1"`
}

type UpdateTenantResponse struct {
	Code int     `json:"code"`
	Item *Tenant `json:"item"`
}

type GetTenantRequest struct {
	ID string `json:"id" validate:"required,uuidv4"`
}

type GetTenantResponse struct {
	Code int     `json:"code"`
	Item *Tenant `json:"item"`
}

type ListTenantsRequest struct {
	Token string `json:"token" validate:"omitempty,base64"`
	Limit *int   `json:"limit" validate:"gte=1,lte=100"`

	Name                    string `json:"name"`
	Domain                  string `json:"domain"`
	Timezone                string `json:"timezone"                  validate:"timezone"`
	SubscriptionType        *int   `json:"subscription_type"         validate:"gte=1"`
	SelfRegistrationEnabled *bool  `json:"self_registration_enabled"`
	Status                  *int   `json:"status"                    validate:"gte=1"`

	//nolint:lll
	Sort  string `json:"order" validate:"omitempty,oneof=name domain timezone subscription_type self_registration_enabled status created_at updated_at"`
	Order string `json:"sort"  validate:"omitempty,oneof=asc desc"`
}

func DefaultListTenantsRequest() *ListTenantsRequest {
	return &ListTenantsRequest{
		Token:                   "",
		Limit:                   util.ToPointer(model.DefaultLimit),
		Name:                    "",
		Domain:                  "",
		Timezone:                "",
		SubscriptionType:        nil,
		SelfRegistrationEnabled: nil,
		Status:                  nil,
		Sort:                    model.DefaultSort,
		Order:                   model.DefaultOrder,
	}
}

type ListTenantsResponse struct {
	Code       int       `json:"code"`
	Items      []*Tenant `json:"items"`
	Next       string    `json:"next"`
	Previous   string    `json:"previous"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
}
