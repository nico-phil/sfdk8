package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/nico-phil/service/fondation/logger"
	"github.com/open-policy-agent/opa/rego"
)

var ErrForbidden = errors.New("attempted action is not allowed")

type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

func (c Claims) HasRole(r string) bool {
	for _, role := range c.Roles {
		if role == r {
			return true
		}
	}
	return false
}

type KeyLookup interface {
	PrivateKey(kid string) (key string, err error)
	PublicKey(kid string) (key string, err error)
}

type Config struct {
	Log       *logger.Logger
	KeyLookup KeyLookup
	Issuer    string
}

type Auth struct {
	keyLookup KeyLookup
	// userBus   *userbus.Core
	method jwt.SigningMethod
	parser *jwt.Parser
	issuer string
}

// New creates an Auth ot support authentication/authorization
func New(cfg Config) (*Auth, error) {
	// var userBus *user
	// if cfg.DB = nil {
	// 	userBus = userbus.NewCore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB))
	// }

	a := Auth{
		keyLookup: cfg.KeyLookup,
		method:    jwt.GetSigningMethod(jwt.SigningMethodES256.Name),
		parser:    jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodES256.Name})),
		issuer:    cfg.Issuer,
	}

	return &a, nil
}

// GenerateToken generates and signes token with a private key
func (a *Auth) GenerateToken(kid string, claims Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	token.Header["kid"] = kid

	privateKeyPEM, err := a.keyLookup.PrivateKey(kid)
	if err != nil {
		return "", fmt.Errorf("private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPEM))
	if err != nil {
		return "", fmt.Errorf("parsing private pem: %w", err)
	}

	str, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return str, nil
}

func (a *Auth) Authenticate(ctx context.Context, bearerToken string) (Claims, error) {
	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return Claims{}, errors.New("expected authurization header format: Bearer <token>")
	}

	var claims Claims
	token, _, err := a.parser.ParseUnverified(parts[1], &claims)
	if err != nil {
		return Claims{}, fmt.Errorf("error parsing token %w", err)
	}

	kidRaw, exists := token.Header["kid"]
	if !exists {
		return Claims{}, fmt.Errorf("kid missing from the header %w", err)
	}

	kid, ok := kidRaw.(string)
	if !ok {
		return Claims{}, fmt.Errorf("kid malformed %w", ok)
	}

	pem, err := a.keyLookup.PublicKey(kid)
	if err != nil {
		return Claims{}, fmt.Errorf("public file:  %w", err)
	}

	input := map[string]any{
		"Key":   pem,
		"Token": parts[1],
		"ISS":   a.issuer,
	}

	if err := a.opaPolicyEvaluation(ctx, regoAuthentication, RuleAuthenticate, input); err != nil {
		return Claims{}, fmt.Errorf("authentication failed %w", err)
	}

	// check the dabase for this user to veriry the are still anable

	return claims, nil

}

func (a *Auth) Authorize(ctx context.Context, claims Claims, userID uuid.UUID, rule string) error {
	input := map[string]any{
		"Roles":   claims.Roles,
		"Subject": claims.Subject,
		"UserID":  userID,
	}

	if err := a.opaPolicyEvaluation(ctx, regoAuthorization, rule, input); err != nil {
		return fmt.Errorf("rego evaluation failed : %w", err)
	}

	return nil
}

// opaPolicyEvaluation asks opa to evaluate the token against the specified token
// policy and public key.
func (a *Auth) opaPolicyEvaluation(ctx context.Context, regoScript string, rule string, input any) error {
	query := fmt.Sprintf("x = data.%s.%s", opaPackage, rule)

	q, err := rego.New(
		rego.Query(query),
		rego.Module("policy.rego", regoScript),
	).PrepareForEval(ctx)
	if err != nil {
		return err
	}

	results, err := q.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	if len(results) == 0 {
		return errors.New("no results")
	}

	result, ok := results[0].Bindings["x"].(bool)
	if !ok || !result {
		return fmt.Errorf("bindings results[%v] ok[%v]", results, ok)
	}

	return nil
}
