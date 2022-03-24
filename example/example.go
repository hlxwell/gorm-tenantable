package main

import (
	"fmt"
	"net/http"

	tenantable "github.com/hlxwell/gorm-tenantable"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DBUser     = "root"
	DBPassword = ""
	DBAddress  = "localhost"
	DBPort     = "3306"
	DBName     = "gorm_by_example"
	TenantKey  = "tenant_id"
)

var Conn *gorm.DB

func init() {
	setupConn()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(TenantKey, "E638D8DC-4C92-4B43-A9FD-FB32EF6369F6")
			return next(c)
		}
	})

	e.Use(tenantable.GormInjector(&tenantable.Config{
		TenantKey:   TenantKey,
		DB:          Conn,
		AutoMigrate: true,
	}))

	e.GET("/content", func(c echo.Context) error {
		return c.String(http.StatusOK, "Content!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

// Helper Methods ============================

func setupConn() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		DBUser, DBPassword, DBAddress, DBPort, DBName,
	)

	var err error
	if Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}

	if err = Conn.AutoMigrate(&tenantable.Tenant{}); err != nil {
		panic(err)
	}
}
