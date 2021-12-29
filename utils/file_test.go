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
		{name: "2", args: args{filename: "/usr/local/projects/tmp.json"}, want: "tmp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFilePrefix(tt.args.filename); got != tt.want {
				t.Errorf("GetFilePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileSuffix(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{filename: "tmp.txt"}, want: ".txt"},
		{name: "2", args: args{filename: "/usr/local/projects/tmp.json"}, want: ".json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileExt(tt.args.filename); got != tt.want {
				t.Errorf("GetFileExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
