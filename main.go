package main

import (
	"be/config"
	"be/model"
	"be/repo"
	"be/rest"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo repo.PostsRepo) *gin.Engine {
	r := gin.Default()

	rest.PostsRoutes(r, repo)

	return r
}

func main() {
	db := config.NewDB()
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&model.Post{})
	if err != nil {
		panic(err)
	}

	postsRepo := repo.NewPostsRepo(conn)

	r := SetupRouter(postsRepo)

	r.Run(":8080")
}
