package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID       int64
	CourseID int
	Title    string
	Body     string
}

func GetPostByID(id string) (Post, error) {
	var post Post

	res := DB.First(&post, id)

	return post, res.Error
}

func GetPostsByCourse(courseID string) ([]Post, error) {
	var posts []Post

	res := DB.Find(&posts, "course_id=?", courseID)

	return posts, res.Error
}

func EditPost(id string, title string, body string) error {
	result := DB.Table("posts").Where("id = ?", id).Update("title", title).Update("body", body)
	if result.Error != nil {
		return result.Error
	}

	return nil
}