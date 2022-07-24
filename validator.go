package jwt_licence

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type LicenceValidator struct {
	PublicKey *rsa.PublicKey
}

func NewLicenceValidator(publicKeyPath string) (*LicenceValidator, error) {
	publicKeyData, readErr := os.ReadFile(publicKeyPath)
	if readErr != nil {
		return nil, readErr
	}

	publicKey, parseErr := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if parseErr != nil {
		return nil, parseErr
	}

	return &LicenceValidator{PublicKey: publicKey}, nil
}

func (lv *LicenceValidator) ValidateLicence(licence string) (*LicenceData, error) {
	licenceBytes, decodingErr := base64.StdEncoding.DecodeString(licence)
	if decodingErr != nil {
		return nil, decodingErr
	}

	claims := &jwt.RegisteredClaims{}
	licenceJWT, jwtErr := jwt.ParseWithClaims(string(licenceBytes), claims, func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return lv.PublicKey, nil
	})

	if !licenceJWT.Valid {
		return nil, jwtErr
	}

	return &LicenceData{
		ExpiresAt: claims.ExpiresAt.Time,
		Subject:   claims.Subject,
		Issuer:    claims.Issuer,
	}, nil
}
