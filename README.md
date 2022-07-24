# Simple JWT based licence

## Generating key pair
Private key:
```
$ openssl genrsa -out keys/id_rsa 4096
```

Public key:
```
$ openssl rsa -in keys/id_rsa -pubout -out keys/id_rsa.pub

```

## Generating licence
```go
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
```

## Validating licence
```go
validator, validatorErr := jwt_licence.NewLicenceValidator("keys/id_rsa.pub")
check(validatorErr)
licenceData, licenceErr := validator.ValidateLicence(licence)
check(licenceErr)
```