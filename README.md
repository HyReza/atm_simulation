
## Installation

### Prerequisites

1. **Go (Golang)**: Ensure that you have Go installed. You can download it from the official website: [https://golang.org/dl/](https://golang.org/dl/).
2. **MySQL Database**: You need MySQL to run the application. You can download and install it from the official website: [https://www.mysql.com/](https://www.mysql.com/).

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

    - You need to create a MySQL database for the application. You can use the following SQL script to create the necessary tables:

    ```sql
    CREATE DATABASE atm_simulation;

    USE atm_simulation;

    CREATE TABLE accounts (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        pin VARCHAR(10) NOT NULL,
        balance FLOAT NOT NULL DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE transactions (
        id INT AUTO_INCREMENT PRIMARY KEY,
        account_id INT NOT NULL,
        type VARCHAR(50) NOT NULL,
        amount FLOAT NOT NULL,
        target_id INT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (account_id) REFERENCES accounts(id)
    );
    ```

    - Update your MySQL credentials in the `pkg/db/db.go` file to match your MySQL configuration.

4. **Run the application**:

    Use the following command to run the application:

    ```bash
    go run cmd/main.go
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

- **`cmd/`**: Contains the entry point of the application.
  - **`main.go`**: The main file that runs the ATM simulation application.

- **`internal/`**: Holds the business logic for the application.
  - **`user/`**: Contains the logic related to user operations, including registration, login, and PIN management.
    - **`user.go`**: Contains functions for user account management.
  - **`transaction/`**: Contains the logic for managing transactions (deposit, withdraw, and transfer).
    - **`transaction.go`**: Contains functions for performing and recording transactions.

- **`pkg/`**: Contains reusable libraries or modules used by the application.
  - **`db/`**: Handles the connection to the MySQL database and query operations.
    - **`db.go`**: Manages the MySQL connection and queries.

- **`go.mod`**: Contains the module dependencies for Go projects.
- **`go.sum`**: Provides cryptographic hashes of module dependencies for verifying integrity.
- **`README.md`**: This file containing project description, setup instructions, and usage.

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
