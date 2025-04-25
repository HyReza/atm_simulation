package user

import (
	"atm-simulation/pkg/db"
	"fmt"
	"log"
)

// Account represents a user's account in the system
type Account struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	PIN       string  `db:"pin"`
	Balance   float64 `db:"balance"`
	CreatedAt string  `db:"created_at"`
}

// Register creates a new account
// Checks if the username is already taken, and if so, returns an error
func Register(name, pin string) (*Account, error) {
	// Check if the username already exists in the database
	var existingAccount Account
	err := db.DB.Get(&existingAccount, "SELECT * FROM accounts WHERE name = ?", name)
	if err == nil {
		// If the username already exists, return an error
		return nil, fmt.Errorf("nama pengguna sudah terdaftar, silakan pilih nama lain")
	}

	// If the username is not taken, create a new account
	account := &Account{Name: name, PIN: pin, Balance: 0.0}
	result, err := db.DB.NamedExec(`INSERT INTO accounts (name, pin, balance) VALUES (:name, :pin, :balance)`, account)
	if err != nil {
		return nil, err
	}

	// Retrieve the ID of the newly created account
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Set the ID of the new account into the account object
	account.ID = int(lastID)
	return account, nil
}

// Login authenticates the user by checking their name and PIN
// If the account is found, it returns the account details, otherwise an error
func Login(name, pin string) (*Account, error) {
	account := &Account{}
	// Use sqlx Get to fetch account details by name and PIN
	err := db.DB.Get(account, "SELECT * FROM accounts WHERE name = ? AND pin = ?", name, pin)

	// Handle errors if the account is not found
	if err != nil {
		// If no account is found, return an error
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("akun tidak ditemukan")
		}
		// Log any other errors
		log.Fatal(err)
	}
	return account, nil
}

// CheckBalance retrieves the balance of the given account by its ID
func CheckBalance(accountID int) (float64, error) {
	var balance float64
	// Query the balance from the database
	err := db.DB.Get(&balance, "SELECT balance FROM accounts WHERE id = ?", accountID)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

// ChangePIN updates the PIN of the user account
// It receives the account ID and the new PIN as parameters
func ChangePIN(accountID int, newPIN string) error {
	// Update the PIN for the user in the database
	_, err := db.DB.Exec("UPDATE accounts SET pin = ? WHERE id = ?", newPIN, accountID)
	return err
}
