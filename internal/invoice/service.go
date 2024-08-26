package invoice

func FetchInvoices() ([]Invoice, error) {
	return GetInvoices()
}

func FetchInvoiceByID(id uint) (Invoice, error) {
	return GetInvoiceByID(id)
}

func AddInvoice(invoice *Invoice) error {
	return CreateInvoice(invoice)
}

func EditInvoice(invoice *Invoice) error {
	return UpdateInvoice(invoice)
}

func RemoveInvoice(id uint) error {
	return DeleteInvoice(id)
}
