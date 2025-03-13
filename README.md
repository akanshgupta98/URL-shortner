# Dockerized URL Shortener (Go) with Redis 🐳🔗

This project is a simple and efficient URL shortening service built with Go, now containerized with Docker Compose for easy deployment and scalability. It transforms long URLs into concise, shareable links, using Redis for persistent storage.

## ✨ Features

* **✂️ URL Shortening:** Easily convert lengthy URLs into short, unique identifiers.
* **↩️ URL Redirection:** Seamlessly redirect users from short URLs to their original destinations.
* **🔒 Concurrency Safe:** Engineered with mutexes to ensure robust thread safety under concurrent load.
* **🐳 Dockerized:** Containerized for consistent deployment across environments.
* **🚀 Easy Deployment:** Simplified setup with Docker Compose.
* **💾 Redis Persistence:** Uses Redis for persistent storage of URL mappings.
* **API Simplicity:** A straightforward API for effortless URL shortening and redirection.

## 🚀 Getting Started

### 📋 Prerequisites

* Docker
* Docker Compose

### 🛠️ Installation and Running with Docker Compose

1.  **Clone the Repository:**

    ```bash
    git clone https://github.com/akanshgupta98/URL-shortner.git
    cd URL-shortner
    ```

2.  **Run with Docker Compose:**

    ```bash
    docker-compose up --build -d
    ```

    * This command will build the Docker image, start the Redis and application containers in detached mode (background).
    * The service will be accessible on `http://localhost:8080`.

3. **Stop the containers:**

    ```bash
    docker-compose down
    ```

### 📡 API Endpoints

#### 🔗 Shorten URL (POST)

* **Endpoint:** `/shorten`
* **Method:** `POST`
* **Request Body (JSON):**

    ```json
    {
      "url": "https://hub.docker.com/layers/library/alpine/latest/images/sha256-2436f2b3b7d2537f4c5b622d7a820f00aaea1b6bd14c898142472947d5f02abe"
    }
    ```

* **Response (JSON):**

    ```json
    {
      "InputURL": "https://hub.docker.com/layers/library/alpine/latest/images/sha256-2436f2b3b7d2537f4c5b622d7a820f00aaea1b6bd14c898142472947d5f02abe",
      "OutputURL": "http://127.0.0.1:8080/api/url-shortner/9a4759b",
      "ErrMsg": ""
    }
    ```

#### ➡️ Redirect URL (GET)

* **Endpoint:** `/api/url-shortner/9a4759b` (e.g., `/your-short-url-key`)
* **Method:** `GET`
* **Description:** Redirects users to the original long URL.

### 📝 Notes

* Ensure Docker and Docker Compose are running on your system before executing the commands.
* If you encounter port conflicts, modify the port mapping in the `docker-compose.yml` file.
* Data is persisted using a Docker volume for the Redis container.
* To check container logs use `docker-compose logs <container name>`. For example `docker-compose logs app`.
* The application uses the redis service name as the hostname.
* The redis service is accessible on port 6379.
* The redis address is set with the environment variable REDIS_ADDR inside of the app container.