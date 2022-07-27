package entity

type Product struct {
	Id          string
	Name        string
	Description string
	Status      string
	MakerID     string
	Maker       User
	CheckerID   *string
	Checker     User
	SignerID    *string
	Signer      User
}
