package auth

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoPublicService struct {
	PublicKey  paseto.V4AsymmetricPublicKey
	PrivateKey paseto.V4AsymmetricSecretKey
}

func NewPasetoPublicService(pub paseto.V4AsymmetricPublicKey, priv paseto.V4AsymmetricSecretKey) *PasetoPublicService {
	return &PasetoPublicService{
		PublicKey:  pub,
		PrivateKey: priv,
	}
}

func (p *PasetoPublicService) GenerateToken(userID, role string) (string, error) {
	token := paseto.NewToken()
	token.SetString("user_id", userID)
	token.SetString("role", role)
	token.SetIssuedAt(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	return token.V4Sign(p.PrivateKey, nil), nil
}

func (p *PasetoPublicService) VerifyToken(tokenStr string) (*paseto.Token, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))

	return parser.ParseV4Public(p.PublicKey, tokenStr, nil)
}
