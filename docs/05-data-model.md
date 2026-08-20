# Data Model

## Entities
### USER
Represents an account holder who manages birthday, reminders, and events.

### BIRTHDAY
Represents a person whose birthday is being tracked.

### EVENT
Represents scheduled activities related to a birthday.

````mermaid
erDiagram
    USER ||--o{ BIRTHDAY : registers
    USER ||--o{ SESSION : generates
    BIRTHDAY ||--o{ EVENT : has

    USER {
        UUID _id
        string name
        string password
        string email
        string timezone
        BIRTHDAY[] birthdays
        SESSION[] sessions
    }
    
    SESSION {
        UUID _id
        string token
        date createdAt
        date expiresAt
    }

    BIRTHDAY {
        UUID _id
        string name
        date birthdate
        string timezone
        date remindMeIn
        string[] contactLinks
        string[] wishlists
        EVENT[] events
    }

    EVENT {
        UUID _id
        date eventDateStart
        date eventDateEnd
        date remindMeIn
        string description
    }
````

## Considerations
### User's fields
- password and email are encrypted.
- 
### Session's fields
- token is encrypted.

### Birthday's fields
- name, contactLinks, and wishlists are encrypted.
- remindMeIn, contactLinks, wishlists and events are optional.
- birthdate's year is always 2000, a leap year.

### Event's fields
- description is encrypted.
- eventDateEnd, remindMeIn and description are optional.
