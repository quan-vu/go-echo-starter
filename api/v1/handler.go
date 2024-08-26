package v1

import (
	"invoice-management/internal/invoice"
	"invoice-management/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetInvoices(c echo.Context) error {
	invoices, err := invoice.FetchInvoices()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch invoices"})
	}
	return c.JSON(http.StatusOK, invoices)
}

func GetInvoice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	invoice, err := invoice.FetchInvoiceByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invoice not found"})
	}
	return c.JSON(http.StatusOK, invoice)
}

func CreateInvoice(c echo.Context) error {
	var inv models.Invoice
	if err := c.Bind(&inv); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := invoice.AddInvoice(&inv); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, inv)
}

func UpdateInvoice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	inv, err := invoice.FetchInvoiceByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invoice not found"})
	}

	if err := c.Bind(&inv); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := invoice.EditInvoice(&inv); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update invoice"})
	}
	return c.JSON(http.StatusOK, inv)
}

func DeleteInvoice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := invoice.RemoveInvoice(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete invoice"})
	}
	return c.NoContent(http.StatusNoContent)
}
