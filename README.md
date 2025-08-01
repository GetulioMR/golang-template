# Go Template

A simple, customizable Go project template to kickstart your next Go application.

## Features

- Clean Go project structure
- Live reload with Air for development
- Docker Compose setup with MySQL database
- PHPMyAdmin for database management
- VS Code Dev Container configuration
- Environment variables configuration with `.env` file support
- Ready for containerized development

## Prerequisites
- Docker and Docker Compose
- VS Code with Dev Containers extension
- Git

## Getting Started

1. **Clone the repository:**
    ```sh
    git clone git@github.com:GetulioMR/golang-template.git
    cd golang-template
    ```

2. **Set up environment variables:**
    ```sh
    cp .env.example .env
    # Edit .env file with your preferred values
    ```

3. **Start the containers:**
    ```sh
    docker compose up -d
    ```

4. **Connect through the VS Code dev container**

5. **Install dependencies:**
    ```sh
    go mod tidy
    ```

6. **Run the application with live reload:**
    ```sh
    air
    ```

## Accessing the Application

Once the containers are running and the application is started:

- **Go Application:** http://localhost:3000 (or the port specified in `APP_PORT`)
- **PHPMyAdmin:** http://localhost:8080 (or the port specified in `PHPMYADMIN_PORT`)

## Development Workflow
1. Make changes to your Go code in the `app/` directory
2. Air will automatically reload the application
3. View changes at http://localhost:3000 (or your configured `APP_PORT`)

## Database Configuration

This template includes a MySQL database configured through environment variables.

### Environment Setup
The database and application configuration is managed through a `.env` file. Default values are provided, but you can customize them:

1. **Copy the example environment file:**
   ```sh
   cp .env.example .env
   ```

2. **Edit the `.env` file with your preferred values:**
   ```env
   # Database Configuration
   MYSQL_ROOT_PASSWORD=your_root_password
   MYSQL_DATABASE=your_database_name
   MYSQL_USER=your_database_user
   MYSQL_PASSWORD=your_database_password
   
   # Application Configuration
   APP_PORT=3000
   PHPMYADMIN_PORT=8080
   MYSQL_PORT=3306
   ```

### Default Database Credentials
If you use the provided `.env` file without changes:
- **Host:** `mysql` (container name) or `localhost:3306` (from host)
- **Database:** `myapp_db`
- **Username:** `myapp_user`
- **Password:** `myapp_password`
- **Root Password:** `rootpassword`

### Connecting to the Database

**From within the Go application (dev container):**
```go
// Example connection string using environment variables
dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", 
    os.Getenv("MYSQL_USER"), 
    os.Getenv("MYSQL_PASSWORD"), 
    os.Getenv("MYSQL_DATABASE"))
```

**From host machine:**
```bash
mysql -h localhost -P 3306 -u myapp_user -p myapp_db
# Use the password from your .env file
```

**Using PHPMyAdmin:**
- Access http://localhost:8080 in your browser (or the port specified in `PHPMYADMIN_PORT`)
- Login with the credentials from your `.env` file

### Security Notes
- The `.env` file is excluded from version control via `.gitignore`
- Always use `.env.example` as a template for required environment variables
- Never commit sensitive credentials to your repository

## Project Structure

```
.
├── .air.toml           # Air live reload configuration
├── .devcontainer/      # VS Code dev container setup
│   └── Dockerfile
├── .dockerignore       # Docker ignore patterns
├── .env                # Environment variables (not in version control)
├── .env.example        # Environment variables template
├── .gitignore          # Git ignore patterns
├── app/                # Application source code
│   └── main.go
├── tmp/                # Temporary files (generated by Air)
├── docker-compose.yml  # Docker Compose configuration
├── go.mod              # Go module definition
├── LICENSE             # License file
└── README.md           # Project documentation
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests.

## License

This project is licensed under the MIT License.