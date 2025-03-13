# Dockerized URL Shortener (Go) 🐳🔗

This project is a simple and efficient URL shortening service built with Go, now containerized with Docker for easy deployment and scalability. It transforms long URLs into concise, shareable links.

## ✨ Features

* **✂️ URL Shortening:** Easily convert lengthy URLs into short, unique identifiers.
* **↩️ URL Redirection:** Seamlessly redirect users from short URLs to their original destinations.
* **🔒 Concurrency Safe:** Engineered with mutexes to ensure robust thread safety under concurrent load.
* **🐳 Dockerized:** Containerized for consistent deployment across environments.
* **🚀 Easy Deployment:** Simplified setup with Docker.
* **API Simplicity:** A straightforward API for effortless URL shortening and redirection.

## 🚀 Getting Started

### 📋 Prerequisites

* Docker

### 🛠️ Installation and Running with Docker (Manual)

1.  **Clone the Repository:**

    ```bash
    git clone [https://github.com/akanshgupta98/URL-shortner.git](https://github.com/akanshgupta98/URL-shortner.git)
    cd URL-shortner
    ```

2.  **Build the Docker Image:**

    ```bash
    docker build -t url-shortener .
    ```

3.  **Run the Docker Container:**

    ```bash
    docker run -p 8080:8080 url-shortener
    ```

    The service will be accessible on `http://localhost:8080`.

### 📡 API Endpoints

#### 🔗 Shorten URL (POST)

* **Endpoint:** `/shorten`
* **Method:** `POST`
* **Request Body (JSON):**

    ```json
    {
      "url": "https://www.example.com/very/long/url"
    }
    ```

* **Response (JSON):**

    ```json
    {
      "InputURL": "http://google.com",
      "OutputURL": "http://127.0.0.1:8080/api/url-shortner/9a4759b",
      "ErrMsg": ""
    }
    ```

#### ➡️ Redirect URL (GET)

* **Endpoint:** `/api/url-shortner/9a4759b` (e.g., `/your-short-url-key`)
* **Method:** `GET`
* **Description:** Redirects users to the original long URL.

### 📝 Notes

* Ensure Docker is running on your system before executing the commands.
* If you encounter port conflicts, modify the port mapping in the `docker run` command.
* For production environments, consider using a persistent data store (e.g., Redis or PostgreSQL) instead of in-memory storage.