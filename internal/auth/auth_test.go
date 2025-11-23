package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	type test struct {
		name    string
		input   http.Header
		want    string
		wantErr error
	}
	tests := []test{
		{
			name:    "Valid API Key",
			input:   http.Header{"Authorization": []string{"ApiKey valid_api_key"}},
			want:    "valid_api_key",
			wantErr: nil,
		},
		{
			name:    "Missing Authorization Header",
			input:   http.Header{},
			want:    "",
			wantErr: errors.New("no authorization header included"),
		},
		{
			name:    "Malformed Authorization Header",
			input:   http.Header{"Authorization": []string{"InvalidHeader"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if got != tc.want {
				t.Errorf("GetApiKey() got = %v, want %v", got, tc.want)
			}

			if err != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("GetApiKey() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
