package tests

import (
	"testing"

	"github.com/jaavier/sider"
)

func TestCounters(t *testing.T) {
	t.Run("Increment key", func(t *testing.T) {
		var key = "counter"
		var expected int64 = 1
		var expected2 int64 = 2
		if sider.Incr(key) != expected {
			t.Errorf("Error incrementing key %s", key)
		}
		if sider.Incr(key) != expected2 {
			t.Errorf("Error incrementing key %s", key)
		}
	})

	t.Run("Decrement key", func(t *testing.T) {
		var key = "counter"
		var expected int64 = 1
		var expected2 int64 = 0
		if sider.Decr(key) != expected {
			t.Errorf("Error incrementing key %s", key)
		}
		if sider.Decr(key) != expected2 {
			t.Errorf("Error incrementing key %s", key)
		}
	})
}