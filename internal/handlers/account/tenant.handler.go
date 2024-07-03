package account

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type TenantHandler struct {
	logger *zap.Logger
}

type TenantHandlerParams struct {
	fx.In
	Logger *zap.Logger
}

func NewTenantHandler(params TenantHandlerParams) *TenantHandler {
	return &TenantHandler{
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
//	@Success		200	{object}	account.Tenant
//	@Failure		404	{object}	shared.Error	"when the tenant is not found"
//	@Failure		500	{object}	shared.Error	"when the server is unable to handle the request"
//	@Failure		400	{object}	shared.Error	"when the provided id is invalid"
//	@Failure		401	{object}	shared.Error	"when the user is not authenticated"
//	@Failure		403	{object}	shared.Error	"when the user is not authorized"
//	@Router			/tenants/{id} [get]
//	@Security		JWT
//	@Security		ApiKey
func (h *TenantHandler) GetTenant(ctx *fiber.Ctx) error {
	funcN := "GetTenant"

	id := ctx.Params("id")

	h.logger.Info("GetTenant", zap.String("id", id), zap.String("func", funcN))

	return ctx.JSON(fiber.Map{
		"status": "ok",
	})
}

// ListTenants godoc
//
//	@Id				ListTenants
//	@Summary		List all tenants
//	@Description	List all tenants
//	@Tags			Tenant
//	@Produce		json
//	@Success		200	{object}	account.ListTenantsResponse
//	@Failure		500	{object}	shared.Error	"when the server is unable to handle the request"
//	@Failure		401	{object}	shared.Error	"when the user is not authenticated"
//	@Failure		403	{object}	shared.Error	"when the user is not authorized"
//	@Router			/tenants [get]
//	@Security		JWT
func (h *TenantHandler) ListTenants(ctx *fiber.Ctx) error {
	funcN := "ListTenants"

	h.logger.Info("ListTenants", zap.String("func", funcN))

	return ctx.JSON(fiber.Map{
		"status": "ok",
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
//	@Success		200		{object}	account.Tenant
//	@Failure		500		{object}	shared.Error	"when the server is unable to handle the request"
//	@Failure		401		{object}	shared.Error	"when the user is not authenticated"
//	@Failure		403		{object}	shared.Error	"when the user is not authorized"
//	@Router			/tenants [post]
//	@Security		JWT
func (h *TenantHandler) CreateTenant(ctx *fiber.Ctx) error {
	funcN := "CreateTenant"

	h.logger.Info("CreateTenant", zap.String("func", funcN))

	return ctx.JSON(fiber.Map{
		"status": "ok",
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
//	@Success		200		{object}	account.Tenant
//	@Failure		500		{object}	shared.Error	"when the server is unable to handle the request"
//	@Failure		401		{object}	shared.Error	"when the user is not authenticated"
//	@Failure		403		{object}	shared.Error	"when the user is not authorized"
//	@Router			/tenants/{id} [put]
//	@Security		JWT
func (h *TenantHandler) UpdateTenant(ctx *fiber.Ctx) error {
	funcN := "UpdateTenant"

	id := ctx.Params("id")

	h.logger.Info("UpdateTenant", zap.String("id", id), zap.String("func", funcN))

	return ctx.JSON(fiber.Map{
		"status": "ok",
	})
}
