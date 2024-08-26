package invoice

import (
	"invoice-management/internal/db"
)

func CreateInvoice(invoice *Invoice) error {
	return db.DB.Create(invoice).Error
}

func GetInvoices() ([]Invoice, error) {
	var invoices []Invoice
	result := db.DB.Find(&invoices)
	return invoices, result.Error
}

func GetInvoiceByID(id uint) (Invoice, error) {
	var invoice Invoice
	result := db.DB.First(&invoice, id)
	return invoice, result.Error
}

func UpdateInvoice(invoice *Invoice) error {
	return db.DB.Save(invoice).Error
}

func DeleteInvoice(id uint) error {
	return db.DB.Delete(&Invoice{}, id).Error
}
