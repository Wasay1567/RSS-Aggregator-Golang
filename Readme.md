* RSS Aggregator
A simple RSS aggregator written in Go.

** Installation: **
1. Clone the repository:
   ```bash
   git clone https://github.com/wasay1567/rss-aggregator.git
   cd rss-aggregator
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Install Goose for database migrations:
   ```bash
    go install github.com/pressly/goose/v3/cmd/goose@latest
    ```
4. Install SQLC for generating Go code from SQL queries:
    ```bash
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    ```
5. Setting up postgres using Docker:
   ```bash
   docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=test -e POSTGRES_DB=rssagg -p 5432:5432 -d postgres:14.5
   ```
6. Set up the environment variables:
    Create a `.env` file in the root directory with the following content:
    ```plaintext
    PORT=8000
    DB_URL=postgres://root:test@localhost:5432/rssagg?sslmode=disable
    ``` 
7. Run the database migrations:
    ```bash
    goose up
    ```
8. Start the server:
    ```bash
    go run main.go
    ```
9. Access the API:
    Open your browser and navigate to `http://localhost:8000/v1/healthz` to check if the server is running.

** Usage: **
- The server will run on port 8000 by default.
- You can change the port by modifying the `PORT` variable in the `.env` file.
- The database connection URL can also be modified in the `.env` file.


