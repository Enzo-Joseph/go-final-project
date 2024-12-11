package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	LecturerId int
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

func CreatePost(courseID int, userID int, title string, body string) (int, error) {
    post := Post{LecturerId: userID, CourseID: courseID, Title: title, Body: body}
    result := DB.Create(&post)

    if result.Error != nil {
        return 0, result.Error
    }

    return post.ID, nil
}

func EditPost(id string, title string, body string) error {
	result := DB.Table("posts").Where("id = ?", id).Update("title", title).Update("body", body)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeletePost(id string) error {
	result := DB.Delete(&Post{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}