package domain

type User struct {
	ID                         int
	Email, FirstName, LastName string
	Terms                      []Term
}
