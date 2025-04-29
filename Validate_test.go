package vies

import "testing"

func TestClient_Validate(t *testing.T) {
	type args struct {
		euVat string
	}
	tests := []struct {
		name        string
		client      *Client
		args        args
		wantIsValid bool
		wantErr     bool
	}{
		{
			name:   "Existing EU VAT",
			client: New(),
			args: args{
				euVat: "PL7251868136",
			},
			wantIsValid: true,
			wantErr:     false,
		},
		{
			name:   "Non-existing EU VAT",
			client: New(),
			args: args{
				euVat: "PL0000000000",
			},
			wantIsValid: false,
			wantErr:     false,
		},
		{
			name:   "Non-existing country code",
			client: New(),
			args: args{
				euVat: "xx0000000000",
			},
			wantIsValid: false,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsValid, err := tt.client.Validate(tt.args.euVat)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("Client.Validate() = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}
