# Groupie Tracker

## Overview
Groupie Tracker is a web application that displays information about various musical artists and their concert details. The project fetches data from a provided API and presents it in a user-friendly manner using HTML, CSS, and JavaScript for the frontend, with a backend powered by Go.

## Features
- **Artist Information**: Displays artist details such as name, image, members, first album, and activity start year.
- **Concert Locations & Dates**: Lists past and upcoming concerts with corresponding dates and locations.
- **Search & Filter**: Allows users to search for artists and filter concerts.
- **Interactive Map**: Displays concert locations on an interactive map.
- **Client-Server Communication**: Implements an event-driven feature where user interactions trigger a request to the server for additional data.

## Technologies Used
### Backend:
- **Go**: Used for handling HTTP requests, API integration, and serving web pages.
- **net/http**: Used to create the web server.
- **encoding/json**: Used for parsing API responses.
- **html/template**: Used for rendering HTML templates dynamically.

### Frontend:
- **HTML & CSS**: Used for structuring and styling the web pages.

## Installation & Setup
### Prerequisites:
- Go (1.22.2 version recommended)
- A web browser (Chrome, Firefox, etc.)

### Steps:
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/groupie-tracker.git
   cd groupie-tracker
   ```
2. Run the server:
   ```sh
   go run main.go
   ```
3. Open a browser and visit:
   ```
   http://localhost:8080
   ```

## API Structure
The application fetches data from an external API that consists of four main endpoints:
- `/artists`: Provides artist details (name, image, members, first album, etc.).
- `/locations`: Lists concert locations.
- `/dates`: Contains concert dates.
- `/relation`: Links the above data to establish artist-concert relations.

## Event-Driven Feature
- Clicking on an artist dynamically loads their upcoming concert details without reloading the page.
- The server processes the request and returns concert data, which is displayed on the UI.

## Error Handling
- Handles API failures gracefully by displaying error messages to users.
- Uses logging to capture server-side errors.
- Prevents crashes due to invalid API responses.

<!-- ## Testing
- Unit tests are implemented using Go's `testing` package.
- Tests cover API fetching, data parsing, and server responses.
- Run tests with:
  ```sh
  go test ./...
  ``` -->

## Future Improvements
- Add user authentication for personalized features.
- Implement caching to reduce API calls and improve performance.
- Enhance UI with animations and better UX.

