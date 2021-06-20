package mta

import (
	"testing"
)

func TestSplitAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name       string
		args       args
		wantLocal  string
		wantDomain string
		wantErr    bool
	}{
		{name: "1", args: args{addr: "local@domain.com"}, wantDomain: "domain.com", wantLocal: "local", wantErr: false},
		{name: "2", args: args{addr: "1033256636@qq.com"}, wantDomain: "qq.com", wantLocal: "1033256636", wantErr: false},
		{name: "3", args: args{addr: "local.domain.com"}, wantDomain: "", wantLocal: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocal, gotDomain, err := SplitAddress(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLocal != tt.wantLocal {
				t.Errorf("SplitAddress() gotLocal = %v, want %v", gotLocal, tt.wantLocal)
			}
			if gotDomain != tt.wantDomain {
				t.Errorf("SplitAddress() gotDomain = %v, want %v", gotDomain, tt.wantDomain)
			}
		})
	}
}
