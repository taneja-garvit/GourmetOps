package routes

import (
	"GourmetOps/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(app *gin.Engine) {
	app.GET("/invoices", controllers.GetInvoices())
	app.GET("/invoices/:invoice_id", controllers.GetInvoice())
	app.POST("/invoices", controllers.CreateInvoice())
	app.PATCH("invoices/:invoice_id", controllers.UpdateInvoice())
}
