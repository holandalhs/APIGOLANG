package internally

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"-"`
	Username  string    `json:"username"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
