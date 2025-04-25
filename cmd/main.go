package main

import (
	"atm-simulation/internal/transaction"
	"atm-simulation/internal/user"
	"atm-simulation/pkg/db"
	"fmt"

	"github.com/urfave/cli/v2"
)

// currentUser stores the account that is currently logged in
var currentUser *user.Account

// Displays the main menu of the ATM application
func mainMenu() {
	fmt.Println("\n===== Menu Utama =====")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Deposit")
	fmt.Println("5. Withdraw")
	fmt.Println("6. Transfer")
	if currentUser != nil {
		// Options available only after login
		fmt.Println("7. View Profile")
		fmt.Println("8. Change PIN")
		fmt.Println("9. Log Out")
		fmt.Println("10. View Transaction History")
	}
	fmt.Println("11. Exit")
}

// Handles user input and operations based on the menu choice
func handleChoice(c *cli.Context) {
	for {
		mainMenu() // Display the main menu
		fmt.Print("Pilih menu (1-11): ")
		var choice int
		_, err := fmt.Scanln(&choice)

		// Validate input choice
		if err != nil || choice < 1 || choice > 11 {
			fmt.Println("Pilihan tidak valid, coba lagi.")
			continue
		}

		// Call the corresponding function based on user choice
		switch choice {
		case 1:
			register()
		case 2:
			login()
		case 3:
			checkBalance()
		case 4:
			deposit()
		case 5:
			withdraw()
		case 6:
			transfer()
		case 7:
			viewProfile()
		case 8:
			changePIN()
		case 9:
			logOut()
		case 10:
			viewTransactionHistory()
		case 11:
			fmt.Println("Terima kasih telah menggunakan aplikasi ATM!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

// Registers a new account
func register() {
	fmt.Print("Masukkan nama: ")
	var name string
	fmt.Scanln(&name)
	fmt.Print("Masukkan PIN: ")
	var pin string
	fmt.Scanln(&pin)

	// Call the register function from the user package
	account, err := user.Register(name, pin)
	if err != nil {
		fmt.Println("Gagal membuat akun:", err)
		return
	}

	// Display account ID after successful registration
	fmt.Printf("Akun berhasil dibuat! ID Akun: %d\n", account.ID)
	fmt.Println("Kembali ke menu utama...\n")
}

// Logs in to the application
func login() {
	fmt.Print("Masukkan nama: ")
	var name string
	fmt.Scanln(&name)
	fmt.Print("Masukkan PIN: ")
	var pin string
	fmt.Scanln(&pin)

	// Call the login function from the user package
	account, err := user.Login(name, pin)
	if err != nil {
		fmt.Println("Login gagal:", err)
		return
	}

	currentUser = account
	fmt.Printf("Login berhasil! Selamat datang, %s.\n", account.Name)
	fmt.Println("Kembali ke menu utama...\n")
}

// Formats the currency with thousand separators (e.g., "Rp 1,000")
func formatCurrencyWithSeparator(amount float64) string {
	amountStr := fmt.Sprintf("%.0f", amount)
	result := ""
	count := 0

	// Format the amount to include thousand separators
	for i := len(amountStr) - 1; i >= 0; i-- {
		result = string(amountStr[i]) + result
		count++
		if count%3 == 0 && i > 0 {
			result = "." + result
		}
	}

	return "Rp " + result
}

// Checks the balance of the logged-in account
func checkBalance() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk melihat saldo.")
		return
	}

	// Get the balance of the current account
	balance, err := user.CheckBalance(currentUser.ID)
	if err != nil {
		fmt.Println("Gagal memeriksa saldo:", err)
		return
	}

	// Display the balance in currency format
	fmt.Printf("Saldo Anda saat ini: %s\n", formatCurrencyWithSeparator(balance))
	fmt.Println("Kembali ke menu utama...\n")
}

// Retrieves the updated balance after a transaction (deposit, withdraw, transfer)
func getUpdatedBalance(accountID int) (float64, error) {
	var balance float64
	err := db.DB.Get(&balance, "SELECT balance FROM accounts WHERE id = ?", accountID)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

// Deposits money into the account
func deposit() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk deposit.")
		return
	}

	var amount float64
	for {
		fmt.Print("Masukkan jumlah deposit: ")
		_, err := fmt.Scanln(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Jumlah uang tidak valid, coba lagi.")
			continue
		}
		break
	}

	// Call the deposit function from the transaction package
	err := transaction.Deposit(currentUser.ID, amount)
	if err != nil {
		fmt.Println("Gagal melakukan deposit:", err)
		return
	}

	// Retrieve the updated balance after deposit
	updatedBalance, err := getUpdatedBalance(currentUser.ID)
	if err != nil {
		fmt.Println("Gagal memeriksa saldo:", err)
		return
	}

	// Display the updated balance after deposit
	fmt.Printf("Deposit berhasil! Saldo Anda sekarang: %s\n", formatCurrencyWithSeparator(updatedBalance))
	fmt.Println("Kembali ke menu utama...\n")
}

// Withdraws money from the account
func withdraw() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk melakukan penarikan.")
		return
	}

	var amount float64
	for {
		fmt.Print("Masukkan jumlah penarikan: ")
		_, err := fmt.Scanln(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Jumlah uang tidak valid, coba lagi.")
			continue
		}
		break
	}

	// Call the withdraw function from the transaction package
	err := transaction.Withdraw(currentUser.ID, amount)
	if err != nil {
		fmt.Println("Gagal melakukan penarikan:", err)
		return
	}

	// Retrieve the updated balance after withdrawal
	updatedBalance, err := getUpdatedBalance(currentUser.ID)
	if err != nil {
		fmt.Println("Gagal memeriksa saldo:", err)
		return
	}

	// Display the updated balance after withdrawal
	fmt.Printf("Penarikan berhasil! Saldo Anda sekarang: %s\n", formatCurrencyWithSeparator(updatedBalance))
	fmt.Println("Kembali ke menu utama...\n")
}

// Transfers money to another account
func transfer() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk melakukan transfer.")
		return
	}

	// Enter the target account ID
	fmt.Print("Masukkan ID akun tujuan: ")
	var targetID int
	fmt.Scanln(&targetID)

	// Check if the target account exists
	var targetAccount user.Account
	err := db.DB.Get(&targetAccount, "SELECT id, name FROM accounts WHERE id = ?", targetID)
	if err != nil {
		// If the target account is not found
		fmt.Println("User ID tujuan tidak terdaftar.")
		return
	}

	// Display target account information
	fmt.Printf("Akun tujuan ditemukan: %s (ID: %d)\n", targetAccount.Name, targetAccount.ID)

	// Ask for user confirmation before transfer
	var confirm string
	fmt.Print("Apakah Anda yakin ingin mentransfer ke akun ini? (y/n): ")
	fmt.Scanln(&confirm)

	if confirm != "y" && confirm != "Y" {
		fmt.Println("Transfer dibatalkan. Kembali ke menu utama.")
		return
	}

	// Enter the transfer amount
	var amount float64
	for {
		fmt.Print("Masukkan jumlah transfer: ")
		_, err := fmt.Scanln(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Jumlah uang tidak valid, coba lagi.")
			continue
		}
		break
	}

	// Perform the transfer
	err = transaction.Transfer(currentUser.ID, targetID, amount)
	if err != nil {
		fmt.Println("Gagal melakukan transfer:", err)
		return
	}

	// Retrieve the updated balance after transfer
	updatedBalance, err := getUpdatedBalance(currentUser.ID)
	if err != nil {
		fmt.Println("Gagal memeriksa saldo:", err)
		return
	}

	// Display the updated balance after transfer
	fmt.Printf("Transfer berhasil! Saldo Anda sekarang: %s\n", formatCurrencyWithSeparator(updatedBalance))
	fmt.Println("Kembali ke menu utama...\n")
}

// Displays the profile of the logged-in account
func viewProfile() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk melihat profil.")
		return
	}

	// Display the account profile information
	fmt.Printf("\n===== Profil Akun =====\n")
	fmt.Printf("ID Akun: %d\n", currentUser.ID)
	fmt.Printf("Nama: %s\n", currentUser.Name)
	balance, err := user.CheckBalance(currentUser.ID)
	if err != nil {
		fmt.Println("Gagal memeriksa saldo:", err)
		return
	}
	fmt.Printf("Saldo: %s\n", formatCurrencyWithSeparator(balance))
	fmt.Println("Kembali ke menu utama...\n")
}

// Changes the PIN of the logged-in account
func changePIN() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk mengganti PIN.")
		return
	}

	// Ask for the old PIN to verify the user
	fmt.Print("Masukkan PIN lama: ")
	var oldPIN string
	fmt.Scanln(&oldPIN)

	// Verify the old PIN
	if currentUser.PIN != oldPIN {
		fmt.Println("PIN lama salah. Gagal mengganti PIN.")
		return
	}

	// Ask for the new PIN
	fmt.Print("Masukkan PIN baru: ")
	var newPIN string
	fmt.Scanln(&newPIN)

	// Update the PIN in the database
	err := user.ChangePIN(currentUser.ID, newPIN)
	if err != nil {
		fmt.Println("Gagal mengganti PIN:", err)
		return
	}

	fmt.Println("PIN berhasil diganti.")
}

// Displays transaction history based on type (deposit, withdrawal, etc.)
func viewTransactionHistory() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk melihat riwayat transaksi.")
		return
	}

	// Display transaction type options
	fmt.Println("\n===== Pilih Jenis Transaksi =====")
	fmt.Println("1. Transferan Masuk")
	fmt.Println("2. Transferan Keluar")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Deposit")
	fmt.Println("5. Kembali ke menu utama")

	var choice int
	fmt.Print("Pilih menu (1-5): ")
	_, err := fmt.Scanln(&choice)

	if err != nil || choice < 1 || choice > 5 {
		fmt.Println("Pilihan tidak valid, coba lagi.")
		return
	}

	var transactionType string
	switch choice {
	case 1:
		transactionType = "transfer_in"
	case 2:
		transactionType = "transfer_out"
	case 3:
		viewWithdrawHistory()
		return
	case 4:
		viewDepositHistory()
		return
	case 5:
		return
	}

	transactions, err := transaction.ViewTransactionHistory(currentUser.ID, transactionType)
	if err != nil {
		fmt.Println("Gagal memuat riwayat transaksi:", err)
		return
	}

	// Display transaction history
	fmt.Println("\n===== Riwayat Transaksi =====")
	for _, trans := range transactions {
		tType := trans["type"].(string)
		amount := trans["amount"].(float64)
		targetID := trans["target_id"]
		createdAt := trans["created_at"].(string)

		// Display transaction information
		fmt.Printf("Tipe Transaksi: %s\n", tType)
		fmt.Printf("Jumlah: %s\n", formatCurrencyWithSeparator(amount))
		fmt.Printf("Tanggal: %s\n", createdAt)

		// Handle targetID which is an interface{} and perform type assertion
		if targetID != nil {
			if tID, ok := targetID.(*int); ok {
				// For transfer transactions, show the recipient's account name
				if tType == "transfer_in" || tType == "transfer_out" {
					var targetName string
					err := db.DB.Get(&targetName, "SELECT name FROM accounts WHERE id = ?", *tID)
					if err != nil {
						fmt.Println("Gagal mendapatkan nama akun tujuan:", err)
					} else {
						fmt.Printf("Nama: %s\n", targetName)
					}
				}

				fmt.Printf("Target ID: %d\n", *tID) // Dereference pointer if successful
			} else {
				fmt.Println("Error: targetID type mismatch.")
			}
		}

		fmt.Println("-----------------------------------")
	}
	fmt.Println("Kembali ke menu utama...\n")
}

// Displays deposit transaction history
func viewDepositHistory() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk melihat riwayat transaksi deposit.")
		return
	}

	// Display deposit transaction history
	transactions, err := transaction.ViewTransactionHistory(currentUser.ID, "deposit")
	if err != nil {
		fmt.Println("Gagal memuat riwayat transaksi deposit:", err)
		return
	}

	fmt.Println("\n===== Riwayat Transaksi Deposit =====")
	for _, trans := range transactions {
		amount := trans["amount"].(float64)
		createdAt := trans["created_at"].(string)

		// Display deposit transaction information
		fmt.Printf("Tipe Transaksi: Deposit\n")
		fmt.Printf("Jumlah: %s\n", formatCurrencyWithSeparator(amount))
		fmt.Printf("Tanggal: %s\n", createdAt)
		fmt.Println("-----------------------------------")
	}
	fmt.Println("Kembali ke menu utama...\n")
}

// Displays withdrawal transaction history
func viewWithdrawHistory() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu untuk melihat riwayat transaksi withdraw.")
		return
	}

	// Display withdrawal transaction history
	transactions, err := transaction.ViewTransactionHistory(currentUser.ID, "withdraw")
	if err != nil {
		fmt.Println("Gagal memuat riwayat transaksi withdraw:", err)
		return
	}

	fmt.Println("\n===== Riwayat Transaksi Withdraw =====")
	for _, trans := range transactions {
		amount := trans["amount"].(float64)
		createdAt := trans["created_at"].(string)

		// Display withdrawal transaction information
		fmt.Printf("Tipe Transaksi: Withdraw\n")
		fmt.Printf("Jumlah: %s\n", formatCurrencyWithSeparator(amount))
		fmt.Printf("Tanggal: %s\n", createdAt)
		fmt.Println("-----------------------------------")
	}
	fmt.Println("Kembali ke menu utama...\n")
}

// Logs out of the application
func logOut() {
	if currentUser == nil {
		fmt.Println("Anda belum login.")
		return
	}
	currentUser = nil
	fmt.Println("Anda telah berhasil log out.")
	fmt.Println("Kembali ke menu utama...\n")
}

// Main function to run the ATM application
func main() {
	// Initialize database connection
	db.InitDB()

	// Start the application with interactive menu
	handleChoice(nil)
}
