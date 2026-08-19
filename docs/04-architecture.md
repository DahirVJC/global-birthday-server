# Architecture
Users interact with the mobile app, which communicates with the Node.js REST API via HTTPS/JSON requests. The API then performs read/write operations on the MongoDB database.

````mermaid
%%{init: {
  "theme": "dark"
}}%%
C4Context
    title Global Birthday System
    Enterprise_Boundary(b0, "C1") {
        Person(user, "User", "Individual managing birthdays of contacts across different timezones")

        System_Boundary("Global Birthday System", "Allows users to manage and view  birthdays and events adjusted to the current timezone"){
            Container(mobile, "Mobile Application", "React Native", "Application for registering birthdays and events, and display them with a calendar-like view")

            System_Boundary(backend, "Backend") {
                Container(api, "API", "Node.js / REST API", "Handles CRUD operations for birthdays, events, and wishlists")
                ContainerDb(mongodb, "MongoDB", "MongoDB", "Stores user accounts, birthday records, party events, and wishlist items")
            }
            Rel(user, mobile, "Uses")
            Rel(mobile, api, "Calls API", "HTTPS/JSON")
            Rel(api, mongodb, "Reads from / writes to", "MongoDB")
        }
    }

    UpdateRelStyle(user, mobile, $textColor="#fff", $lineColor="#fff")
    UpdateRelStyle(mobile, api, $textColor="#fff", $lineColor="#fff")
    UpdateRelStyle(api, mongodb, $textColor="#fff", $lineColor="#fff")
````