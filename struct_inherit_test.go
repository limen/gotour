package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mapStr map[string]interface{}

type entity struct {
	Data mapStr
}

func (en *entity) data() mapStr {
	return en.Data
}

func (en *entity) setData(data mapStr) {
	en.Data = data
}

type userEntity struct {
	entity        // inheritance
	Data   mapStr // shadows `entity.Data`
}

type tagEntity struct {
	entity // No shadow. tagEntity.Data is pointing to entity.Data
}

func TestStructInheritance(t *testing.T) {
	userEn := userEntity{}
	alice := make(mapStr)
	alice["name"] = "alice"
	alice["age"] = 22
	userEn.Data = make(mapStr)
	userEn.Data["name"] = alice["name"]
	userEn.Data["age"] = alice["age"]
	// userEn.data() would return the shadowed `Data`
	// So NotEqual
	assert.NotEqual(t, alice, userEn.data())
	// The shadowed `Data` should be nil
	assert.Nil(t, userEn.data())

	// explicitly set shadowed `Data`
	userEn.entity.Data = alice
	assert.Equal(t, alice, userEn.data())

	tagEn := tagEntity{}
	tagEn.Data = alice
	// No shadow
	assert.Equal(t, alice, userEn.data())

	bob := make(mapStr)
	bob["name"] = "bob"
	bob["age"] = 25
	// No shadow
	tagEn.setData(bob)
	assert.Equal(t, bob, tagEn.data())

	// shadowed `Data` set
	userEn.setData(bob)
	assert.NotEqual(t, bob, userEn.Data)
	assert.Equal(t, bob, userEn.entity.Data)
}
