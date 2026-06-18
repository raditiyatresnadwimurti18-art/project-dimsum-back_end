package routes

import (
	"project-dimsum/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Memanggil fungsi dari folder controllers
	r.POST("/api/dimsum", controllers.TambahDimsum)
	r.GET("/api/dimsum", controllers.GetSemuaDimsum)
	r.PUT("/api/dimsum/:id", controllers.UpdateDimsum)
	r.DELETE("/api/dimsum/:id", controllers.HapusDimsum)
}
