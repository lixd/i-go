package jwt

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(int64(10086))
	if err != nil {
		fmt.Printf("err: %v \n", err.Error())
		return
	}
	fmt.Println(token)
}
func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEwMDg2fQ.H2FuEzQ7_Fp7G-6i0-Sl6vBjr-yHcsk7RoC5Qk_rVgY"
	userId, err := ParseToken(tokenString)
	if err != nil {
		fmt.Printf("err: %v \n", err.Error())
		return
	}
	fmt.Println("userId: ", userId)
}

func BenchmarkGenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateToken(1)
	}
}

func BenchmarkParseToken(b *testing.B) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEyMzQ1fQ.vpZ2CycFtkbE63lSLyjG3y8mwtASRInPjynTh2be4Ks"
	for i := 0; i < b.N; i++ {
		ParseToken(tokenString)
	}
}

func TestParseToken1(t *testing.T) {
	type args struct {
		tokenString string
	}
	want1, _ := GenerateToken(1)
	want2, _ := GenerateToken(2)
	tests := []struct {
		name       string
		args       args
		wantUserId int64
		wantErr    bool
	}{
		{"1", args{tokenString: want1}, 1, false},
		{"2", args{tokenString: want2}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserId, err := ParseToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserId != tt.wantUserId {
				t.Errorf("ParseToken() gotUserId = %v, want %v", gotUserId, tt.wantUserId)
			}
		})
	}
}

func TestGenerateToken1(t *testing.T) {
	type args struct {
		userId int64
	}
	want1, _ := GenerateToken(int64(1))
	want2, _ := GenerateToken(int64(2))
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1", args{userId: 1}, want1, false},
		{"2", args{userId: 2}, want2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
