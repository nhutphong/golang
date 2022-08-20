package main

import (
	"fmt"
	"phong/tricks"
)

/*
		khi nhúng Author vào làm field của BlogPost , thì tức la BlogPost đã inheritance Author 
		co the access fields va method cua Author

		blog.firstName, blog.lastName, blog.bio
		blog.fullName()
*/


type Author struct {
	firstName string
	lastName  string
	bio       string
}

func (a Author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type BlogPost struct {
	title   string
	content string
	Author //embed struct Author == class BlogPost(Author)

	/*
		khi nhúng Author vào làm field của BlogPost , thì tức la BlogPost đã inheritance Author 
		co the access fields va method cua Author

		blog.firstName, blog.lastName, blog.bio
		blog.fullName()
	*/
}

func (p BlogPost) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type WebSite struct {
	BlogPosts []BlogPost
	// BlogPost
}

func (w WebSite) contents() {
	fmt.Println("Contents of WebSite\n")
	for _, blog := range w.BlogPosts {
		blog.details()
		fmt.Println("fullName = ", blog.firstName)
		fmt.Println("fullName = ", blog.lastName)
		fmt.Println("Author = ", blog.Author)
		fmt.Println()
	}

	// tricks.Format()

	// fmt.Println("firstName = ", w.firstName)
	// fmt.Println("lastName = ", w.lastName)
	// fmt.Println("fullName = ", w.firstName + w.lastName)
	// fmt.Println("title = ", w.title)
	// fmt.Println("content = ", w.content)
}	


func main() {
	Author1 := Author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	BlogPost1 := BlogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		Author1,
	}
	BlogPost2 := BlogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		Author1,
	}
	BlogPost3 := BlogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		Author1,
	}
	w := WebSite{
		BlogPosts: []BlogPost{BlogPost1, BlogPost2, BlogPost3},
	}
	w.contents()
}
