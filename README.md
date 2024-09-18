# Simple-App Project

This project is a simple web application built using `go-app v10`. It demonstrates how to implement a Progressive Web App (PWA) in Go, including features like routing, API calls, and component-based UI.

## Project Structure

```bash
myapp/
├── main.go                # Main entry point of the application
├── components/            # Folder containing all UI components
│   ├── navbar.go          # Navigation bar component
│   ├── welcome.go         # Welcome page component
│   └── filemanager.go     # API Calls view component
├── web/
│   └── styles.css         # CSS styles for the application
├── Makefile               # Makefile for building and running the app
└── README.md              # Project documentation (this file)
