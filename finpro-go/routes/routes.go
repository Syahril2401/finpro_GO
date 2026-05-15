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
	noteCtrl controller.NoteController,
	plannerCtrl *controller.PlannerController,
	targetCtrl *controller.TargetController,
	workspaceCtrl *controller.WorkspaceController,
) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		// Test Routes
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
			assessment.GET("/status", assessCtrl.GetStatus)
			assessment.POST("/chat", assessCtrl.Chat)
		}

		// Dashboard Routes
		dashboard := api.Group("/dashboard")
		dashboard.Use(middleware.AuthMiddleware())
		{
			// Dashboard aggregation (real metrics)
			dashboard.GET("", workspaceCtrl.GetDashboardMetrics)

			// Search
			dashboard.GET("/search", workspaceCtrl.Search)

			// Notifications
			dashboard.GET("/notifications", workspaceCtrl.GetNotifications)
			dashboard.PUT("/notifications/:id/read", workspaceCtrl.MarkNotificationRead)
			dashboard.PUT("/notifications/read-all", workspaceCtrl.MarkAllNotificationsRead)

			// Progress
			dashboard.GET("/progress", workspaceCtrl.GetProgress)

			// Planner CRUD
			dashboard.GET("/planner", plannerCtrl.GetSessions)
			dashboard.POST("/planner", plannerCtrl.CreateSession)
			dashboard.GET("/planner/:id", plannerCtrl.GetSession)
			dashboard.PUT("/planner/:id", plannerCtrl.UpdateSession)
			dashboard.DELETE("/planner/:id", plannerCtrl.DeleteSession)
			dashboard.PATCH("/planner/:id/complete", plannerCtrl.CompleteSession)
			dashboard.PATCH("/planner/:id/skip", plannerCtrl.SkipSession)

			// Weekly Targets CRUD
			dashboard.GET("/weekly-targets", targetCtrl.GetTargets)
			dashboard.POST("/weekly-targets", targetCtrl.CreateTarget)
			dashboard.GET("/weekly-targets/:id", targetCtrl.GetTarget)
			dashboard.PUT("/weekly-targets/:id", targetCtrl.UpdateTarget)
			dashboard.DELETE("/weekly-targets/:id", targetCtrl.DeleteTarget)
			dashboard.POST("/weekly-targets/:id/subtasks", targetCtrl.CreateSubtask)
			dashboard.PUT("/weekly-targets/:id/subtasks/:subtaskId", targetCtrl.UpdateSubtask)
			dashboard.DELETE("/weekly-targets/:id/subtasks/:subtaskId", targetCtrl.DeleteSubtask)
			dashboard.PATCH("/weekly-targets/:id/subtasks/:subtaskId/toggle", targetCtrl.ToggleSubtask)

			// Notes CRUD
			dashboard.GET("/notes", noteCtrl.GetNotes)
			dashboard.POST("/notes", noteCtrl.CreateNote)
			dashboard.GET("/notes/search", noteCtrl.SearchNotes)
			dashboard.GET("/notes/templates", noteCtrl.GetTemplates)
			dashboard.GET("/notes/:id", noteCtrl.GetNoteByID)
			dashboard.PUT("/notes/:id", noteCtrl.UpdateNote)
			dashboard.DELETE("/notes/:id", noteCtrl.DeleteNote)
			dashboard.PUT("/notes/:id/pin", noteCtrl.TogglePin)
			dashboard.PUT("/notes/:id/archive", noteCtrl.ToggleArchive)

			// Legacy stubs (keeping for compatibility)
			dashboard.GET("/ai-strategies", dashCtrl.GetAIStrategies)
			dashboard.GET("/settings", dashCtrl.GetSettings)
		}

		// Admin Routes
		api.POST("/admin/login", adminCtrl.Login)
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			admin.GET("/users", adminCtrl.GetUsers)
			admin.GET("/evaluations", adminCtrl.GetEvaluations)
		}
	}
}
