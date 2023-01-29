package pkg

import (
	"net"
	"testing"
)

func TestIsPublicIP(t *testing.T) {
	tests := []struct {
		ip      string
		want    bool
		wantErr bool
	}{
		{
			ip:      "200.16.1.1",
			want:    true,
			wantErr: false,
		},
		{
			ip:      "192.168.1.1",
			want:    false,
			wantErr: false,
		},
		{
			ip:      "172.31.1.1",
			want:    false,
			wantErr: false,
		},
		{
			ip:      "10.0.0.1",
			want:    false,
			wantErr: false,
		},
		{
			ip:      "8.8.8.8",
			want:    true,
			wantErr: false,
		},
		{
			ip:      "localhost",
			want:    false,
			wantErr: true,
		},
		{
			ip:      "",
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		ip := net.ParseIP(tt.ip)
		if ip == nil && !tt.wantErr {
			t.Errorf("Failed to parse IP address %q", tt.ip)
			continue
		}
		if got := IsPublicIP(ip); got != tt.want {
			t.Errorf("IsPublicIP(%q) = %v, want %v", tt.ip, got, tt.want)
		}
	}
}
