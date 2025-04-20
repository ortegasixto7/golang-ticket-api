package signup

import (
	"testing"
)

func TestExecute_Validation(t *testing.T) {
	tests := []struct {
		name    string
		request *Request
		wantErr bool
	}{
		{
			name: "Valid Request",
			request: &Request{
				Name:        "Test Name",
				Description: "Test Description",
			},
			wantErr: false,
		},
		{
			name: "Invalid Request - Missing Name",
			request: &Request{
				Description: "Test Description",
			},
			wantErr: true,
		},
		{
			name: "Invalid Request - Missing Description",
			request: &Request{
				Name: "Test Name",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Execute(tt.request, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
