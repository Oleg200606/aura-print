package handlers

import (
	"slices"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestUUIDSplit(t *testing.T) {
	id, err := uuid.Parse("12345678-1234-5678-9012-345678901234")
	if err != nil {
		t.Errorf("UUIDParse(%s) = %v; want %v", id.String(), err, nil)
	}
	expected := []string{"12345678", "1234", "5678", "9012", "345678901234"}
	actual := strings.Split(id.String(), "-")
	if slices.Compare(expected, actual) != 0 {
		t.Errorf("UUIDSplit(%s) = %v; want %v", id.String(), actual, expected)
	}
}
