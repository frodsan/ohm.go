package rom_test

import (
	"testing"

	"github.com/frodsan/rom"
)

const URL = "redis://127.0.0.1:6379/13"

func TestFlushDB(t *testing.T) {
	db := rom.Connect(URL)

	db.Flush()
}
