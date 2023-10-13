package dto

import "github.com/carlosCACB333/cb-back/model"

type UpdateTag struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type UpdateCategory struct {
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Icon   string `json:"icon"`
	Img    string `json:"img"`
}

type UpdatePost struct {
	Title      string       `json:"title"`
	Slug       string       `json:"slug"`
	Summary    string       `json:"summary"`
	Content    string       `json:"content"`
	Banner     string       `json:"banner"`
	CategoryId string       `json:"categoryId"`
	Tags       []*model.Tag `json:"tags" gorm:"many2many:post_tags"`
}

type UpdateComment struct {
	Content string `json:"content"`
}
