# Main

Here's the documentation for the `main.go` file:

### Main Application Documentation

The `main.go` file serves as the entry point for the application, initializing the server, setting up middleware, connecting to the database, performing migrations, and defining routes.

#### Initialization

* **Database Connection**: A connection to the PostgreSQL database is established using the `DBConnector` from the `backend/repository/db/postgres` package. Debug mode is enabled for the database connection.
* **Database Migration**: Database migration is performed to ensure that the necessary tables (`Wallet`, `Transaction`, `SubTransaction`) are created in the database. Any errors during migration result in a fatal log.

#### Middleware

* **Gin Default Middleware**: Default middleware provided by Gin, including logging and error handling.
* **CORS Middleware**: CORS middleware is configured to allow requests from specified origins (`http://localhost:8080`) and to allow credentials.
* **Session Middleware**: Gin session middleware is set up with a cookie store and configured with a session expiration time of 24 hours.

#### Routes

* **`POST /verify`**: Endpoint for verifying the authenticity of a user by validating their message signature against the provided message and signature in the request body.
* **`GET /nonce`**: Endpoint for generating a nonce, typically used for authentication and security purposes.
* **`GET /login`**: Endpoint for validating the user session by checking if the provided signature matches the session signature.

#### Dependencies

* **Gin**: HTTP web framework for Go.
* **Gin CORS**: Middleware for configuring Cross-Origin Resource Sharing (CORS) in Gin.
* **Gin Sessions**: Middleware for session management in Gin.
* **Gin Cookie Store**: Session store using cookies in Gin.
* **PostgreSQL**: Database management system used for storing application data.

#### Error Handling

* **Fatal Log**: Any errors during database migration or server startup result in a fatal log, causing the application to terminate.

#### Conclusion

The `main.go` file serves as the central component of the application, responsible for initializing the server, setting up middleware, connecting to the database, defining routes, and handling errors. By leveraging the features provided by Gin and PostgreSQL, developers can create robust web applications with secure authentication and database management functionalities.\


**Attention**\
It could be more clean and robust if we have more time to develop.
