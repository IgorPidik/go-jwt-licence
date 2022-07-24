package main

import (
	"go-jwt-licence"
	"log"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func generateLicence() string {
	generator, generatorErr := jwt_licence.NewLicenceGenerator("keys/id_rsa")
	check(generatorErr)

	expiresAt, parsingErr := time.Parse("2006-01-02", "2022-12-31")
	check(parsingErr)

	licenceData := &jwt_licence.LicenceData{
		ExpiresAt: expiresAt,
		Subject:   "Licence subject",
		Issuer:    "Licence issuer",
	}

	licence, licenceErr := generator.CreateLicence(licenceData)
	check(licenceErr)
	return licence
}

func validateLicence(licence string) *jwt_licence.LicenceData {
	validator, validatorErr := jwt_licence.NewLicenceValidator("keys/id_rsa.pub")
	check(validatorErr)
	licenceData, licenceErr := validator.ValidateLicence(licence)
	check(licenceErr)
	return licenceData
}

func main() {
	licence := generateLicence()
	log.Println(licence)

	licenceData := validateLicence(licence)
	log.Println(licenceData)
}
