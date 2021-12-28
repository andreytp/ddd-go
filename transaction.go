package tavern

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	amount   int
	from     uuid.UUID
	to       uuid.UUID
	createAt time.Time
}
