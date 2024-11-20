package auth

import (
	"log"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {

	test_hash, err := HashPassword("Welcome1!")
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
	}

	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "match",
			args:    args{password: "Welcome1!", hash: test_hash},
			wantErr: false,
		},
		{
			name:    "no match",
			args:    args{password: "lk2jlkjsfl80020", hash: test_hash},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPasswordHash(tt.args.password, tt.args.hash); (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
