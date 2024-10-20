package invoice

// import (
// 	"instalasi-pro/database"
// 	helper "instalasi-pro/helpers"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func Initiator(router *gin.Engine) {
// 	api := router.Group("/api/invoices")
// 	{
// 		api.GET("/", FindAll)
// 		api.GET("/:id", FindById)
// 	}
// }

// func FindAll(c *gin.Context) {
// 	invoices, err := NewService(NewRepository(database.DB)).FindAll()
// 	if err != nil {
// 		response := helper.APIResponse("Failed to get invoices", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("List of invoices", http.StatusOK, "success", invoices)
// 	c.JSON(http.StatusOK, response)

// }

// func FindById(c *gin.Context) {
// 	invoiceID := c.Param("id")

// 	id, err := strconv.Atoi(invoiceID)
// 	if err != nil {
// 		response := helper.APIResponse("Failed to get invoice", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	invoice, err := NewService(NewRepository(database.DB)).GetInvoice(id)
// 	if err != nil {
// 		response := helper.APIResponse("Failed to get invoice", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Detail of invoice", http.StatusOK, "success", invoice)
// 	c.JSON(http.StatusOK, response)
// }
