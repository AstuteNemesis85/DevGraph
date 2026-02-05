package analysis

import "github.com/google/uuid"

var JobQueue = make(chan uuid.UUID, 100)
