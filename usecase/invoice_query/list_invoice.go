package invoice_query

type ListInvoiceUsecase struct {
	Repo InvoiceReadRepoInterface
}

func (h *ListInvoiceUsecase) Execute(q ListInvoiceQuery) (*ListInvoiceResult, error) {
	items, total, err := h.Repo.List(q.CompanyID, q.Status, q.Limit, q.Offset)
	if err != nil {
		return nil, err
	}

	return &ListInvoiceResult{
		Data:  items,
		Total: total,
	}, nil
}
