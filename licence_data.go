package jwt_licence

import "time"

type LicenceData struct {
	ExpiresAt time.Time
	Subject   string
	Issuer    string
}
