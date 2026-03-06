package routes

import (
	"finpro/controller"
	"finpro/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authCtrl *controller.AuthController,
	assessCtrl *controller.AssessmentController,
	dashCtrl *controller.DashboardController,
	adminCtrl *controller.AdminController,
	testCtrl *controller.TestController,
) {
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		// Test Routes (Basic Task)
		test := api.Group("/test")
		{
			test.GET("/add", testCtrl.AddData)
			test.POST("/add", testCtrl.AddData)
			test.GET("/list", testCtrl.GetData)
		}

		// Auth Routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authCtrl.Register)
			auth.POST("/login", authCtrl.Login)
		}

		// Assessment Routes
		assessment := api.Group("/assessment")
		assessment.Use(middleware.AuthMiddleware())
		{
			assessment.GET("/questions", assessCtrl.GetQuestions)
			assessment.POST("/submit", assessCtrl.Submit)
		}

		// Dashboard Routes
		dashboard := api.Group("/dashboard")
		dashboard.Use(middleware.AuthMiddleware())
		{
			dashboard.GET("", dashCtrl.GetDashboard)
		}

		// Admin Routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			admin.GET("/users", adminCtrl.GetUsers)
			admin.GET("/evaluations", adminCtrl.GetEvaluations)
		}
	}
}
