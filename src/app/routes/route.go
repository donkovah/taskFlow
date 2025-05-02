package routes

import (
	"be/src/app/controllers"
	authController "be/src/app/controllers/Auth"
	commentController "be/src/app/controllers/Comment"
	noteController "be/src/app/controllers/Note"
	projectController "be/src/app/controllers/Project"
	taskController "be/src/app/controllers/Task"
	userController "be/src/app/controllers/User"
	"be/src/domain/service"
	"be/src/infrastructure/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.Engine,
	projectService *service.ProjectService,
	taskService *service.TaskService,
	authService *service.AuthService,
	commentService *service.CommentService,
	noteService *service.NoteService,
	timelineService *service.TimelineService,
	userService *service.UserService,
) {

	projectController := projectController.NewProjectController(projectService)
	taskController := taskController.NewTaskController(taskService)
	authController := authController.NewAuthController(authService)
	commentController := commentController.NewCommentController(commentService)
	noteController := noteController.NewNoteController(noteService)
	timelineController := controllers.NewTimelineController(timelineService)
	userController := userController.NewUserController(userService)

	// Public auth routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
	}

	// Protected routes
	v1 := r.Group("/v1")
	v1.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))
	{
		// Auth routes that require authentication
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/logout", authController.Logout)
		}

		userRoute := v1.Group("/users")
		{
			userRoute.GET("/:id", userController.GetUser)
			userRoute.GET("/", userController.GetUsers)
			userRoute.POST("/", userController.CreateUser)
			userRoute.PUT("/:id", userController.UpdateUser)
			userRoute.DELETE("/:id", userController.DeleteUser)
		}

		projectRoute := v1.Group("/projects")
		{
			projectRoute.GET("", projectController.GetProjects)
			projectRoute.GET("/:id", projectController.GetProject)
			projectRoute.POST("", projectController.CreateProject)
			projectRoute.PUT("/:id", projectController.UpdateProject)
			projectRoute.DELETE("/:id", projectController.DeleteProject)
		}

		taskRoute := v1.Group("/tasks")
		{
			taskRoute.GET("/:id", taskController.GetTask)
			taskRoute.GET("/", taskController.GetTasks)
			taskRoute.POST("/", taskController.CreateTask)
			taskRoute.PUT("/:id", taskController.UpdateTask)
			taskRoute.DELETE("/:id", taskController.DeleteTask)
			taskRoute.PATCH("/:id/start", taskController.StartTask)
			taskRoute.PATCH("/:id/block", taskController.BlockTask)
			taskRoute.PATCH("/:id/complete", taskController.CompleteTask)
		}

		noteRoute := v1.Group("/notes")
		{
			noteRoute.GET("/:id", noteController.GetNote)
			noteRoute.GET("/", noteController.GetNotes)
			noteRoute.POST("/", noteController.CreateNote)
			noteRoute.PUT("/:id", noteController.UpdateNote)
			noteRoute.DELETE("/:id", noteController.DeleteNote)
		}

		timelineRoute := v1.Group("/timelines")
		{
			timelineRoute.GET("/:id", timelineController.GetTimeline)
			timelineRoute.GET("/", timelineController.GetTimelines)
			timelineRoute.POST("/", timelineController.CreateTimeline)
			timelineRoute.PUT("/:id", timelineController.UpdateTimeline)
			timelineRoute.DELETE("/:id", timelineController.DeleteTimeline)
		}

		commentRoute := v1.Group("/comments")
		{
			commentRoute.GET("/:id", commentController.GetComment)
			commentRoute.GET("/", commentController.GetComments)
			commentRoute.POST("/", commentController.CreateComment)
			commentRoute.PUT("/:id", commentController.UpdateComment)
			commentRoute.DELETE("/:id", commentController.DeleteComment)
		}
	}
}
