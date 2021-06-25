package main

import "testing"

func TestHasher(t *testing.T) {
	testCases := []struct{
		url string
		errorExists bool
	}{
		{"https://google.com", false},
		{"https://www.facebook.com", false},
		{"https://www.adjust.com", false},
		{"boomer", true},
	}

	for _, tc := range testCases {
		_, err := Hasher(tc.url)
		if err != nil {
			if !tc.errorExists {
				t.Errorf("Expected: No Error \t Got: %s", err.Error())
				continue
			}
		} else {
			if tc.errorExists {
				t.Error("Expected: Error \t Got: No Error")
				continue
			}
		}
	}
}
