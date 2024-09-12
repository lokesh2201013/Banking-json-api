package main


import ("math/rand"
 "time"
)

type CreateAccountRequest struct{
	FirstName string `json:"firstname"`
      LastName string `json:"lastname"`
}

// Account represents a bank account with ID, FirstName, LastName, Number, and Balance.
type Account struct {
	ID        int    `json:"id"`        // ID of the account
	FirstName string `json:"firstname"` // First name of the account holder
	LastName  string `json:"lastname"`  // Last name of the account holder
	Number    int64  `json:"number"`    // Unique number of the account
	CreatedAt time.Time `json:"createdAt"`
	Balance   int64  `json:"balance"`   // Account balance
}

// NewAccount is a constructor function that creates a new Account with random ID and account number.
// It takes the first name and last name as input and initializes the account fields.
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}
