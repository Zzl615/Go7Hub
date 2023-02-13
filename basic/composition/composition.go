/**
 * @Author: Noaghzil
 * @Date:   2023-02-13 16:14:31
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-13 16:24:36
 */
package composition

import (
	"fmt"
)

// Author represents a book author

type Author struct {
	firstName string
	lastName  string
	bio       string
}

func (a *Author) FullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// BlogPost represents a blog post

type BlogPost struct {
	title   string
	content string
	auther  Author
}

func (b *BlogPost) Details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.auther.FullName())
	fmt.Println("Bio: ", b.auther.bio)
}

// Website represents a website

type Website struct {
	posts []BlogPost
}

func (w *Website) Contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.posts {
		v.Details()
		fmt.Println()
	}
}
