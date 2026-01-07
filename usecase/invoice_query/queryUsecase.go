package invoice_query

import "invoice-payment-system/dto"

type InvoiceReadRepoInterface interface {
	FindDetailByID(id uint64) (*dto.InvoiceDetail, error)
	GetDashboard(companyID uint64) (*dto.InvoiceDashboard, error)
	List(companyID uint64, page, limit int) ([]dto.InvoiceList, error)
}

//type QueryUsecase struct {
//	Repo InvoiceReadRepoInterface
//}
//
//func NewQueryUsecase(repo InvoiceReadRepoInterface) *QueryUsecase {
//	return &QueryUsecase{
//		Repo: repo,
//	}
//}
