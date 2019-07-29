package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

type gender struct {
	G    bool
	Name string
}

type user struct {
	Name   string    `json:"name"`
	BornAt time.Time `json:"bornAt"`
	Gender gender    `json:"gender"`
}

func TestReflect(t *testing.T) {
	u := user{
		Name:   "json",
		BornAt: time.Now(),
		Gender: gender{G: true, Name: "Male"},
	}
	uType := reflect.TypeOf(u)
	uNumField := uType.NumField()
	assert.Equal(t, uNumField, 3)
	// 获取u.Name的元数据
	uFieldName := uType.Field(0)
	assert.Equal(t, uFieldName.Type.Name(), "string")
	assert.Equal(t, uType.Field(2).Type, reflect.TypeOf(gender{}))
	assert.Equal(t, uFieldName.Name, "Name")
	assert.Equal(t, uFieldName.Tag.Get("json"), "name")
	assert.Equal(t, uFieldName.Tag.Get("foo"), "")
	// 获取u的第一个field, 这里是u.Name
	assert.Equal(t, reflect.ValueOf(u).Field(0).String(), "json")
	// 取第2个属性, 并转为interface{}类型
	fieldBornAt := reflect.ValueOf(u).Field(1).Interface()
	assert.Equal(t, fieldBornAt.(time.Time).Second(), time.Now().Second())
	// 取第3个属性, 并转为interface{}类型
	fieldGender := reflect.ValueOf(u).Field(2).Interface()
	assert.Equal(t, fieldGender.(gender).G, true)
}
