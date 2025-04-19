package signup

import (
	"testing"
)

func TestRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request Request
		wantErr bool
	}{
		{
			name: "Valid Request",
			request: Request{
				Name:        "Test Name",
				Description: "Test Description",
			},
			wantErr: false,
		},
		{
			name: "Missing Name",
			request: Request{
				Description: "Test Description",
			},
			wantErr: true,
		},
		{
			name: "Missing Description",
			request: Request{
				Name: "Test Name",
			},
			wantErr: true,
		},
		{
			name:    "Missing Both",
			request: Request{},
			wantErr: true,
		},
		{
			name: "Name with Spaces",
			request: Request{
				Name:        "  Test Name  ",
				Description: "Test Description",
			},
			wantErr: false,
		},
		{
			name: "Description with Spaces",
			request: Request{
				Name:        "Test Name",
				Description: "  Test Description  ",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.request.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Request.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
