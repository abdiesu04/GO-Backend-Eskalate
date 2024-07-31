// Package routes Blog API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Host: localhost:8080
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package routes




// swagger:route GET /blogs blogs getBlogs
//
// GetBlogs returns all blogs.
//
// Responses:
//
//	200: successResponse
func GetBlogs() {}

// swagger:route GET /blogs/{id} blogs getBlog
//
// GetBlog returns a blog by its ID.
//
// Responses:
//
//	200: successResponseR
//    400: errorResponse
func GetBlog() {}

// swagger:route POST /blogs blogs createBlog
//
// CreateBlog creates a new blog and returns it.
//
// Responses:
//
//	201: successResponse
//	400: errorResponse
func CreateBlog() {}

// swagger:route PUT /blogs/{id} blogs updateBlog
//
// UpdateBlog updates a blog by its ID.
//
// Responses:
//
//	200: successResponse
//	400: errorResponse
//	404: errorResponse
func UpdateBlog() {}

// swagger:route DELETE /blogs/{id} blogs deleteBlog
//
// DeleteBlog deletes a blog by its ID.
//
// Responses:
//
//	200: successResponse
//    404: errorResponse
func DeleteBlog() {}