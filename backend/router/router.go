package router

import (
	"c0ding-backend/api"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	user := r.Group("/api/users")
	{
		user.POST("/register", api.UserRegister)
		user.POST("/login", api.UserLogin)
	}

	post := r.Group("/api/posts")
	{
		post.POST("/", api.UploadPost)
		post.POST("/:postId", api.UploadLikes)
		post.DELETE("/:postId")
	}

}
