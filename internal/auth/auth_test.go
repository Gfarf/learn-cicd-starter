package auth

import (
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		header     string
		identifier string
		want       string
		errorWant  string
	}{
		"Correct Key":      {header: "Authorization", identifier: "ApiKey ItisaCorrectKey", want: "ItisaCorrectKey", errorWant: ""},
		"Incorrect Header": {header: "", identifier: "ApiKey ItisaCorrectKey", want: "", errorWant: "no authorization header"},
		"Malformed Key":    {header: "Authorization", identifier: "Api wrong key", want: "", errorWant: "malformed authorization header"},
		"Malformed Key 2 ": {header: "Authorization", identifier: "ApiKey", want: "", errorWant: "malformed authorization header"},
	}
	for name, tc := range tests {
		//just a comment
		t.Run(name, func(t *testing.T) {
			res := &http.Response{
				Header: http.Header{},
			}
			res.Header.Set(tc.header, tc.identifier)
			got, err := GetAPIKey(res.Header)
			if err != nil {
				if strings.Contains(err.Error(), tc.errorWant) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			} else {
				diff := cmp.Diff(tc.want, got)
				if diff != "" {
					t.Fatalf(diff)
				}
			}
		})
	}
}
