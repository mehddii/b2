package main

import (
	"testing"
)

func TestInsertGet(t *testing.T) {
	s := NewSkipMap[string, int]()

	tests := []struct {
		Name     string
		Key      string
		Value    int
		Expected int
		Error    bool
	}{
		{
			Name:     "Key already exists",
			Key:      "Hello",
			Value:    1,
			Expected: 1,
		},
		{
			Name:     "Key does not exist",
			Key:      "World",
			Expected: 0,
			Error:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if !test.Error {
				s.Insert(test.Key, test.Value)
			}

			result, err := s.Get(test.Key)

			if result != test.Expected || (err != nil) != test.Error {
				t.Errorf("Got %q; wanted %q", result, test.Expected)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	s := NewSkipMap[string, int]()
	s.Insert("Hello", 1)

	tests := []struct {
		Name string
		Key  string
	}{
		{
			Name: "Key already exists",
			Key:  "Hello",
		},
		{
			Name: "Key does not exist",
			Key:  "World",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			s.Delete(test.Key)
			_, err := s.Get(test.Key)

			if err != nil {
				t.Errorf("Got nil; wanted error")
			}
		})
	}
}
