package router

import (
    "github.com/DaniloFaraum/studere-backend/controllers"
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
    controllers.InitializeController()

	v1 := router.Group("/api/v1")
    {
	answerRoutes := v1.Group("/tag")
        {
            answerRoutes.GET("/", controllers.ListTagsController)
            answerRoutes.POST("/", controllers.CreateTagController)
        }
	}
}
//     v1 := router.Group("/api/v1")
// 	{
// 		registerRoutes(v1, "answers", controllers.ListAnswerController, controllers.ShowAnswerController, controllers.CreateAnswerController, controllers.UpdateAnswerController, controllers.DeleteAnswerController)
// 		registerRoutes(v1, "answerTags", controllers.ListAnswerTagController, controllers.ShowAnswerTagController, controllers.CreateAnswerTagController, controllers.UpdateAnswerTagController, controllers.DeleteAnswerTagController)
// 		registerRoutes(v1, "comments", controllers.ListCommentController, controllers.ShowCommentController, controllers.CreateCommentController, controllers.UpdateCommentController, controllers.DeleteCommentController)
// 		registerRoutes(v1, "courses", controllers.ListCourseController, controllers.ShowCourseController, controllers.CreateCourseController, controllers.UpdateCourseController, controllers.DeleteCourseController)
// 		registerRoutes(v1, "courseTags", controllers.ListCourseTagController, controllers.ShowCourseTagController, controllers.CreateCourseTagController, controllers.UpdateCourseTagController, controllers.DeleteCourseTagController)
// 		registerRoutes(v1, "images", controllers.ListImageController, controllers.ShowImageController, controllers.CreateImageController, controllers.UpdateImageController, controllers.DeleteImageController)
// 		registerRoutes(v1, "questions", controllers.ListQuestionController, controllers.ShowQuestionController, controllers.CreateQuestionController, controllers.UpdateQuestionController, controllers.DeleteQuestionController)
// 		registerRoutes(v1, "questionnaires", controllers.ListQuestionnaireController, controllers.ShowQuestionnaireController, controllers.CreateQuestionnaireController, controllers.UpdateQuestionnaireController, controllers.DeleteQuestionnaireController)
// 		registerRoutes(v1, "questionnaireQuestions", controllers.ListQuestionnaireQuestionController, controllers.ShowQuestionnaireQuestionController, controllers.CreateQuestionnaireQuestionController, controllers.UpdateQuestionnaireQuestionController, controllers.DeleteQuestionnaireQuestionController)
// 		registerRoutes(v1, "questionTags", controllers.ListQuestionTagController, controllers.ShowQuestionTagController, controllers.CreateQuestionTagController, controllers.UpdateQuestionTagController, controllers.DeleteQuestionTagController)
// 		registerRoutes(v1, "ratings", controllers.ListRatingController, controllers.ShowRatingController, controllers.CreateRatingController, controllers.UpdateRatingController, controllers.DeleteRatingController)
// 		registerRoutes(v1, "roles", controllers.ListRoleController, controllers.ShowRoleController, controllers.CreateRoleController, controllers.UpdateRoleController, controllers.DeleteRoleController)
// 		registerRoutes(v1, "tags", controllers.ListTagController, controllers.ShowTagController, controllers.CreateTagController, controllers.UpdateTagController, controllers.DeleteTagController)
// 		registerRoutes(v1, "users", controllers.ListUserController, controllers.ShowUserController, controllers.CreateUserController, controllers.UpdateUserController, controllers.DeleteUserController)
// 	}
// }

// func registerRoutes(group *gin.RouterGroup, resource string, list, show, create, update, delete gin.HandlerFunc) {
//     routes := group.Group("/" + resource)
//     {
//         routes.GET("/", list)
//         routes.GET("/:id", show)
//         routes.POST("/", create)
//         routes.PUT("/:id", update)
//         routes.DELETE("/:id", delete)
//     }
// }