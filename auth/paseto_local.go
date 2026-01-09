package auth

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoService struct {
	key paseto.V4SymmetricKey
}

func NewPasetoService(key paseto.V4SymmetricKey) *PasetoService {
	return &PasetoService{key: key}
}

func (p *PasetoService) GenerateToken(userID, role string) (string, error) {
	token := paseto.NewToken()

	token.SetString("user_id", userID)
	token.SetString("role", role)

	token.SetIssuedAt(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	return token.V4Encrypt(p.key, nil), nil
}

func (p *PasetoService) VerifyToken(tokenStr string) (*paseto.Token, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))

	return parser.ParseV4Local(p.key, tokenStr, nil)
}
