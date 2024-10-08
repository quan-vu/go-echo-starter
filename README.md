# Go Echo Starter

This is a starter project for Go Echo.

## Getting Started

Create Postgres database and run the following SQL to create the table.

```sql
-- Create a new database
CREATE DATABASE ontrack;

-- Create a new user
CREATE USER ontrack_user WITH ENCRYPTED PASSWORD 'password';

-- Grant all privileges on the new database to the new user
GRANT ALL PRIVILEGES ON DATABASE ontrack TO ontrack_user;
```

### Prerequisites

- Go 1.16

### Installation

1. Clone the repo
   ```sh
   git clone
   ```
2. Install dependencies
   ```sh
   go mod download
   ```
3. Run the server
   ```sh
   go run cmd/server/main.go
   ```
4. Visit `http://localhost:8080` in your browser

## Rest API

### Create a invoices

```bash
curl --location 'http://localhost:8080/api/v1/invoices' \
--header 'Content-Type: application/json' \
--data '{
    "customer_name": "John Doe",
    "amount": 100.00,
    "due_date": "2024-09-01T15:04:05Z",
    "status": "pending"
}'
```

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - [@your_twitter](https://twitter.com/your_twitter) -
