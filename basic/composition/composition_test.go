/**
 * @Author: Noaghzil
 * @Date:   2023-02-13 16:14:41
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-13 16:24:41
 */
package composition

import (
	"testing"
)

func TestComposition(t *testing.T) {
	auther_one := Author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	auther_two := Author{
		"Sam",
		"Adolf",
		"Java Enthusiast",
	}

	blog_one := BlogPost{
		"Go Interface",
		"Go interface is named collection of method signatures",
		auther_one,
	}

	blog_two := BlogPost{
		"Java Interface",
		"Java interface is named collection of method signatures",
		auther_two,
	}

	blog_three := BlogPost{
		"Go Goroutine",
		"Go Goroutine is lightweight thread",
		auther_one,
	}

	website := Website{
		posts: []BlogPost{blog_one, blog_two, blog_three},
	}
	website.Contents()
}
