package main

import (
	"invoice-payment-system/config"
	"invoice-payment-system/handler/company"
	"invoice-payment-system/handler/invoice"
	"invoice-payment-system/repository/company_read"
	"invoice-payment-system/repository/company_write"
	"invoice-payment-system/repository/invoice_read"
	"invoice-payment-system/repository/invoice_write"
	"invoice-payment-system/usecase/company_command"
	"invoice-payment-system/usecase/company_query"
	"invoice-payment-system/usecase/invoice_command"
	"invoice-payment-system/usecase/invoice_query"

	"github.com/gin-gonic/gin"
)

func main() {
	writeDB := config.InitWriteDB()
	readDB := config.InitReadDB()
	invoiceReadRepo := invoice_read.NewInvoiceReadRepo(readDB)
	invoiceWriteRepo := invoice_write.NewInvoiceWriteRepo(writeDB)

	createInvoiceCmd := &invoice_command.CreateInvoiceUsecase{
		DB:   writeDB,
		Repo: invoiceWriteRepo,
	}

	submitInvoiceCmd := &invoice_command.SubmiteInvoiceUsecase{
		DB:   writeDB,
		Repo: invoiceWriteRepo,
	}

	approveInvoiceCmd := &invoice_command.ApproveInvoiceUsecase{
		DB:   writeDB,
		Repo: invoiceWriteRepo,
	}

	payInvoiceCmd := &invoice_command.PayInvoiceUsecase{
		DB:   writeDB,
		Repo: invoiceWriteRepo,
	}

	getInvoiceDetailQuery := &invoice_query.GetInvoiceDetailUsecase{
		Repo: invoiceReadRepo,
	}

	listInvoiceQuery := &invoice_query.ListInvoiceUsecase{
		Repo: invoiceReadRepo,
	}

	invoiceDashboardQuery := &invoice_query.InvoiceDashboardUsecase{
		Repo: invoiceReadRepo,
	}
	invooiceHandler := &invoice.InvoiceHandler{
		CreateCmd:  createInvoiceCmd,
		SubmitCmd:  submitInvoiceCmd,
		ApproveCmd: approveInvoiceCmd,
		PayCmd:     payInvoiceCmd,

		DetailQuery:    getInvoiceDetailQuery,
		ListQuery:      listInvoiceQuery,
		DashboardQuery: invoiceDashboardQuery,
	}

	companyWriteRepo := company_write.NewCompanyWriteRepo(writeDB)
	companyReadRepo := company_read.NewCompanyReadRepo(readDB)

	companyWrite := &company_command.CreateCompanyUsecase{
		Repo: companyWriteRepo,
	}

	getCompanyQuery := &company_query.GetCompanyUsecase{
		Repo: companyReadRepo,
	}
	listCompanyQuery := &company_query.ListCompanyUsecase{
		Repo: companyReadRepo,
	}
	companyHandler := &company.CompanyHandler{
		CreateCmd: companyWrite,
		GetQuery:  getCompanyQuery,
		ListQuery: listCompanyQuery,
	}

	r := gin.Default()
	company.NewRouterCompany(r, companyHandler)
	invoice.NewRouterInvoice(r, invooiceHandler)
	r.Run(":8080")

}
