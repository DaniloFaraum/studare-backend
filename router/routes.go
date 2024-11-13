package router

import (
	"github.com/DaniloFaraum/studere-backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	controllers.InitializeController()

	v1 := router.Group("/api/v1")
	{
		registerRoutes(v1, "answers", controllers.ListAnswersController, controllers.ShowAnswerController, controllers.CreateAnswerController, controllers.UpdateAnswerController, controllers.DeleteAnswerController)
		registerRoutes(v1, "comments", controllers.ListCommentsController, controllers.ShowCommentController, controllers.CreateCommentController, controllers.UpdateCommentController, controllers.DeleteCommentController)
		registerRoutes(v1, "courses", controllers.ListCoursesController, controllers.ShowCourseController, controllers.CreateCourseController, controllers.UpdateCourseController, controllers.DeleteCourseController)
		//registerRoutes(v1, "courseTags", controllers.ListCourseTagController, controllers.ShowCourseTagController, controllers.CreateCourseTagController, controllers.UpdateCourseTagController, controllers.DeleteCourseTagController)
		registerRoutes(v1, "questions", controllers.ListQuestionsController, controllers.ShowQuestionController, controllers.CreateQuestionController, controllers.UpdateQuestionController, controllers.DeleteQuestionController)
		registerRoutes(v1, "questionnaires", controllers.ListQuestionnairesController, controllers.ShowQuestionnaireController, controllers.CreateQuestionnaireController, controllers.UpdateQuestionnaireController, controllers.DeleteQuestionnaireController)
		registerRoutes(v1, "ratings", controllers.ListRatingsController, controllers.ShowRatingController, controllers.CreateRatingController, controllers.UpdateRatingController, controllers.DeleteRatingController)
		registerRoutes(v1, "roles", controllers.ListRolesController, controllers.ShowRoleController, controllers.CreateRoleController, controllers.UpdateRoleController, controllers.DeleteRoleController)
		registerRoutes(v1, "tags", controllers.ListTagsController, controllers.ShowTagController, controllers.CreateTagController, controllers.UpdateTagController, controllers.DeleteTagController)

		imageRoutes := v1.Group("/images")
		{
			imageRoutes.GET("/all", controllers.ListImagesController)
			imageRoutes.GET("/:id", controllers.ShowImageController)
			imageRoutes.POST("/", controllers.CreateImageController)
			imageRoutes.DELETE("/:id", controllers.DeleteImageController)
		}
		userRoutes := v1.Group("/users")
		{
			userRoutes.GET("/all", controllers.ListUsersController)
			userRoutes.GET("/:id", controllers.ShowUserController)
			userRoutes.POST("/", controllers.CreateUserController)
			userRoutes.PUT("/:id", controllers.UpdateUserController)
			userRoutes.POST("/login", controllers.Login)
		}
	}
}

func registerRoutes(group *gin.RouterGroup, resource string, list, show, create, update, delete gin.HandlerFunc) {
	routes := group.Group("/" + resource)
	{
		routes.GET("/all", list)
		routes.GET("/:id", show)
		routes.POST("/", create)
		routes.PUT("/:id", update)
		routes.DELETE("/:id", delete)
	}
}
