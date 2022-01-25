package validate

import (
	"fmt"
	"testing"
)

func TestURL(t *testing.T) {
	fmt.Println(IsNumbers("123abc"))
	fmt.Println(IsNumbers("1234567890"))
}

func TestPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{phone: "13452340416"}, want: true},
		{name: "2", args: args{phone: "1345234041611"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Phone(tt.args.phone); got != tt.want {
				t.Errorf("Phone() = %v, want %v", got, tt.want)
			}
		})
	}
}
