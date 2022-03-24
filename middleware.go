package tenantable

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

const (
	GormDBKey       = "GORM_DB"
	WrongSubdomain  = "Missing Tenant Subdomain."
	MissingTenantID = "TenantID is missing."
	TenantNotExists = "The tenant you are trying to access does not exist."
)

func GormInjector(config *Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// default will use the empty scope.
			c.Set(GormDBKey, config.DB)
			var tenantID string
			var result *gorm.DB
			var tenant Tenant

			if config.TenantKeyMethod == TenantByDomain { // Check tenant by subdomain
				host := c.Request().Host
				re := regexp.MustCompile(`^([^.]+).`)
				match := re.FindStringSubmatch(host)
				if len(match) == 0 {
					return c.String(http.StatusBadRequest, WrongSubdomain)
				}
				result = config.DB.Where("sub_domain = ?", match[1]).First(&tenant)
			} else if config.TenantKeyMethod == TenantByKey { // check tenant by key
				tenantID := c.Get(config.TenantKey)
				if tenantID == nil {
					return c.String(http.StatusBadRequest, MissingTenantID)
				}
				result = config.DB.Where("uuid = ?", tenantID).First(&tenant)
			}

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return c.String(http.StatusUnauthorized, TenantNotExists)
			} else if result.Error != nil {
				fmt.Println(result.Error)
				return c.String(http.StatusInternalServerError, "Some Error Happened during tenant checking.")
			}

			// Set Scoped Gorm object into echo context.
			c.Set(GormDBKey, config.DB.Set(config.TenantKey, tenantID))
			return next(c)
		}
	}
}
