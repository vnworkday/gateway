package tenant

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
}

type HandlerParams struct {
	fx.In
	Logger *zap.Logger
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		logger: params.Logger.Named("tenant"),
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
	return ctx.JSON(fiber.Map{
		"message": "get tenant",
	})
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
	return ctx.JSON(fiber.Map{
		"message": "list tenants",
	})
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
	return ctx.JSON(fiber.Map{
		"message": "create tenant",
	})
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
	return ctx.JSON(fiber.Map{
		"message": "update tenant",
	})
}
