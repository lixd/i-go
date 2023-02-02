package health_check

import (
	"testing"
	"time"
)

func TestReachableByTCP(t *testing.T) {
	type args struct {
		protocol string
		host     string
		port     string
		timeout  time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "localhost", args: args{
				host:    "172.20.150.100",
				port:    "22",
				timeout: time.Second * 3,
			}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReachableByTCP(tt.args.host, tt.args.port, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("ReachableByTCP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReachableByPing(t *testing.T) {
	type args struct {
		addr    string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "localhost", args: args{
			addr:    "127.0.0.1",
			timeout: time.Second * 3,
		}, wantErr: false},
		{name: "any", args: args{
			addr:    "1.1.1.1",
			timeout: time.Second * 3,
		}, wantErr: false},
		{name: "error", args: args{
			addr:    "127.0.0.2",
			timeout: time.Second * 3,
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReachableByPing(tt.args.addr, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("ReachableByPing() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReachableByHTTP(t *testing.T) {
	type args struct {
		protocol string
		host     string
		port     string
		timeout  time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "baidu",
			args: args{
				protocol: "http",
				host:     "www.baidu.com",
				port:     "",
				timeout:  time.Second * 3,
			},
			wantErr: false,
		},
		{
			name: "github",
			args: args{
				protocol: "https",
				host:     "github.com",
				port:     "443",
				timeout:  time.Second * 3,
			},
			wantErr: false,
		},
		{
			name: "google",
			args: args{
				protocol: "https",
				host:     "google.com",
				port:     "443",
				timeout:  time.Second * 3,
			},
			wantErr: false,
		},
		{
			name: "loclahost",
			args: args{
				protocol: "http",
				host:     "127.0.0.1",
				port:     "8080",
				timeout:  time.Second * 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReachableByHTTP(tt.args.protocol, tt.args.host, tt.args.port, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("ReachableByHTTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
