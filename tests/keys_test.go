package tests

import (
	"testing"

	"github.com/jaavier/sider"
)

func TestKeys(t *testing.T) {
	t.Run("Add Keys", func(t *testing.T) {
		var key string = "my-key"
		var value string = "my-value"
		if result, err := sider.Set(key, value); err != nil && !result {
			t.Errorf("Error adding key '%v'", key)
		}
	})

	t.Run("Read Key", func(t *testing.T) {
		var key string = "my-key"
		var expected string = "my-value"
		if result, err := sider.Get(key); err != nil && result != expected {
			t.Errorf("Error reading key '%v'", key)
		}
	})

	t.Run("Delete Key", func(t *testing.T) {
		var newKey string = "delete-key"
		sider.Set(newKey, "value")
		if err := sider.DeleteKey(newKey); err != nil {
			t.Errorf("Error deleting key '%s'", newKey)
		}
	})
}
