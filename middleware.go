package tenantable

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

const (
	GormDBKey = "GORM_DB"
)

func GormInjector(config *Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// default will use the empty scope.
			c.Set(GormDBKey, config.DB)

			tenantID := c.Get(config.TenantKey)
			if tenantID == nil {
				c.String(http.StatusBadRequest, "TenantID is missing.")
				return nil
			}

			// Check if tenantID existed in the tenants table.
			var tenant Tenant
			result := config.DB.Where("uuid = ?", tenantID).First(&tenant)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.String(http.StatusUnauthorized, "The tenant you are trying to access does not exist.")
				return nil
			}

			// Set Scoped Gorm object into echo context.
			c.Set(GormDBKey, config.DB.Set(config.TenantKey, tenantID))
			return next(c)
		}
	}
}
