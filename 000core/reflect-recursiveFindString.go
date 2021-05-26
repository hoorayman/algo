package main

import (
	"fmt"
	"reflect"
)

type foo struct {
	A string
	Bar
	M map[string]string
	S []string
}

type Bar struct {
	B string
	I interface{}
	P *struct{ O int }
}

func find(input interface{}, findString *string) bool {
	v := reflect.ValueOf(input)
	switch v.Kind() {
	case reflect.String:
		if v.String() == *findString { // string value
			return true
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if find(v.Type().Field(i).Name, findString) { // struct field
				return true
			}
			if find(v.Field(i).Interface(), findString) { // struct value
				return true
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if find(key.Interface(), findString) { // map key
				return true
			}
			value := v.MapIndex(key)
			if find(value.Interface(), findString) { // map value
				return true
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if find(v.Index(i).Interface(), findString) { // slice element
				return true
			}
		}
	case reflect.Interface, reflect.Ptr:
		if find(v.Elem().Interface(), findString) { // dereference pointer
			return true
		}
	}
	return false
}

func main() {
	t := map[int]*struct {
		Str string
		P   float64
	}{1: &struct {
		Str string
		P   float64
	}{"iamhere", 3.4}}

	f := foo{A: "abc", Bar: Bar{"cde", t, &struct{ O int }{9}},
		M: map[string]string{"abd": "akf"}, S: []string{"ok", "haha"}}

	findStr := "iamhere"

	fmt.Println(find(f, &findStr))
}
