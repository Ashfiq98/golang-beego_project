# Golang Beego Project - Cat Voting and Favourites

## Project Overview

This project is a web application built using the [Beego](https://beego.vip/) framework and Go programming language. It implements a voting system for cat images, fetching data from [The Cat API](https://thecatapi.com). The frontend interactions are powered by a JavaScript library/framework or vanilla JS, as per preference.

Key features include:
- Voting (upvote/downvote) functionality.
- Marking and managing favourites.
- Viewing vote history.
- API integration for backend communication.

### Project Repository
Clone the project from the repository:
```
git clone https://github.com/Ashfiq98/golang-beego_project.git
```

### Web Application URL
Access the application at:
```
http://localhost:8080/
```

## Features

1. **Voting System**
   - Upvote API: `POST http://localhost:8080/vote/up`
   - Downvote API: `POST http://localhost:8080/vote/down`
   - Vote History API: `GET http://localhost:8080/vote/history`

2. **Favourites Management**
   - View Favourites: `GET http://localhost:8080/favourites`
   - Delete Favourites.

3. **Backend Integration**
   - Fetches data from The Cat API using Go channels.
   - Stores the data in a custom API for frontend interaction.

4. **Frontend Interaction**
   - Interacts with the backend through APIs to display and manage data dynamically.

5. **Code Coverage**
   - Achieves approximately 60% unit test coverage.

## Project Structure
```
|-- conf
|   |-- app.conf            # Application configuration file
|-- controllers
|   |-- catController.go    # Controller for managing cat-related functionalities
|   |-- default.go          # Default controller
|-- routers
|   |-- router.go           # Application routes
|-- static
|   |-- css
|   |   |-- styles.css      # Styling for the application
|   |-- js
|       |-- main.js         # Main JavaScript logic
|       |-- reload.min.js   # JavaScript for hot-reloading
|-- tests
|   |-- catController_test.go  # Unit tests for cat controller
|-- views
|   |-- index.tpl           # HTML template for the main page
|-- coverage                # Code coverage reports
|-- go.mod                  # Go module configuration
|-- go.sum                  # Go dependencies
|-- help.txt                # Instructions for using the application
|-- main.go                 # Main entry point of the application
```

## Setup Instructions

### Prerequisites
- Go 1.19 or higher installed.
- Beego framework installed.
- MySQL or SQLite (if using a database for additional persistence).

### Beego Installation
1. Install Beego:
   ```
   go get github.com/beego/beego/v2
   ```

2. Install Bee CLI tool (optional for development):
   ```
   go install github.com/beego/bee/v2@latest
   ```

### Project Setup
1. Clone the repository:
   ```
   git clone https://github.com/Ashfiq98/golang-beego_project.git
   cd golang-beego_project
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Configure `app.conf`:
   - Add your API key for The Cat API.
   - Configure other settings like port and database connection (if needed).

4. Run the application:
   ```
   bee run
   ```
   or
   ```
   go run main.go
   ```

### Running Tests
Run unit tests to check functionality and code coverage:
```sh
go test ./... -cover
```

## How It Works

1. **Backend**
   - Fetches cat data from The Cat API.
   - Processes requests using Go channels.
   - Provides APIs for voting and managing favourites.

2. **Frontend**
   - Dynamically displays data fetched from backend APIs.
   - Handles user interactions (voting, favourites) through JavaScript.

## Additional Notes

- **Configuration Management**: The Beego `app.conf` file is used for managing API keys and other configurations.
- **Code Quality**: Unit tests are written for the main controller functionalities to ensure reliability.
- **API Interaction**: All frontend interactions with the backend are through REST APIs.

### Troubleshooting
- Ensure your Go environment variables are correctly set.
- Verify the `app.conf` file for valid API key and configurations.
- Use the Bee CLI tool for easier debugging and running the application.

For further details, refer to the official [Beego documentation](https://beego.vip/).

