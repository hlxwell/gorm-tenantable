# GORM Tenantable

The purpose of Gorm-Tenantable project for solving multi-tenancy requirement in echo framework based project.  

## Features

- Be able to automatically inject a scoped GORM database object into your echo.context. 
- Be able to recognize which tenant of current logged in user. 

## How to use

### Add middleware to your project.

```go
e.Use(tenantable.GormInjector(&tenantable.Config{
  TenantKey:   "tenant_key",
  DB:          Conn,
  AutoMigrate: false,
}))
```

## Example Code

You could run `go run example/example.go` to try the example. But make sure to change the database config.
