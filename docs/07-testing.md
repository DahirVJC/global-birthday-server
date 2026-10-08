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
- `birthdays_test`
    - `TestGetBirthdaysRoute`: Makes a GET request to /api/birthday/{Year} with a valid access token and verifies converted intervals in the user's timezone
    - `TestGetBirthdayRoute`: Performs a GET request to /api/birthday/{Year}/{id} and verifies the processed birthday
    - `TestGetBirthdayLeapYearFallbackRoute`: Performs a GET request for a February 29 birthday in a non-leap year and verifies conversion to February 28
    - `TestGetBirthdayNotFoundRoute`: Performs a GET request for a missing birthday ID and verifies that it returns a 404 HTTP status
    - `TestGetBirthdayForbiddenAsNotFoundRoute`: Performs a GET request for another user's birthday and verifies that it returns a 404 HTTP status
    - `TestPostBirthdayRoute`: Creates a birthday via POST /api/birthday/ and retrieves it with GET /api/birthday/{Year} to check persisted fields and year normalization
    - `TestPutBirthdayRoute`: Modifies a birthday and retrieves it to validate the data changes
    - `TestDeleteBirthdayRoute`: Deletes a birthday and attempts to fetch it; if it returns a 404 Not Found status, the test is successful
