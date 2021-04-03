// リスト7.11
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Posts struct {
	// Posts []Post `json:"posts"`
	Posts []Post
}
type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file: ", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var posts Posts
		err := decoder.Decode(&posts)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON data: ", err)
			return
		}
		psts := posts.Posts
		for _, pst := range psts {
			fmt.Println(pst)
		}
	}
}
