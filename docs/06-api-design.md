# API Overview

## User Endpoints

- `GET` `user/me/`: 
  - Receives a valid access token
  - Returns the user's name and configuration
  
- `POST` `user/`: 
  - Receives user data
  - Creates a new user
  - Returns 201 if successful
  
- `PUT` `user/me/`: 
  - Receives a valid access token
  - Receives new user's information (except password)
  - Changes the user's information
  - Returns 201 if successful
  
- `PATCH` `user/me/password`:
  - Receives a valid access token
  - Receives old and new passwords
  - Replaces the user's old password and encryption key with new values
  - Encrypted information is decrypted and re-encrypted with the new encryption key
  - Returns 200 if successful
  
- `Delete` `user/me/`
  - Receives a valid access token
  - User is deleted
  - Returns 200 if successful

## Birthday Endpoints

- `GET` `birthday/{Year}`:
  - Receives a valid access token
  - Receives the consulted year as an integer
  - Converts the birthdays for the given year to a time interval represented by a start and end date in the users' timezone
  - If a birthdate is February 29 and the given year is not a leap year, it converts it into February 28
  - Returns all processed birthdays for the user

- `GET` `birthday/{Year}/{id}`:
  - Receives a valid access token
  - Receives the consulted year as an integer
  - Receives the birthday's ID
  - Converts the birthdate for the given year to a time interval represented by a start and end date in the users' timezone
  - If the birthdate is February 29 and the given year is not a leap year, it converts it into February 28
  - Returns the processed birthday for the user
  - Returns 404 if the birthday doesn't exist
  - Returns 404 if the birthday belongs to another user

- `POST` `birthday/`
  - Receives a valid access token
  - Receives birthday data, including wishlists
  - Creates the birthday for the user
  - Returns 201 if successful

- `PUT` `birthday/{id}`
  - Receives a valid access token
  - Receives the birthday's ID
  - Receives new birthday data, including wishlists
  - Updates the selected birthday
  - Returns 200 if successful
  
- `DELETE` `birthday/{id}`
  - Receives a valid access token
  - Receives the birthday's ID
  - Deletes the specified birthday
  - Returns 200 if successful

## Event Birthday

- `GET` `event/`:
  - Receives a valid access token
  - Receives a boolean as a query parameter to indicate whether it should list past birthdays
  - Converts registered events to a time interval represented by a start and end date in the users' timezone
  - Returns all processed events for the user

- `GET` `event/{id}`:
  - Receives a valid access token
  - Receives the event's ID
  - Converts the event for the given year to a time interval represented by a start and end date in the users' timezone
  - Returns the processed event for the user
  - Returns 404 if the event doesn't exist
  - Returns 404 if the event belongs to another user

- `POST` `event/`
  - Receives a valid access token
  - Receives event data
  - Creates the event for the user
  - Returns 201 if successful

- `PUT` `event/{id}`
  - Receives a valid access token
  - Receives the event's ID
  - Receives new event data
  - Updates the selected event
  - Returns 200 if successful

- `DELETE` `birthday/{id}`
  - Receives a valid access token
  - Receives the event's ID
  - Deletes the specified event
  - Returns 200 if successful

## Authentication Endpoints

- `POST` `/auth/login`
  - Receives credentials
  - Authenticates user and creates a new session
  - Returns access token, refresh token, and expiration details
  
- `POST` `/auth/logout`
  - Receives a valid access token
  - Logs out the user and delete the current session
  - Returns 200 if successful
  
- `POST` `/auth/refresh`
  - Receives a valid refresh token
  - Creates a new access token for the current session
  - Returns 200 if successful

- `POST` `/auth/verify-email`
  - Receives an email verification token
  - Verifies the user's email address and invalidates the verification token
  - Returns 200 if successful

- `POST` `/auth/resend-verification`
  - Receives the user's email address
  - Creates a new email verification token and sends a verification email
  - Returns 200 if successful

- `POST` `/auth/forgot-password`
  - Receives the user's email address
  - Creates a password reset token and sends a password reset email
  - Returns 200 if successful

- `POST` `/auth/reset-password`
  - Receives a valid password reset token and a new password
  - Updates the user's password, email and encryption key, and invalidates the password reset token
    - User's password turns into the new password and the user's email is replaced by the one stored in the auth token encrypted by the new encryption key
  - Warning: deletes all encrypted optional data and replaces mandatory fields with a default value according to its type:
    - string: "-"
    - numeric: "0"
    - date: "2000/01/01"
    - bool: false
  - Returns 200 if successful