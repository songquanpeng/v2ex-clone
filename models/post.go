// Credit: https://segmentfault.com/a/1190000013297747
package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Post struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` // Foreign key
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistPostByID(id int) bool {
	var post Post
	db.Select("id").Where("id = ?", id).First(&post)

	if post.ID > 0 {
		return true
	}

	return false
}

func GetPostNum(maps interface{}) (count int) {
	db.Model(&Post{}).Where(maps).Count(&count)

	return
}

func GetPosts(pageNum int, pageSize int, maps interface{}) (posts []Post) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&posts)

	return
}

func GetPost(id int) (post Post) {
	db.Where("id = ?", id).First(&post)
	db.Model(&post).Related(&post.Tag)

	return
}

func EditPost(id int, data interface{}) bool {
	db.Model(&Post{}).Where("id = ?", id).Updates(data)

	return true
}

func AddPost(data map[string]interface{}) bool {
	db.Create(&Post{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeletePost(id int) bool {
	db.Where("id = ?", id).Delete(Post{})

	return true
}

func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	e := scope.SetColumn("CreatedOn", time.Now().Unix())
	return e
}

func (post *Post) BeforeUpdate(scope *gorm.Scope) error {
	e := scope.SetColumn("ModifiedOn", time.Now().Unix())
	return e
}
