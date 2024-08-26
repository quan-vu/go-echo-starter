package v1

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("/invoices", GetInvoices)
	v1.GET("/invoices/:id", GetInvoice)
	v1.POST("/invoices", CreateInvoice)
	v1.PUT("/invoices/:id", UpdateInvoice)
	v1.DELETE("/invoices/:id", DeleteInvoice)
}
