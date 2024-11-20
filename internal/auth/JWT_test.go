package auth

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestValidateJWT(t *testing.T) {

	userID, _ := uuid.Parse("47fdffa3-37b9-4282-9dc5-9a1fd24a04d8")
	tokenSecret := "my-super-secret-key"
	expIn := time.Duration(time.Second * 1000)

	jwt, _ := MakeJWT(userID, tokenSecret, expIn)

	fmt.Println("Initial JWT:")
	fmt.Printf("%+v\n", jwt)

	type args struct {
		tokenString string
		tokenSecret string
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "test-valid-jwt",
			args: args{
				tokenString: jwt,
				tokenSecret: tokenSecret,
			},
			want:    userID,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateJWT(tt.args.tokenString, tt.args.tokenSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
