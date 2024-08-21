package handler

import (
	"be/model"
	"be/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostsHandler struct {
	repo repo.PostsRepo
}

func NewPostsHandler(repo repo.PostsRepo) PostsHandler {
	return PostsHandler{
		repo: repo,
	}
}

func (p PostsHandler) CreatePosts(c *gin.Context) {
	var newPost model.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := p.repo.Save(newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Postingan berhasil ditambahkan"})
}

func (p PostsHandler) GetPosts(c *gin.Context) {
	posts, err := p.repo.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (p PostsHandler) GetPostsById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	post, err := p.repo.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, post)
}
