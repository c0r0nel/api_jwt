package main

import (
	"github.com/go-chi/jwtauth"
	"testing"
)

var (
	TokenAuthHS256 *jwtauth.JWTAuth
	TokenSecret    = []byte("secretpass")

	TokenAuthRS256 *jwtauth.JWTAuth

	PrivateKeyRS256String = `-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBALxo3PCjFw4QjgOX06QCJIJBnXXNiEYwDLxxa5/7QyH6y77nCRQy
J3x3UwF9rUD0RCsp4sNdX5kOQ9PUyHyOtCUCAwEAAQJARjFLHtuj2zmPrwcBcjja
IS0Q3LKV8pA0LoCS+CdD+4QwCxeKFq0yEMZtMvcQOfqo9x9oAywFClMSlLRyl7ng
gQIhAOyerGbcdQxxwjwGpLS61Mprf4n2HzjwISg20cEEH1tfAiEAy9dXmgQpDPir
C6Q9QdLXpNgSB+o5CDqfor7TTyTCovsCIQDNCfpu795luDYN+dvD2JoIBfrwu9v2
ZO72f/pm/YGGlQIgUdRXyW9kH13wJFNBeBwxD27iBiVj0cbe8NFUONBUBmMCIQCN
jVK4eujt1lm/m60TlEhaWBC3p+3aPT2TqFPUigJ3RQ==
-----END RSA PRIVATE KEY-----
`

	PublicKeyRS256String = `-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALxo3PCjFw4QjgOX06QCJIJBnXXNiEYw
DLxxa5/7QyH6y77nCRQyJ3x3UwF9rUD0RCsp4sNdX5kOQ9PUyHyOtCUCAwEAAQ==
-----END PUBLIC KEY-----
`
)

func init() {
	TokenAuthHS256 = jwtauth.New("HS256", TokenSecret, nil)
}

//func Test_main(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test cases.
//	}
//	for range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			main()
//		})
//	}
//}
func Test_validate(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid user & pass",
			args: args{
				username: "testuser",
				password: "supersecret",
			},
			want: true,
		},
		{
			name: "Invalid password",
			args: args{
				username: "testuser",
				password: "supcret",
			},
			want: false,
		},
		{
			name: "Invalid user",
			args: args{
				username: "teser",
				password: "supersecret",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateJwt(t *testing.T) {
	type args struct {
		username   string
		expiration int64
	}
	tests := []struct {
		name string
		args args
		want string
		pass bool
	}{
		{
			name: "simple check",
			args: args{
				username:   "foo",
				expiration: 1234,
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEyMzQsInVzZXJfaWQiOiJmb28ifQ.K5j9LgmmT9kQMvK7IKYO6wVEiZDKlBV-z1vxR77vDAw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateJwt(tt.args.username, tt.args.expiration); got != tt.want {
				t.Errorf("generateJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}
