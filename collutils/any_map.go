package collutils

import (
	"reflect"
	"strconv"
	"strings"
)

// M represents any value map
type M map[string]Any

// Get value by key
func (m M) Get(key string) Any {
	return m[key]
}

// GetDefault returns value by key or default if not found
func (m M) GetDefault(key string, defaults Any) Any {
	value, ok := m[key]
	if !ok {
		return defaults
	}
	return value
}

// GetBool returns bool value by key or false if not found
func (m M) GetBool(key string) bool {
	value, _ := m[key].(bool)
	return value
}

// GetStringAndTrim returns trimmed string value by key or "" (empty string) if not found
func (m M) GetStringAndTrim(key string) string {
	value, _ := m[key].(string)
	return strings.TrimSpace(value)
}

// GetInt returns int value by key or 0 if not found
func (m M) GetInt(key string) int {
	value, _ := m[key].(int)
	return value
}

// ParseInt returns or parses string to int value by key. Returns 0 if failed
func (m M) ParseInt(key string) int {
	v, _ := m[key]

	if reflect.TypeOf(v).Kind() == reflect.String {
		value, err := strconv.Atoi(v.(string))
		if err != nil {
			return 0
		}
		return value
	}

	value, _ := v.(int)
	return value
}

// -----

// Pick returns an array of values by selected keys
func (m M) Pick(keys ...string) L {
	var data L
	for _, key := range keys {
		if v, ok := m[key]; ok {
			data = append(data, v)
		} else {
			data = append(data, nil)
		}
	}
	return data
}

// Extend assigns elements from extended
func (m M) Extend(extended M) M {
	for i, v := range extended {
		m[i] = v
	}
	return m
}

// -----

// Pluck extracts an array of selected values
func Pluck(ms []M, key string) L {
	var data L
	for _, m := range ms {
		v := m.Get(key)
		data = append(data, v)
	}
	return data
}
