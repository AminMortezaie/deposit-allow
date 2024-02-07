# Delivery

Here's the detailed documentation for the delivery layer's wallet functionality:

### Wallet Delivery&#x20;

The `delivery/wallet` folder contains functionality related to wallet management and authentication. It includes endpoints for verifying and validating user sessions.

#### Endpoints

**Verify**

* **Method**: `POST`
* **Path**: `/verify`
* **Description**: Verifies the authenticity of a user by validating their message signature against the provided message and signature in the request body. If the verification is successful, the user is logged in, and a session is created.
*   **Request Body**:

    ```json
    {
      "Message": "string",
      "Signature": "string"
    }
    ```
* **Response**:
  * `200 OK`: Successful authentication.
  * `401 Unauthorized`: Authentication failed.

**ValidateSession**

* **Method**: `GET`
* **Path**: `/validate-session`
* **Description**: Validates the user session by checking if the provided signature matches the session signature. If the session is valid, the user is considered authenticated.
* **Response**:
  * `200 OK`: Session authenticated.
  * `401 Unauthorized`: Session not authenticated.

#### Dependencies

* **Gin Sessions**: Used for session management.
* **SIWE-Go**: Library for parsing and verifying SIWA (Sign in with Apple) messages.

#### Middleware

No middleware is explicitly defined in this file, but the use of Gin sessions implies the use of session middleware.

#### Initialization

* **Database Connection**: A connection to the PostgreSQL database is established using the `DBConnector` from the `backend/repository/db/postgres` package.

#### Service Interaction

* **Service Initialization**: The `wallet.Service` is initialized with a repository instance to interact with the database.
* **Function Invocation**: The `Verify` and `ValidateSession` functions interact with the wallet service to perform authentication and session validation.

#### Error Handling

* **Panic**: Errors during database connection or binding of request body result in panics.
* **Logging**: Errors during session saving are logged.

#### Logging

* **Log Fatal**: Errors during session saving are logged as fatal errors.

#### Conclusion

The wallet delivery functionality provides endpoints for verifying and validating user sessions, ensuring secure and authenticated access to wallet-related resources within the application. By leveraging SIWA messages and Gin sessions, developers can implement robust authentication mechanisms for user interactions with wallet functionalities.
