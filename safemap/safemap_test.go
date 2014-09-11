package safemap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialLenIsZero(t *testing.T) {
	t.Parallel()
	assert.Equal(t, New().Len(), 0)
}

func TestKeyNotFound(t *testing.T) {
	t.Parallel()
	result, found := New().Find("unknown_key")
	assert.Nil(t, result)
	assert.False(t, found)
}

func TestNewKeyInsertion(t *testing.T) {
	t.Parallel()
	sm := New()
	sm.Insert("key", "value")
	assert.Equal(t, sm.Len(), 1)
	result, found := sm.Find("key")
	assert.Equal(t, result, "value")
	assert.True(t, found)
}

func TestDeleteAKey(t *testing.T) {
	t.Parallel()
	sm := New()
	sm.Insert("key", "value")
	sm.Delete("key")
	result, found := sm.Find("key")
	assert.Nil(t, result)
	assert.False(t, found)
}

func TestUpdateAnExistingKey(t *testing.T) {
	t.Parallel()
	sm := New()
	sm.Insert("key", "value")
	sm.Update("key", func(value interface{}, found bool) interface{} {
		return value.(string) + "_updated"
	})
	result, found := sm.Find("key")
	assert.Equal(t, result, "value_updated")
	assert.True(t, found)
}

func TestUpdateANonExistingKey(t *testing.T) {
	t.Parallel()
	sm := New()
	sm.Update("key", func(value interface{}, found bool) interface{} {
		return fmt.Sprintf("%s_updated", value)
	})
	result, found := sm.Find("key")
	assert.Equal(t, result, "%!s(<nil>)_updated")
	assert.True(t, found)
}

func TestClose(t *testing.T) {
	t.Parallel()
	sm := New()
	sm.Close()
	assert.Panics(t, func() {
		sm.Insert("key", "value")
	})
}
