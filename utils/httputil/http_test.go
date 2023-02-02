package httputil

import "testing"

func TestIsValidURL(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "normal", args: args{urlStr: "https://kubeclipper.io"}, want: true},
		{name: "withoutScheme", args: args{urlStr: "kubeclipper.io"}, want: false},
		{name: "empty", args: args{urlStr: ""}, want: false},
		{name: "special", args: args{urlStr: "//kubeclipper.io"}, want: false},
		{name: "normalIP", args: args{urlStr: "http://192.168.1.1:8080/api/v1/healthz"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidURL(tt.args.urlStr); got != tt.want {
				t.Errorf("IsValidURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
