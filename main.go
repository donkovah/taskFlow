package main

import (
	"be/src/app/routes"
	"be/src/domain/service"
	"be/src/infrastructure/config"
	"be/src/infrastructure/persistence"
	infraService "be/src/infrastructure/service"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Load application configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set Gin mode based on config
	if !config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	projectService := service.NewProjectService(persistence.NewProjectRepository(config.DB))
	taskService := service.NewTaskService(persistence.NewTaskRepository(config.DB))
	authService := service.NewAuthService(infraService.NewAuthService(persistence.NewUserRepository(config.DB), config.App.JWTSecret))
	commentService := service.NewCommentService(persistence.NewCommentRepository(config.DB))
	noteService := service.NewNoteService(persistence.NewNoteRepository(config.DB))
	timelineService := service.NewTimelineService(persistence.NewTimelineRepository(config.DB))
	userService := service.NewUserService(persistence.NewUserRepository(config.DB))

	routes.InitRoutes(
		r,
		projectService,
		taskService,
		authService,
		commentService,
		noteService,
		timelineService,
		userService,
	)

	// Listen and serve on the specified port
	if err := r.Run(fmt.Sprintf(":%d", config.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
