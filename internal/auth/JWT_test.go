package auth

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestValidateJWT(t *testing.T) {

	/// TEST 1 - VALID JWT ///
	userID, _ := uuid.Parse("47fdffa3-37b9-4282-9dc5-9a1fd24a04d8")
	tokenSecret := "my-super-secret-key"
	expIn1 := time.Duration(time.Second * 1000)
	jwt, _ := MakeJWT(userID, tokenSecret, expIn1)

	/// TEST 2 - EXPIRED TOKEN ///
	expIn2 := time.Duration(time.Millisecond)
	jwt2, _ := MakeJWT(userID, tokenSecret, expIn2)

	/// TEST 3 - WRONG SECRET KEY ///
	wrongSecret := "making-this-up"
	jwt3, _ := MakeJWT(userID, wrongSecret, expIn1)

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
		{
			name: "test-expired-token",
			args: args{
				tokenString: jwt2,
				tokenSecret: tokenSecret,
			},
			want:    uuid.UUID{},
			wantErr: true,
		},
		{
			name: "test-wrong-secret-key",
			args: args{
				tokenString: jwt3,
				tokenSecret: tokenSecret,
			},
			want:    uuid.UUID{},
			wantErr: true,
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

func TestGetBearerToken(t *testing.T) {

	headers1 := http.Header{
		"Authorization": []string{"Bearer B10ijfjlkj393jslsjf2"},
	}

	headers2 := http.Header{
		"Content-Type": []string{"application/json"},
	}

	headers3 := http.Header{
		"Authorization": []string{"Bearer        B10ijfjlkj393jslsjf2"},
	}

	type args struct {
		headers http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Valid bearer token",
			args:    args{headers: headers1},
			want:    "B10ijfjlkj393jslsjf2",
			wantErr: false,
		},
		{
			name:    "No Authorization header included",
			args:    args{headers: headers2},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Lots of white space between 'Bearer' and token",
			args:    args{headers: headers3},
			want:    "B10ijfjlkj393jslsjf2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBearerToken(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBearerToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
