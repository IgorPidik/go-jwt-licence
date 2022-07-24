package jwt_licence

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type LicenceGenerator struct {
	PrivateKey *rsa.PrivateKey
}

func NewLicenceGenerator(privateKeyPath string) (*LicenceGenerator, error) {
	privateKeyData, readErr := os.ReadFile(privateKeyPath)
	if readErr != nil {
		return nil, readErr
	}

	privateKey, parseErr := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if parseErr != nil {
		return nil, parseErr
	}

	return &LicenceGenerator{PrivateKey: privateKey}, nil
}

func (lg *LicenceGenerator) CreateLicence(data *LicenceData) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(data.ExpiresAt),
		Subject:   data.Subject,
		Issuer:    data.Issuer,
	}

	licenceJWT, jwtErr := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(lg.PrivateKey)
	if jwtErr != nil {
		return "", jwtErr
	}

	return base64.StdEncoding.EncodeToString([]byte(licenceJWT)), nil
}
