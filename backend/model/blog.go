package model

type Post struct {
	Model
	Slug       string    `json:"slug" gorm:"not null;unique" validate:"required"`
	Title      string    `json:"title" gorm:"not null" validate:"required"`
	Summary    string    `json:"summary" gorm:"not null" validate:"required"`
	Content    string    `json:"content" gorm:"not null" validate:"required"`
	Banner     string    `json:"banner" gorm:"not null" validate:"required"`
	AuthorId   string    `json:"authorId" gorm:"not null" validate:"required"`
	CategoryId string    `json:"categoryId" gorm:"not null" validate:"required"`
	Author     User      `json:"author" gorm:"not null" validate:"-"`
	Category   *Category `json:"category" gorm:"not null" validate:"-"`
	Tags       []*Tag    `json:"tags" gorm:"many2many:post_tags" validate:"required"`
}

type Category struct {
	Model
	Slug   string `json:"slug" gorm:"not null;unique" validate:"required"`
	Name   string `json:"name" gorm:"not null;unique" validate:"required"`
	Detail string `json:"detail" gorm:"not null" validate:"required"`
	Icon   string `json:"icon" gorm:"not null" validate:"required"`
	Img    string `json:"img" gorm:"not null" validate:"required"`
}

type Tag struct {
	Model
	Name  string  `json:"name" gorm:"not null;unique" validate:"required"`
	Icon  string  `json:"icon" gorm:"not null" validate:"required"`
	Posts []*Post `json:"posts" gorm:"many2many:post_tags" validate:"-"`
}

type Comment struct {
	Model
	Content  string `json:"content" gorm:"not null" validate:"required"`
	AuthorId string `json:"authorId" gorm:"not null" validate:"required"`
	PostId   string `json:"postId" gorm:"not null" validate:"required"`
	Author   User   `json:"author" gorm:"not null" validate:"-"`
	Post     Post   `json:"post" gorm:"not null" validate:"-"`
}
