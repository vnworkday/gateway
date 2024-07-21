package tenant

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Handler struct {
	logger       *zap.Logger
	tenantClient *GRPCClient
}

type HandlerParams struct {
	fx.In
	Logger           *zap.Logger
	TenantGRPCClient *GRPCClient
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		logger:       params.Logger.Named("tenant"),
		tenantClient: params.TenantGRPCClient,
	}
}

// GetTenant godoc
//
//	@Id				GetTenant
//	@Summary		Get a tenant
//	@Description	Get a tenant by ID
//	@Tags			Tenant
//	@Produce		json
//	@Param			id	path		string	true	"Tenant ID"
//	@Success		200	{object}	GetTenantResponse
//	@Failure		404	{object}	model.Error	"when the tenant is not found"
//	@Failure		500	{object}	model.Error	"when the server is unable to handle the request"
//	@Failure		400	{object}	model.Error	"when the provided id is invalid"
//	@Failure		401	{object}	model.Error	"when the user is not authenticated"
//	@Failure		403	{object}	model.Error	"when the user is not authorized"
//	@Router			/tenants/{id} [get]
//	@Security		JWT
//	@Security		ApiKey
func (h *Handler) GetTenant(ctx *fiber.Ctx) error {
	var request *GetTenantRequest

	if err := ctx.ParamsParser(request); err != nil {
		return errors.Wrap(err, "handler: failed to parse request")
	}

	response, err := h.tenantClient.GetTenant(
		ctx.Context(),
		request,
		ToGetTenantRequest,
		ToGetTenantResponse,
	)
	if err != nil {
		return errors.Wrap(err, "handler: failed to get tenant")
	}

	return ctx.JSON(response)
}

// ListTenants godoc
//
//	@Id				ListTenants
//	@Summary		List all tenants
//	@Description	List all tenants
//	@Tags			Tenant
//	@Produce		json
//	@Success		200	{object}	ListTenantsResponse
//	@Failure		500	{object}	model.Error	"when the server is unable to handle the request"
//	@Failure		401	{object}	model.Error	"when the user is not authenticated"
//	@Failure		403	{object}	model.Error	"when the user is not authorized"
//	@Router			/tenants [get]
//	@Security		JWT
func (h *Handler) ListTenants(ctx *fiber.Ctx) error {
	var request *ListTenantsRequest

	if err := ctx.QueryParser(request); len(ctx.Queries()) > 0 && err != nil {
		return errors.Wrap(err, "handler: failed to parse request")
	}

	response, err := h.tenantClient.ListTenants(
		ctx.Context(),
		request,
		ToListTenantsRequest,
		ToListTenantsResponse,
	)
	if err != nil {
		return errors.Wrap(err, "handler: failed to list tenants")
	}

	return ctx.JSON(response)
}

// CreateTenant godoc
//
//	@Id				CreateTenant
//	@Summary		Create a tenant
//	@Description	Create a tenant
//	@Tags			Tenant
//	@Accept			json
//	@Produce		json
//	@Param			tenant	body		account.Tenant	true	"Tenant object that needs to be created"
//	@Success		200		{object}	CreateTenantResponse
//	@Failure		500		{object}	model.Error	"when the server is unable to handle the request"
//	@Failure		401		{object}	model.Error	"when the user is not authenticated"
//	@Failure		403		{object}	model.Error	"when the user is not authorized"
//	@Router			/tenants [post]
//	@Security		JWT
func (h *Handler) CreateTenant(ctx *fiber.Ctx) error {
	var request *CreateTenantRequest

	if err := ctx.BodyParser(request); err != nil {
		return errors.Wrap(err, "handler: failed to parse request")
	}

	response, err := h.tenantClient.CreateTenant(
		ctx.Context(),
		request,
		ToCreateTenantRequest,
		ToCreateTenantResponse,
	)
	if err != nil {
		return errors.Wrap(err, "handler: failed to create tenant")
	}

	return ctx.JSON(response)
}

// UpdateTenant godoc
//
//	@Id				UpdateTenant
//	@Summary		Update a tenant
//	@Description	Update an existing tenant by ID
//	@Tags			Tenant
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Tenant ID"
//	@Param			tenant	body		account.Tenant	true	"Tenant object that needs to be updated"
//	@Success		200		{object}	UpdateTenantResponse
//	@Failure		500		{object}	model.Error	"when the server is unable to handle the request"
//	@Failure		401		{object}	model.Error	"when the user is not authenticated"
//	@Failure		403		{object}	model.Error	"when the user is not authorized"
//	@Router			/tenants/{id} [put]
//	@Security		JWT
func (h *Handler) UpdateTenant(ctx *fiber.Ctx) error {
	var request *UpdateTenantRequest

	if err := ctx.BodyParser(request); err != nil {
		return errors.Wrap(err, "handler: failed to parse request")
	}

	response, err := h.tenantClient.UpdateTenant(
		ctx.Context(),
		request,
		ToUpdateTenantRequest,
		ToUpdateTenantResponse,
	)
	if err != nil {
		return errors.Wrap(err, "handler: failed to update tenant")
	}

	return ctx.JSON(response)
}
