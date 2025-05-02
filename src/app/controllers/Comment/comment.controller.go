package commentController

import (
	"be/src/domain/models"
	"be/src/domain/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service *service.CommentService
}

func NewCommentController(service *service.CommentService) *CommentController {
	return &CommentController{service: service}
}
func (tc *CommentController) GetComment(c *gin.Context) {
	id := c.Param("id")
	comment, err := tc.service.GetComment(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comment"})
	}
	c.JSON(http.StatusOK, comment)
}

func (tc CommentController) GetComments(c *gin.Context) {
	comments, err := tc.service.GetComments(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch comment"})
	}
	c.JSON(http.StatusOK, comments)
}

func (tc CommentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdComment, err := tc.service.CreateComment(context.Background(), &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
	}

	c.JSON(http.StatusOK, createdComment)
}

func (tc CommentController) UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var commentBody *models.Comment

	if err := c.ShouldBindJSON(&commentBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := tc.service.GetComment(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comment"})
	}
	comment.Name = commentBody.Name

	updatedComment, err := tc.service.UpdateComment(context.Background(), comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
	}

	c.JSON(http.StatusOK, updatedComment)
}

func (ts CommentController) DeleteComment(c *gin.Context) {
	id := c.Param(("id"))
	err := ts.service.DeleteComment(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
	}
	c.JSON(http.StatusNoContent, nil)
}
