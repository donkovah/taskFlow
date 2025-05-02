package noteController

import (
	"be/src/domain/models"
	"be/src/domain/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	service *service.NoteService
}

func NewNoteController(service *service.NoteService) *NoteController {
	return &NoteController{service: service}
}
func (tc *NoteController) GetNote(c *gin.Context) {
	id := c.Param("id")
	note, err := tc.service.GetNote(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get note"})
	}
	c.JSON(http.StatusOK, note)
}

func (tc NoteController) GetNotes(c *gin.Context) {
	notes, err := tc.service.GetNotes(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch note"})
	}
	c.JSON(http.StatusOK, notes)
}

func (tc NoteController) CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdNote, err := tc.service.CreateNote(context.Background(), &note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
	}

	c.JSON(http.StatusOK, createdNote)
}

func (tc NoteController) UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var noteBody *models.Note

	if err := c.ShouldBindJSON(&noteBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := tc.service.GetNote(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch note"})
	}
	note.Name = noteBody.Name
	updatedNote, err := tc.service.UpdateNote(context.Background(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
	}

	c.JSON(http.StatusOK, updatedNote)
}

func (ts NoteController) DeleteNote(c *gin.Context) {
	id := c.Param(("id"))
	err := ts.service.DeleteNote(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
	}
	c.JSON(http.StatusNoContent, nil)
}
