# URL Shortener (Go) 🔗

A simple and efficient URL shortening service built with Go, designed to transform long URLs into concise, shareable links.

## ✨ Features

* **✂️ URL Shortening:** Easily convert lengthy URLs into short, unique identifiers.
* **↩️ URL Redirection:** Seamlessly redirect users from short URLs to their original destinations.
* **🔒 Concurrency Safe:** Engineered with mutexes to ensure robust thread safety under concurrent load.
* **API Simplicity:** A straightforward API for effortless URL shortening and redirection.

## 🚀 Getting Started

### 📋 Prerequisites

* Go 1.18 or higher
* (Optional) `golangci-lint` for enhanced code quality

### 🛠️ Installation

1.  **Clone the Repository:**

    ```bash
    git clone https://github.com/akanshgupta98/URL-shortner.git
    cd URL-shortner
    ```

2.  **Build the Application:**

    ```bash
    go build -o url-shortner cmd/main.go
    ```

### 🏃 Running the Application

```bash
./url-shortener
```

The service will launch on http://localhost:8080 by default.

📡 API Endpoints
🔗 Shorten URL (POST)
Endpoint: /api/url-shortner

Method: POST

Request Body (JSON):

```
{
  "url": "https://www.example.com/very/long/url"
}
```
Response (JSON):

```
{
    "InputURL": "https://www.example.com/very/long/url",
    "OutputURL": "http://127.0.0.1:8080/api/url-shortner/9a4759b",
    "ErrMsg": ""
}
```
➡️ Redirect URL (GET)

Endpoint: api/url-shortner/9a4759b (e.g., /your-short-url-key)

Method: GET

Description: Redirects users to the original long URL.
