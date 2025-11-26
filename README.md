# **iGaming Platform**

### **Project Description**  

This application is a simple iGaming platform with the following functionalities:  

- **Management of players and their accounts**  
- **Organization of tournaments**  
- **Distribution of prizes after completed tournaments**  
- **Ranking of players based on their account balances**  

The application uses **Golang** for business logic and **MySQL** for the database.

### Project Structure

- `internal/domain` – entity definitions (`Player`, `Tournament`, `PlayerRanking`)  
- `internal/repositories` – database access, executing SQL queries and procedures  
- `internal/services` – business logic, prize distribution  
- `internal/api/handlers` – HTTP handlers for endpoints  
- `internal/api/routers` – API route definitions  
- `internal/config` – application configuration and environment setup (`env_config.go`)  
- `internal/database` – database connection and migrations  
- `internal/shared` – reusable components, including API response structures and business logic helpers (`fallacies/`, `responses/`)  
- `cmd` – application entry point (`main.go`) and documentation (`docs/`)  
- `docker-compose.yaml` – containers for MySQL and Adminer  

### Running the Application

**Copy** `.env.example` to `.env` and fill in the values (DB user, password, port, etc.)

**Start** the **Docker** databases: 
```bash
docker-compose up -d
```

**Run the Application** : 
```bash
cd cmd
go run main.go
```

### Accessing the API

API endpoints are available through the routes defined in `routers`.
If using Swagger, the documentation will be available at `/swagger/index.html`.

### API Endpoints

- GET `/players/ranking` – returns the player ranking based on account balance

- POST `/tournaments/{id}/distribute` – distributes prizes for a specific tournament

### API Testing

Run all tests in the project:


```bash
go test ./... -v
```

### What did I learn while working on this project?
I had not used the `RANK()` function before. I learned about the `RANK()` function through this project because it ranks rows in a result set according to some conditions, and I applied the knowledge in ranking players in tournaments.
### What were the main challenges and how were they resolved?
The biggest problem was testing the application services, particularly the issue with the mocked database. I solved this problem by using a setup like the dependency injection principle, where I could construct a mocked repository that communicates with the mock database and not the actual database.

### What performance considerations did I take into account and how did I handle edge cases?
Used a single `UPDATE ... CASE` statement instead of multiple separate queries to update player balances in one operation, improving performance and reducing complexity.

The entire prize distribution logic is executed within a transaction to ensure data consistency, if any query fails, the transaction rolls back.

I used `FOR UPDATE` locking on the tournament row to prevent concurrent prize distribution or new bets while the procedure runs.

Edge cases:

- Ensured that prizes cannot be distributed twice for the same tournament.

- Verified that the tournament has at least three participants before distributing prizes.

- The procedure signals clear errors if any of these conditions are violated.

### Why did I use CTEs (Common Table Expressions) and window functions?
I applied a Common Table Expression `(WITH ranked_players AS (.))` and the RANK() window function in order to assign each player a rank according to the sum of their account balance. RANK() functionality will make it easier to extract ranks and eliminate the need for writing complex sub-queries.

### How do I see the application of CTEs and window functions in more complex scenarios and real-world use cases?

CTEs and window functions such as `RANK()` help in practical work related to ranking, calculating sums, and verifying data in an efficient manner without complex sub-queries. These provide readability and performance gains when it comes to leader boards, reporting, and analytic pages.
