package rom

import (
	"testing"
)

const URL = "redis://127.0.0.1:6379/13"

func TestFlushDB(t *testing.T) {
	db := Connect(URL)

	db.Flush()
}
