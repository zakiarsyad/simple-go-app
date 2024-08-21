package repo

import (
	"be/model"

	"gorm.io/gorm"
)

type PostsRepo struct {
	db *gorm.DB
}

func NewPostsRepo(db *gorm.DB) PostsRepo {
	return PostsRepo{db}
}

func (p PostsRepo) Save(data model.Post) error {
	if err := p.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (p PostsRepo) Get() ([]model.Post, error) {
	var posts []model.Post
	if err := p.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p PostsRepo) GetById(id int) (*model.Post, error) {
	var post model.Post
	if err := p.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
