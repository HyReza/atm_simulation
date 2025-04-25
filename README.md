# ATM Simulation Application

This project is an ATM simulation application developed in **Go (Golang)**, designed to simulate basic banking functionalities such as registration, login, balance checking, deposits, withdrawals, transfers, transaction history, and more.

## Features

- **User Registration & Login**: Allows users to create accounts with a PIN and log in securely.
- **Balance Checking**: Users can check their account balance in Indonesian Rupiah (Rp).
- **Deposit and Withdraw**: Users can deposit and withdraw money from their accounts, with checks for sufficient balance.
- **Transfer**: Users can transfer money between accounts with transaction history tracking.
- **Transaction History**: Users can view their transaction history, including deposits, withdrawals, and transfers.
- **PIN Change**: Users can change their PIN after confirming the old PIN.
- **Account Deletion**: Accounts can be deleted if the balance is zero, with a confirmation prompt.
- **Formatted Currency**: The balance is displayed in a user-friendly format with thousand separators (e.g., "Rp 1.000").

## Installation

### Prerequisites

1. **Go (Golang)**: Make sure you have Go installed. You can download and install it from the official website: [https://golang.org/dl/](https://golang.org/dl/).
2. **Database**: The application uses **SQLite** (or another relational database) to store user and transaction data.

### Steps

1. **Clone the repository**:

    ```bash
    git clone https://github.com/your-username/atm-simulation.git
    cd atm-simulation
    ```

2. **Install dependencies**:

    If you’re using Go modules (recommended), run the following command to install dependencies:

    ```bash
    go mod tidy
    ```

3. **Set up the database**:

    Ensure that the SQLite database (`db` folder) exists and contains the necessary tables. The application handles database connections and queries.

4. **Run the application**:

    Use the following command to run the application:

    ```bash
    go run main.go
    ```

    This will start the application with an interactive terminal menu.

## Usage

Once the application is running, you’ll see an interactive menu with the following options:

1. **Register**: Create a new user account by providing a name and PIN.
2. **Login**: Log in using your username and PIN.
3. **Check Balance**: View your current account balance.
4. **Deposit**: Deposit money into your account.
5. **Withdraw**: Withdraw money from your account.
6. **Transfer**: Transfer money to another account.
7. **View Profile**: View your account profile and balance.
8. **Change PIN**: Change your PIN after entering the old PIN.
9. **Log Out**: Log out of the current account.
10. **View Transaction History**: View the history of your transactions (deposits, withdrawals, transfers).
11. **Exit**: Exit the application.

## Code Structure

- **`main.go`**: Contains the main logic for running the application and handling user input.
- **`user.go`**: Handles user-related operations such as registration, login, and PIN management.
- **`transaction.go`**: Handles transaction-related operations including deposits, withdrawals, transfers, and transaction history.
- **`db/`**: Contains the database connection and schema management (SQLite).

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request if you have improvements or bug fixes.

### How to contribute:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature/your-feature`).
6. Open a pull request.

## License

This project is open-source and available under the [MIT License](LICENSE).

---

## Contact

If you have any questions or suggestions, feel free to open an issue on the repository or contact me directly.
