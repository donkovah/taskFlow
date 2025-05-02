package structs

import (
	"be/src/domain/models"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Task struct {
	ProjectID   uuid.UUID `json:"projectId" validate:"required"`
	Title       string    `json:"title" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"max=255"`
	Deadline    time.Time `json:"deadline,omitempty"`
}

type TaskValidator struct {
	validator *validator.Validate
}

// NewTaskValidator initializes a new TaskValidator.
func NewTaskValidator() *TaskValidator {
	v := validator.New()
	v.RegisterValidation("uuid", validateUUID)
	return &TaskValidator{validator: v}
}

// Validate validates the Task struct.
func (tv *TaskValidator) Validate(task *Task) error {
	err := tv.validator.Struct(task)
	if err != nil {
		// Return a user-friendly error message
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("validation failed for field '%s': %s", err.Field(), err.Tag())
		}
	}
	return nil
}

// validateUUID checks if a field is a valid UUID.
func validateUUID(fl validator.FieldLevel) bool {
	_, err := uuid.Parse(fl.Field().String())
	return err == nil
}

// MapModelTaskToStructsTask maps models.Task to structs.Task for validation.
func MapModelTaskToStructsTask(modelTask *models.Task) *Task {
	var deadline time.Time

	if modelTask.Deadline != nil {
		deadline = *modelTask.Deadline
	}

	return &Task{
		ProjectID:   modelTask.ProjectID,
		Title:       modelTask.Title,
		Description: modelTask.Description,
		Deadline:    deadline,
	}
}
