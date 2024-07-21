package tenant

import (
	"time"

	tenantv1 "buf.build/gen/go/ntduycs/vnworkday/protocolbuffers/go/account/tenant/v1"
	sharedv1 "buf.build/gen/go/ntduycs/vnworkday/protocolbuffers/go/shared/v1"
	"github.com/gookit/goutil/arrutil"
)

func ToListTenantsRequest(from *ListTenantsRequest) *tenantv1.ListTenantsRequest {
	if from == nil {
		from = DefaultListTenantsRequest()
	}

	return &tenantv1.ListTenantsRequest{
		Pagination: &sharedv1.RequestPagination{
			Token: from.Token,
			Limit: int32(*from.Limit),
		},
		Filters: []*sharedv1.RequestFilter{
			{
				Field:           "name",
				Value:           from.Name,
				Operator:        sharedv1.Operator_OPERATOR_EQ,
				IsCaseSensitive: true,
				ValueType:       sharedv1.ValueType_VALUE_TYPE_STRING,
			},
		},
		Sorts: []*sharedv1.RequestSort{
			{
				Field: from.Sort,
				Order: from.Order,
			},
		},
	}
}

func ToListTenantsResponse(from *tenantv1.ListTenantsResponse) *ListTenantsResponse {
	return &ListTenantsResponse{
		Code: 0,
		Items: arrutil.Map[*tenantv1.Tenant, *Tenant](
			from.GetTenants(),
			func(input *tenantv1.Tenant) (target *Tenant, find bool) {
				return toEntity(input), true
			},
		),
		Next:       from.GetPagination().GetNextToken(),
		Previous:   from.GetPagination().GetPreviousToken(),
		Total:      int(from.GetPagination().GetTotal()),
		TotalPages: int(from.GetPagination().GetTotalPages()),
	}
}

func ToGetTenantRequest(from *GetTenantRequest) *tenantv1.GetTenantRequest {
	return &tenantv1.GetTenantRequest{
		Id: from.ID,
	}
}

func ToGetTenantResponse(from *tenantv1.GetTenantResponse) *GetTenantResponse {
	return &GetTenantResponse{
		Code: 0,
		Item: toEntity(from.GetTenant()),
	}
}

func ToCreateTenantRequest(from *CreateTenantRequest) *tenantv1.CreateTenantRequest {
	return &tenantv1.CreateTenantRequest{
		Info:                    nil,
		Name:                    from.Name,
		Domain:                  from.Domain,
		Timezone:                from.Timezone,
		SubscriptionType:        tenantv1.TenantSubscriptionType(from.SubscriptionType),
		SelfRegistrationEnabled: from.SelfRegistrationEnabled,
	}
}

func ToCreateTenantResponse(from *tenantv1.CreateTenantResponse) *CreateTenantResponse {
	return &CreateTenantResponse{
		Code: 0,
		Item: toEntity(from.GetTenant()),
	}
}

func ToUpdateTenantRequest(from *UpdateTenantRequest) *tenantv1.UpdateTenantRequest {
	return &tenantv1.UpdateTenantRequest{
		Info:                    nil,
		Id:                      from.ID,
		Name:                    from.Name,
		SubscriptionType:        tenantv1.TenantSubscriptionType(from.SubscriptionType),
		SelfRegistrationEnabled: from.SelfRegistrationEnabled,
	}
}

func ToUpdateTenantResponse(from *tenantv1.UpdateTenantResponse) *UpdateTenantResponse {
	return &UpdateTenantResponse{
		Code: 0,
		Item: toEntity(from.GetTenant()),
	}
}

func toEntity(from *tenantv1.Tenant) *Tenant {
	return &Tenant{
		ID:                      from.GetId(),
		Name:                    from.GetName(),
		Status:                  int(from.GetStatus()),
		Domain:                  from.GetDomain(),
		Timezone:                from.GetTimezone(),
		ProductionType:          int(from.GetProductionType()),
		SubscriptionType:        int(from.GetSubscriptionType()),
		SelfRegistrationEnabled: from.GetSelfRegistrationEnabled(),
		CreatedAt:               from.GetCreatedAt().AsTime().Format(time.RFC3339),
		UpdatedAt:               from.GetUpdatedAt().AsTime().Format(time.RFC3339),
	}
}
