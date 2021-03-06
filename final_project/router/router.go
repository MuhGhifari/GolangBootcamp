package router

import (
	"github.com/MuhGhifari/GolangBootcamp/final-project/controllers"
	"github.com/MuhGhifari/GolangBootcamp/final-project/middlewares"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", middlewares.UserAuthorization(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.UserAuthorization(), controllers.DeleteUser)
	}
	
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("", controllers.GetPhotos)
		photoRouter.POST("", controllers.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}
	
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("", controllers.CreateComment)
		commentRouter.GET("", controllers.GetComments)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}
	
	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("", controllers.CreateSocialMedia)
		socialMediaRouter.GET("", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
