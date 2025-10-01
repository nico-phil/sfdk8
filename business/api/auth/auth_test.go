package auth_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/nico-phil/service/business/api/auth"
	"github.com/nico-phil/service/fondation/logger"
)

func Test_Auth(t *testing.T) {

	log, _ := newUnit(t)

	cfg := auth.Config{
		Log:       log,
		KeyLookup: &keyStore{},
		Issuer:    "service project",
	}

	a, err := auth.New(cfg)
	if err != nil {
		t.Fatalf("should be able to create an autenticator %v", err)
	}

	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "service project",
			Subject:   "5cf37266-3473-4006-984f-9325122678b7",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		Roles: []string{"ADMIN"},
	}

	token, err := a.GenerateToken(kid, claims)
	if err != nil {
		t.Fatalf("Should be able to generate a JWT: %v", err)
	}

	parsedClaims, err := a.Authenticate(context.Background(), "Bearer "+token)
	if err != nil {
		t.Fatalf("Should be able to authenticate the claims : %s", err)
	}

	userID := uuid.MustParse(claims.Subject)

	err = a.Authorize(context.Background(), parsedClaims, userID, auth.RuleAdminOnly)
	if err != nil {
		t.Errorf("Should be able to authorize the RoleAdmin claims : %s", err)
	}

	err = a.Authorize(context.Background(), parsedClaims, userID, auth.RuleUserOnly)
	if err == nil {
		t.Error("Should NOT be able to authorize the RoleUser claim")
	}

	err = a.Authorize(context.Background(), parsedClaims, userID, auth.RuleAdminOrSubject)
	if err != nil {
		t.Errorf("Should be able to authorize the RuleAdminOrSubject claim with RoleAdmin only : %s", err)
	}

}

func newUnit(t *testing.T) (*logger.Logger, func()) {
	var buf bytes.Buffer
	log := logger.New(&buf, logger.LevelInfo, "TEST", func(context.Context) string { return "00000000-0000-0000-0000-000000000000" })

	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()

		fmt.Println("******************** LOGS ********************")
		fmt.Print(buf.String())
		fmt.Println("******************** LOGS ********************")
	}

	return log, teardown
}

type keyStore struct{}

func (ks *keyStore) PrivateKey(kid string) (string, error) {
	return privateKeyPEM, nil
}

func (ks *keyStore) PublicKey(kid string) (string, error) {
	return publicKeyPEM, nil
}

const (
	kid = "s4sKIjD9kIRjxs2tulPqGLdxSfgPErRN1Mu3Hd9k9NQ"

	privateKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEuAIBAAKCAQUJgAC77cy851Sur1j2v8ontUX1uGliX7pWvfFqsrMoCVwCQZVH
C9LpbiIsp6FMiFSjEGP8nTW2P1NjhtTeSW6UN/WVk3cL8LEmaHfd+y7V7j21wUhE
7xs3UPG+DqVpeI7GCK+yH/JfleTO55q+Fmj8xXH5UD+JA42WLBk9MiXE+8ZSPd0C
1W9OzzgUJvX0SyPBDzBXs8QYpRLuPzYu2wxKXDqOvNAzmzMSIxaKlocF8gS9JKvJ
lyvl4fvM+dRYgOEgp2N9tlj0pevoZ7LCvLl4QUkaBIjGG/TFR8CY3AECllrmz4da
mGHzRSYdR2gyQRQp/tUhs1CRvPSQnAZP42l513/Lcr8CAwEAAQKCAQUBbgXuQtZ/
EdPxbIuXq3rZ6hLq7gPSXfT2DCVS+S00l+AEqOk2JaXrIvu6u8nvck9GsXdS9Dg6
wyIiRw8vm16sqRVpmyWeIu4OiUeNHbpJUU9xVPMzCMeQVjrj72Fe09mtHW6QAPXi
A/XJXBsyg73uWSsTqku9q8657M4poywDDx+kjSR8Q8oe7N17vsQP9xagrzmmYY1o
F5QPT7UCeR5sEJxlgx0Vs55EMPidwMXypw4k0DJzSsKneSpI44P5YemuU+2GlqCm
DBfCqHknpJ40+IYm4wrCSbnmrsSEgVnI3EG54HQeJcrYe/+PAfSHrlE+t5s+jCOL
x47goJz0f6Ec07uX96ECgYMDGQBlOh7IgbCB7JdpDznbegcXFVnZIzV30ez10/9Q
cz79/cdV6XceiL9btK7wsdJx9+zaaulITlXH0POKdbMHUQE8wflMllK0v1qK8HtQ
rio/vbCkDBeyE4f6CPFumC6i7ZZRlxAMjl0RZAFUS1GjnoK6I0ZbWn0z+hiXmInI
MrzWXwKBgwMRG+5mghS+PM+EB5s/oY/4j3fR74AOHfsLl8mov43MKWkho+Kpf6E+
re4lwkLZyqYOVtIdmTaNMKglz/mY7swA7YB5d5/vWCfGnrD9rESaYYIgPodIKozV
WlYynUgIU0kIoLViRY2KPdD9BjH3tGaHyfuMoGrw+pgqMLbtM/xygP+hAoGDAiFA
ulmaLPeva8ZH3X8Qoy5mjaKqorio3PhE3EqmNKTpXS8Pzqy2sRIJsX6tAubh8mcs
PopgWM00Ai5UJpIDRTaXiTU+u0BpIcqo8PulbrYyap19RW7jJBh12KApkYemGXUP
dI5QBLImn/yJRXJ8cepdrKrwh4PSrth57FJ/+l2YpjUCgYMA7GHQgaSgwjZk9Iel
wp7OTjRECz1k/Nsh/veQi/JAqzu8n5hMYmQ/FDQiA9RddF2DacXSNX8v4YrI1bms
mNMtMQpRKEFQMiwErdSRzY7UiPbaywKIkL3e8U3lrg+U5IzO7H4WnqP6XakHB2ea
G86BIFk8F7ck+7E3p1xLd1ezpnYYgQKBgwK6uON/MQQlf4qPCWQpbKHO8NgUw7SG
HbTKzBeBpjs/mc0f7sysCdRUYENYpe9WQCdses7PdmkwDkTmz2ba4hp2AF3RScFn
KRpVlIKX/No7SwzZY0ngBzhAdp25Zjveetc7XlAq+BuKaWhKVnuo7Xg0gtlpsXZU
d1ye2vTtIWWaHA5V
-----END PRIVATE KEY-----`

	publicKeyPEM = `-----BEGIN Public-----
MIIBJjANBgkqhkiG9w0BAQEFAAOCARMAMIIBDgKCAQUJgAC77cy851Sur1j2v8on
tUX1uGliX7pWvfFqsrMoCVwCQZVHC9LpbiIsp6FMiFSjEGP8nTW2P1NjhtTeSW6U
N/WVk3cL8LEmaHfd+y7V7j21wUhE7xs3UPG+DqVpeI7GCK+yH/JfleTO55q+Fmj8
xXH5UD+JA42WLBk9MiXE+8ZSPd0C1W9OzzgUJvX0SyPBDzBXs8QYpRLuPzYu2wxK
XDqOvNAzmzMSIxaKlocF8gS9JKvJlyvl4fvM+dRYgOEgp2N9tlj0pevoZ7LCvLl4
QUkaBIjGG/TFR8CY3AECllrmz4damGHzRSYdR2gyQRQp/tUhs1CRvPSQnAZP42l5
13/Lcr8CAwEAAQ==
-----END Public-----`
)
