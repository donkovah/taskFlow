package models

import "fmt"

// Status represents the task status as a string.
type Status string

// Status constants representing task status values.
const (
	StatusTodo      Status = "Todo"
	StatusProgress  Status = "In Progress"
	StatusBlocked   Status = "Blocked"
	StatusCompleted Status = "Completed"
)

// ValidStatuses holds all possible statuses for quick lookup and validation.
var ValidStatuses = map[Status]bool{
	StatusTodo:      true,
	StatusProgress:  true,
	StatusBlocked:   true,
	StatusCompleted: true,
}

// String returns the string representation of the status.
func (s Status) String() string {
	return string(s)
}

// IsValid checks if the status is a valid predefined status.
func (s Status) IsValid() bool {
	_, exists := ValidStatuses[s]
	return exists
}

// ParseStatus converts a string into a Status type, validating it.
func ParseStatus(status string) (Status, error) {
	s := Status(status)
	if !s.IsValid() {
		return "", fmt.Errorf("invalid status: %s", status)
	}
	return s, nil
}
