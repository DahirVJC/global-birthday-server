# Requirements

## Functional Requirements
- The system shall allow users to register, log in, and log out securely.
- The system shall allow users to CRUD birthdates, each containing day and month, name of the contact, and timezone or country.
- The system shall support a comprehensive list of timezone and countries
- The system shall calculate and display the birthday as a time interval window showing the exact start and end time in the user's current timezone.
- The birthday window shall be defined as 00:00 to 00:00 (midnight-to-midnight) of the birthday date in the contact's timezone, converted to the user's timezone.
- The system shall allow users to CRUD party events associated with a specific contact's birthday.
- Each event shall be stored in the contact's timezone and displayed with its equivalent in the user's current timezone.
- The system shall handle Daylight Saving Time transitions correctly for both contact and user timezones.
- The system shall allow users to CRUD wishlists attached to a specific contact's birthday.
- Each wishlist item shall support URL or simple text.
- The system shall allow the user to set and change their own current timezone, affecting all interval displays and event conversions.
- The system shall isolate each user's data — no user shall be able to access another user's contacts, events, or wishlists.
- The system shall allow users to reset their password via email.

## Non-Functional Requirements
- All passwords, names and wishlists shall be stored encrypted.
- All communication between client and server shall be encrypted via HTTPS/TLS.
- Session tokens shall expire after a configurable period and support revocation.
- Input validation and sanitization shall be enforced to prevent SQL injection and XSS attacks.

## Domain Requirements
- Contacts born on February 29 shall have their birthday treated as February 28 in non-leap years.
- The user's current timezone" is selectable independent of their physical location.