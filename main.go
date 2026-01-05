package main

import (
	"invoice-payment-system/config"
	"invoice-payment-system/handler/company"
	"invoice-payment-system/handler/invoice"
	model2 "invoice-payment-system/model"
	"invoice-payment-system/repository/company_read"
	"invoice-payment-system/repository/company_write"
	"invoice-payment-system/repository/invoice_read"
	"invoice-payment-system/repository/invoice_write"
	"invoice-payment-system/usecase/company_command"
	"invoice-payment-system/usecase/company_query"
	"invoice-payment-system/usecase/invoice_command"
	"invoice-payment-system/usecase/invoice_query"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	config.LoadEnv()

	dsn := config.BuildDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB:", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(
		&model2.Invoices{},
		&model2.Company{},
		&model2.Item{},
	); err != nil {
		log.Fatal("failed to auto migrate:", err)
	}

	log.Println("database connected & migrated")
	return db
}

func main() {
	db := initDB()
	invoiceReadRepo := invoice_read.NewInvoiceReadRepo(db)
	invoiceWriteRepo := invoice_write.NewInvoiceWriteRepo(db)

	createInvoiceCmd := &invoice_command.CreateInvoiceUsecase{
		DB:   db,
		Repo: invoiceWriteRepo,
	}

	submitInvoiceCmd := &invoice_command.SubmiteInvoiceUsecase{
		DB:   db,
		Repo: invoiceWriteRepo,
	}

	approveInvoiceCmd := &invoice_command.ApproveInvoiceUsecase{
		DB:   db,
		Repo: invoiceWriteRepo,
	}

	payInvoiceCmd := &invoice_command.PayInvoiceUsecase{
		DB:   db,
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

	companyWriteRepo := company_write.NewCompanyWriteRepo(db)
	companyReadRepo := company_read.NewCompanyReadRepo(db)

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
