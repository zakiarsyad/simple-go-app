package rest

import (
	"be/handler"
	"be/repo"

	"github.com/gin-gonic/gin"
)

func PostsRoutes(r *gin.Engine, repo repo.PostsRepo) {
	h := handler.NewPostsHandler(repo)

	post := r.Group("/posts")
	post.POST("", h.CreatePosts)
	post.GET("", h.GetPosts)
	post.GET("/:id", h.GetPostsById)
}
