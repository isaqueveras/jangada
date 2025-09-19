package template

const WebControllerTest = `package controller

import "testing"

// Test{{ .Entity }}Controller ...
func Test{{ .Entity }}Controller(t *testing.T) {
	t.Run("GetByID", func(t *testing.T) {
		t.Fail()
	})
	t.Run("GetAll", func(t *testing.T) {
		t.Fail()
	})
	t.Run("Create", func(t *testing.T) {
		t.Fail()
	})
	t.Run("Update", func(t *testing.T) {
		t.Fail()
	})
	t.Run("Delete", func(t *testing.T) {
		t.Fail()
	})
}
`
