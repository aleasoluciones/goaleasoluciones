// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package safemap_test

import (
	"fmt"
	"testing"

	. "github.com/aleasoluciones/goaleasoluciones/safemap"

	"github.com/stretchr/testify/assert"
)

func TestInitialLenIsZero(t *testing.T) {
	t.Parallel()
	assert.Equal(t, NewSafeMap().Len(), 0)
}

func TestKeyNotFound(t *testing.T) {
	t.Parallel()
	result, found := NewSafeMap().Find("unknown_key")
	assert.Nil(t, result)
	assert.False(t, found)
}

func TestNewSafeMapKeyInsertion(t *testing.T) {
	t.Parallel()
	sm := NewSafeMap()
	sm.Insert("key", "value")
	assert.Equal(t, sm.Len(), 1)
	result, found := sm.Find("key")
	assert.Equal(t, result, "value")
	assert.True(t, found)
}

func TestDeleteAKey(t *testing.T) {
	t.Parallel()
	sm := NewSafeMap()
	sm.Insert("key", "value")
	sm.Delete("key")
	result, found := sm.Find("key")
	assert.Nil(t, result)
	assert.False(t, found)
}

func TestDeleteANonExistingKey(t *testing.T) {
	t.Parallel()
	sm := NewSafeMap()
	sm.Delete("key")
	result, found := sm.Find("key")
	assert.Nil(t, result)
	assert.False(t, found)
}

func TestUpdateAnExistingKey(t *testing.T) {
	t.Parallel()
	sm := NewSafeMap()
	sm.Insert("key", "value")
	sm.Update("key", func(value Value, found bool) Value {
		return value.(string) + "_updated"
	})
	result, found := sm.Find("key")
	assert.Equal(t, result, "value_updated")
	assert.True(t, found)
}

func TestUpdateANonExistingKey(t *testing.T) {
	t.Parallel()
	sm := NewSafeMap()
	sm.Update("key", func(value Value, found bool) Value {
		return fmt.Sprintf("%s_updated", value)
	})
	result, found := sm.Find("key")
	assert.Equal(t, result, "%!s(<nil>)_updated")
	assert.True(t, found)
}

func TestClose(t *testing.T) {
	t.Parallel()
	sm := NewSafeMap()
	sm.Close()
	assert.Panics(t, func() {
		sm.Insert("key", "value")
	})
}

func TestKeys(t *testing.T) {
	t.Parallel()
	sm := NewSafeMap()
	sm.Insert("key", "value")
	sm.Insert("key1", "value1")

	result := sm.Keys()

	assert.Len(t, result, 2)
	assert.Contains(t, result, "key")
	assert.Contains(t, result, "key1")
}
