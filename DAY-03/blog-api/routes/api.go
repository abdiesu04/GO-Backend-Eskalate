package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Blog represents a blog post with a title, description, body, author, and publication status.
// swagger:model
type Blog struct {
	// The ID of the blog
	//
	// example: 1
	ID int `json:"id"`

	// The title of the blog
	// required: true
	// example: My first blog
	Title string `json:"title"`

	// The description of the blog
	//
	// example: This is my first blog
	Description string `json:"description"`

	// The body of the blog
	// required: true
	// example: This is the body of my first blog
	Body string `json:"body"`

	// The author of the blog
	// required: true
	// example: John Doe
	Author string `json:"author"`

	// The publication status of the blog
	// required: true
	// example: true
	IsPublished bool `json:"isPublished"`
}

var blogs = []Blog{
	{
		ID:          1,
		Title:       "My first blog",
		Description: "This is my first blog",
		Body:        "This is the body of my first blog",
		Author:      "John Doe",
		IsPublished: true,
	},
	{
		ID:          2,
		Title:       "My second blog",
		Description: "This is my second blog",
		Body:        "This is the body of my second blog",
		Author:      "Jane Doe",
		IsPublished: false,
	},
	{
		ID:          3,
		Title:       "My third blog",
		Description: "This is my third blog",
		Body:        "This is the body of my third blog",
		Author:      "John Doe",
		IsPublished: true,
	},
	{
		ID:          4,
		Title:       "My fourth blog",
		Description: "This is my fourth blog",
		Body:        "This is the body of my fourth blog",
		Author:      "Jane Doe",
		IsPublished: true,
	},
}


func (b *Blog) GetBlogs(c *gin.Context) {
 c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":              
blogs})
}


func (b *Blog) GetBlog(c *gin.Context){
	id := c.Param("id")
	intID , err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusNotFound , gin.H{
			"status" : http.StatusNotFound,
			"message" : err.Error(),
			} )
			return 
	}

	// find blog whose if matched the Param id

	for _ , blog := range blogs{
		if blog.ID == intID {
			c.JSON(http.StatusOK , gin.H{
				"status":http.StatusOK,
				"data" : blog,
				 
			})
			return
	}

	c.JSON(http.StatusOK , gin.H{
		"status":http.StatusNotFound,
		"message" : "Blog not found",
	})

}
}

func (b *Blog) CreateBlog(c *gin.Context){
	var incomingBlog Blog
	
	incomingBlog.ID = len(blogs) + 1

	err := c.BindJSON(&incomingBlog)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status":http.StatusBadRequest , 
			"message" : err.Error(),
		})
		return
	}

	blogs = append(blogs , incomingBlog)

	c.JSON(http.StatusCreated , gin.H{
		"status":http.StatusCreated,
		"data":incomingBlog,
	})
}


func (b *Blog) UpdateBlog(c *gin.Context){
	id := c.Param("id")

	intID  , err  := strconv.Atoi(id)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": err.Error(),
		})
		return 
	}

	// find the blog with matching id

	for index , blog := range blogs{
		if blog.ID == intID{
			//parse the request body to get the updated blog data

			var updatedBlog Blog

			err := c.BindJSON(&updatedBlog)

			if err != nil{
				c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
				return
			}

			// update the blog with the new data

			updatedBlog.ID = intID
			blogs[index] = updatedBlog

			// Respind with the updated Blog

			c.JSON(http.StatusOK , gin.H{
				"status" : http.StatusOK, 
				"data" : updatedBlog,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Blog not found"})
}


func (b *Blog) DeleteBlog(c *gin.Context){
	id := c.Param("id")

	intID , err := strconv.Atoi(id)

	if err != nil{
		c.JSON(http.StatusNotFound , gin.H{
			"status":http.StatusNotFound,
			"message":err.Error(),
		})
		return 
	}

	for index , blog := range blogs{
		if blog.ID == intID{
			blogs = append(blogs[:index],blogs[index+1:]...)
			c.JSON(
				http.StatusOK ,
				gin.H{
					"status": http.StatusOK ,
					"message": "Blog Deleted Successfully",
				},
			)
			return 
		}
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Blog could not be deleted. Blog not found"})
	}
}