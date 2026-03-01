package main

import (
	"context"
	"finpro/config"
	"finpro/controller"
	"finpro/repository"
	"finpro/routes"
	"finpro/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database cok
	config.ConnectDatabase()

	// ndelok repo
	userRepo := repository.NewUserRepository(config.DB)
	assessRepo := repository.NewAssessmentRepository(config.DB)
	dashRepo := repository.NewDashboardRepository(config.DB)

	// Initialize Services
	aiService := service.NewAIService()
	authService := service.NewAuthService(userRepo)
	assessService := service.NewAssessmentService(assessRepo, aiService)
	dashService := service.NewDashboardService(dashRepo, assessRepo)
	adminService := service.NewAdminService(userRepo, assessRepo)

	// ndelok controller
	authCtrl := controller.NewAuthController(authService)
	assessCtrl := controller.NewAssessmentController(assessService)
	dashCtrl := controller.NewDashboardController(dashService)
	adminCtrl := controller.NewAdminController(adminService)

	// nggenakno ruter
	r := gin.Default()
	routes.SetupRoutes(r, authCtrl, assessCtrl, dashCtrl, adminCtrl)

	// nyambung server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server started on port %s", port)

	// ngenteni sinyal shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
