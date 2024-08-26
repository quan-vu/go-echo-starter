package invoice

import (
	"invoice-management/internal/db"
	"invoice-management/internal/models"
)

func CreateInvoice(invoice *models.Invoice) error {
	return db.DB.Create(invoice).Error
}

func GetInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice
	result := db.DB.Find(&invoices)
	return invoices, result.Error
}

func GetInvoiceByID(id uint) (models.Invoice, error) {
	var invoice models.Invoice
	result := db.DB.First(&invoice, id)
	return invoice, result.Error
}

func UpdateInvoice(invoice *models.Invoice) error {
	return db.DB.Save(invoice).Error
}

func DeleteInvoice(id uint) error {
	return db.DB.Delete(&models.Invoice{}, id).Error
}
