# Testing

## Routes
The routes/test folder contains test files for each route in routes. All of those files have functions to test each endpoint.
- `status_test`
    - `TestStatusRoute`: Checks if /api/status/ returns "Server is up."
- `users_test`
    - `TestGetUsersRoute`: Makes a GET request to /api/users/ and verifies that the returned data matches the mock data
    - `TestGetUserRoute`: Performs a GET request to /api/users/me with a valid access token and verifies the returned user
    - `TestGetUserNotFoundRoute`: Performs a GET request to /api/users/me with an invalid access token and verifies that it returns a 404 HTTP status
    - `TestPostUserRoute`: Creates a new user via POST /api/users/ and retrieves all users with GET /api/users/ to check if the general fields of the saved user match the inserted user
    - `TestPutUserRoute`: Modifies a user and retrieves it to validate the data changes
    - `TestDeleteUserRoute`: Deletes a user and attempts to fetch it; if it returns a 404 Not Found status, the test is successful
