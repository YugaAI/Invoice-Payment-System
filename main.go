package main

import (
	"context"
	"invoice-payment-system/config"
	"invoice-payment-system/handler/company"
	"invoice-payment-system/handler/invoice"
	"invoice-payment-system/handler/user"
	"invoice-payment-system/redis_client"
	"invoice-payment-system/repository/companyRepo/company_read"
	"invoice-payment-system/repository/companyRepo/company_write"
	"invoice-payment-system/repository/invoiceRepo/invoice_read"
	"invoice-payment-system/repository/invoiceRepo/invoice_write"
	"invoice-payment-system/repository/user/user_read"
	"invoice-payment-system/repository/user/user_write"
	"invoice-payment-system/usecase/companyUsecase/company_read_usecase"
	"invoice-payment-system/usecase/companyUsecase/company_write_usecase"
	"invoice-payment-system/usecase/invoiceUsecase/invoice_read_usecase"
	"invoice-payment-system/usecase/invoiceUsecase/invoice_write_usecase"
	"invoice-payment-system/usecase/user/user_read_usecase"
	"invoice-payment-system/usecase/user/user_write_usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	writeDB := config.InitWriteDB()
	readDB := config.InitReadDB()
	redisCfg := config.LoadRedisConfig()
	redisClient := redis_client.NewRedisClient(*redisCfg)
	ctx := context.Background()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	invoiceReadRepo := invoice_read.NewInvoiceReadRepo(readDB, redisClient, ctx)
	invoiceWriteRepo := invoice_write.NewInvoiceWriteRepo(writeDB)
	invoiceReadUC := invoice_read_usecase.NewInvoiceReadUsecase(invoiceReadRepo)
	invoiceWriteUC := invoice_write_usecase.NewInvoiceWriteUsecase(writeDB, invoiceWriteRepo)
	invoiceHandler := invoice.NewInvoiceHandler(r, invoiceReadUC, invoiceWriteUC)
	invoiceHandler.RegisterInvoiceRoutes()

	companyWriteRepo := company_write.NewCompanyWriteRepo(writeDB)
	companyReadRepo := company_read.NewCompanyReadRepo(readDB, redisClient, ctx)
	companyWriteUC := company_write_usecase.NewCompanyWriteUsecase(writeDB, companyWriteRepo)
	companyReadUC := company_read_usecase.NewCompanyReadUsecase(companyReadRepo)
	companyHandler := company.NewCompanyHandler(r, companyReadUC, companyWriteUC)
	companyHandler.RegisterCompanyRoutes()

	userReadRepo := user_read.NewLogin(readDB, redisClient, ctx)
	userWriteRepo := user_write.NewSignIn(writeDB)
	userReadUC := user_read_usecase.NewLoginUsecase(userReadRepo)
	userWriteUC := user_write_usecase.NewWriteUsecase(writeDB, userWriteRepo)
	userHandler := user.NewUserHandler(r, userReadUC, userWriteUC)
	userHandler.RegisterUserRoutes()

	r.Run(":8080")

}
