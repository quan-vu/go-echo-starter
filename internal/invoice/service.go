package invoice

import (
	"invoice-management/internal/models"
)

func FetchInvoices() ([]models.Invoice, error) {
	return GetInvoices()
}

func FetchInvoiceByID(id uint) (models.Invoice, error) {
	return GetInvoiceByID(id)
}

func AddInvoice(invoice *models.Invoice) error {
	return CreateInvoice(invoice)
}

func EditInvoice(invoice *models.Invoice) error {
	return UpdateInvoice(invoice)
}

func RemoveInvoice(id uint) error {
	return DeleteInvoice(id)
}
