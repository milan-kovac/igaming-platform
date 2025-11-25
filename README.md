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
