package routes

import (
	"be/src/app/controllers"
	"be/src/domain/service"

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

	projectController := controllers.NewProjectController(projectService)
	taskController := controllers.NewTaskController(taskService)
	authController := controllers.NewAuthController(authService)
	commentController := controllers.NewCommentController(commentService)
	noteController := controllers.NewNoteController(noteService)
	timelineController := controllers.NewTimelineController(timelineService)
	userController := controllers.NewUserController(userService)

	authGroup := r.Group("/auth")
	authRoute := authGroup.Group("/comments")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
		authRoute.POST("/logout", authController.Logout)
	}

	v1 := r.Group("/v1")

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
		taskRoute.PATCH("/:id/status", taskController.GetTask)
		taskRoute.DELETE("/:id", taskController.DeleteTask)
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
