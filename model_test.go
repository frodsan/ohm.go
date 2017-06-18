package rom_test

import (
	"testing"

	"github.com/frodsan/rom"
)

type User struct {
	rom.Model
}

func TestIsNew(t *testing.T) {
	user := User{}

	if !user.IsNew() {
		t.Error("Expected record to be new")
	}

	user.Id = "1"

	if user.IsNew() {
		t.Error("Expected record to not be new")
	}
}

func TestIsPersisted(t *testing.T) {
	user := User{rom.Model{Id: "1"}}

	if !user.IsPersisted() {
		t.Error("Expected record to be persisted")
	}

	user.Id = ""

	if user.IsPersisted() {
		t.Error("Expected record to not be persisted")
	}
}
