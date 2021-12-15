package utils

import "testing"

func TestGetFilePrefix(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{filename: "tmp.txt"}, want: "tmp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFilePrefix(tt.args.filename); got != tt.want {
				t.Errorf("GetFilePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
