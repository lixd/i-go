package daily

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{pattern: "abba", s: "dog cat cat dog"}, want: true},
		{name: "2", args: args{pattern: "abba", s: "dog dog dog dog"}, want: false},
		{name: "3", args: args{pattern: "abba", s: "dog cat cat fish"}, want: false},
		{name: "4", args: args{pattern: "aaa", s: "cat cat cat cat"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern2(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
