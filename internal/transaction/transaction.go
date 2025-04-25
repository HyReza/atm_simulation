package transaction

import (
	"atm-simulation/pkg/db"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Checks if the account exists by querying the database
func checkAccountExists(accountID int) (bool, error) {
	var count int
	err := db.DB.Get(&count, "SELECT COUNT(*) FROM accounts WHERE id = ?", accountID)
	if err != nil {
		return false, err
	}
	// Returns true if the account exists, false otherwise
	return count > 0, nil
}

// Retrieves the transaction history based on account ID and transaction type
func ViewTransactionHistory(accountID int, transactionType string) ([]map[string]interface{}, error) {
	// Check if the account exists
	exists, err := checkAccountExists(accountID)
	if err != nil {
		return nil, err
	}
	if !exists {
		// If the account does not exist, return an error
		return nil, fmt.Errorf("user id tidak terdaftar")
	}

	// Prepare the query based on the transaction type
	var rows *sqlx.Rows // Use sqlx.Rows instead of db.Rows
	if transactionType == "all" {
		// If 'all', retrieve all transactions
		rows, err = db.DB.Queryx("SELECT type, amount, target_id, created_at FROM transactions WHERE account_id = ? ORDER BY created_at DESC", accountID)
	} else {
		// If a specific transaction type is given, filter by that type
		rows, err = db.DB.Queryx("SELECT type, amount, target_id, created_at FROM transactions WHERE account_id = ? AND type = ? ORDER BY created_at DESC", accountID, transactionType)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []map[string]interface{}

	// Iterate through each row and store the transaction in a map for easy access
	for rows.Next() {
		var tType, createdAt string
		var amount float64
		var targetID *int

		err := rows.Scan(&tType, &amount, &targetID, &createdAt)
		if err != nil {
			return nil, err
		}

		// Store the transaction in a map
		transaction := map[string]interface{}{
			"type":       tType,
			"amount":     amount,
			"target_id":  targetID,
			"created_at": createdAt,
		}

		transactions = append(transactions, transaction)
	}

	// If no transactions are found
	if len(transactions) == 0 {
		return nil, fmt.Errorf("tidak ada riwayat transaksi untuk kategori ini")
	}

	return transactions, nil
}

// Deposits money into the specified account
func Deposit(accountID int, amount float64) error {
	// Check if the account exists
	exists, err := checkAccountExists(accountID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user id tidak terdaftar")
	}

	// Update the account balance
	_, err = db.DB.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, accountID)
	if err != nil {
		return err
	}

	// Insert the deposit transaction into the transactions table
	_, err = db.DB.Exec(`INSERT INTO transactions (account_id, type, amount) VALUES (?, 'deposit', ?)`, accountID, amount)
	return err
}

// Withdraws money from the specified account
func Withdraw(accountID int, amount float64) error {
	// Check if the account exists
	exists, err := checkAccountExists(accountID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user id tidak terdaftar")
	}

	// Check if the account has enough balance
	var balance float64
	err = db.DB.Get(&balance, "SELECT balance FROM accounts WHERE id = ?", accountID)
	if err != nil {
		return err
	}
	if balance < amount {
		return fmt.Errorf("saldo tidak mencukupi")
	}

	// Update the account balance by subtracting the withdrawal amount
	_, err = db.DB.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, accountID)
	if err != nil {
		return err
	}

	// Insert the withdrawal transaction into the transactions table
	_, err = db.DB.Exec(`INSERT INTO transactions (account_id, type, amount) VALUES (?, 'withdraw', ?)`, accountID, amount)
	return err
}

// Transfers money between two accounts (sender and receiver)
func Transfer(accountID, targetID int, amount float64) error {
	// Check if the sender account exists
	exists, err := checkAccountExists(accountID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user id tidak terdaftar")
	}

	// Check if the target account exists
	exists, err = checkAccountExists(targetID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user id tujuan tidak terdaftar")
	}

	// Check if the sender has enough balance to transfer
	var balance float64
	err = db.DB.Get(&balance, "SELECT balance FROM accounts WHERE id = ?", accountID)
	if err != nil {
		return err
	}
	if balance < amount {
		return fmt.Errorf("saldo tidak mencukupi")
	}

	// Withdraw the amount from the sender's account
	_, err = db.DB.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, accountID)
	if err != nil {
		return err
	}

	// Deposit the amount into the target account
	_, err = db.DB.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, targetID)
	if err != nil {
		return err
	}

	// Insert the transaction for the sender
	_, err = db.DB.Exec(`INSERT INTO transactions (account_id, type, amount, target_id) VALUES (?, 'transfer_out', ?, ?)`, accountID, amount, targetID)
	if err != nil {
		return err
	}

	// Insert the transaction for the receiver
	_, err = db.DB.Exec(`INSERT INTO transactions (account_id, type, amount, target_id) VALUES (?, 'transfer_in', ?, ?)`, targetID, amount, accountID)
	return err
}
