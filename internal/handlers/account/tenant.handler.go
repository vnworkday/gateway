package accounthandler

import (
	"errors"

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
//	@Summary		Get a tenant by ID
//	@Description	Get a tenant by ID
//	@Tags			Tenant
//	@Produce		json
//	@Param			id	path		string	true	"Tenant ID"
//	@Success		200	{object}	accountmodel.Tenant
//	@Router			/tenants/{id} [get]
func (h *TenantHandler) GetTenant(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	h.logger.Info("GetTenant", zap.String("id", id))

	return errors.Join(ctx.JSON(fiber.Map{
		"status": "ok",
	}))
}

// ListTenants godoc
//
//	@Summary		List all tenants
//	@Description	List all tenants
//	@Tags			Tenant
//	@Produce		json
//	@Success		200	{object}	accountmodel.Tenant
//	@Router			/tenants [get]
func (h *TenantHandler) ListTenants(ctx *fiber.Ctx) error {
	h.logger.Info("ListTenants")

	return errors.Join(ctx.JSON(fiber.Map{
		"status": "ok",
	}))
}

func (h *TenantHandler) CreateTenant(ctx *fiber.Ctx) error {
	h.logger.Info("CreateTenant")

	return errors.Join(ctx.JSON(fiber.Map{
		"status": "ok",
	}))
}

func (h *TenantHandler) UpdateTenant(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	h.logger.Info("UpdateTenant", zap.String("id", id))

	return errors.Join(ctx.JSON(fiber.Map{
		"status": "ok",
	}))
}
