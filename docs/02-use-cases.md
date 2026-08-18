# Use Cases

## Birthday Registration
Allows users to register a new contact's birthday information including name, birthdate (day/month/year), timezone, and optional notes.
````mermaid
sequenceDiagram
    Contact->>User: Provides birthday details
    User->>+App: Registers birthday information
    App->>-User: Birthday registered
    App-->>User: Birthday reminder at an appropriate time for both parties
````

##  Birthday Consulting
View all registered birthdays displayed as time intervals in the user's current timezone.
````mermaid
sequenceDiagram
    User->>+App: Queries upcoming birthdays
    App->>-User: Returns birthdates converted to user's timezone
````
## Birthday Events
Schedule a party or celebration event for a contact, with automatic timezone conversion.
````mermaid
sequenceDiagram
    User->>+App: Registers an event
    App->>-User: Event and birthday details confirmed
    App-->>User: Birthday event reminder sent
````