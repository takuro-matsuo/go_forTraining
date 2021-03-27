package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	// post := Post{Content: "aaa", Author: "sau"}
	// fmt.Println(post)

	// Db.Create(&post)
	// fmt.Println(post)

	// comment1 := Comment{Content: "good", Author: "joe"}
	// Db.Model(&post).Association("Comments").Append(comment1)

	// comment2 := Comment{Content: "bad", Author: "ma"}
	// Db.Model(&post).Association("Comments").Append(comment2)

	var readPost Post
	Db.Where("author = $1", "sau").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)

	fmt.Println(comments)
	fmt.Println(comments[0])
	fmt.Println(comments[1])
}
