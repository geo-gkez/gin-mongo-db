# Gin Go Framework with MongoDB

This project demonstrates basic implementations of the Gin Go framework with MongoDB.

## Prerequisites

- Go 1.16 or later
- MongoDB
- Git

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/gin-members.git
   cd gin-members
    ```

2. Install the dependencies:
    ```sh
      go mod tidy
    ```

3. Set up your MongoDB connection in the config package.

## Running the Application

1. Start the MongoDB server.
2. Run the application:
    ```sh
    go run main.go
    ```
3. The application will be available at http://localhost:8080.

## API Endpoints

- Create a Member

```json
URL: /members
Method: POST
Request Body: {
"name": "John",
"surname": "Doe",
"email": "john.doe@example.com"
}
Response: {
"id": "60c72b2f9b1e8a5f6d8b4567",
"name": "John",
"surname": "Doe",
"email": "john.doe@example.com"
}

Get All Members
URL: /members
Method: GET
Response:
[
{
"id": "60c72b2f9b1e8a5f6d8b4567",
"name": "John",
"surname": "Doe",
"email": "john.doe@example.com"
}
]
```

## Use Docker compose

Build the docker image:

 ```sh 
   docker build -t gin-members .
 ```

Run the docker compose:

 ```sh
   docker-compose up
 ```