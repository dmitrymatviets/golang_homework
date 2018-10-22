package model

type Animal struct {
	Passport             *AnimalPassport
	SafetyCertificate    *SafetyCertificate
	OwnershipCertificate *OwnershipCertificate
	Weight               float64
}
