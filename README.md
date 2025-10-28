# 🤖 AIpply

> AIpply for that role with confidence!


[![Go Report Card](https://goreportcard.com/badge/github.com/DeleMike/AIpply?style=for-the-badge)](https://goreportcard.com/report/github.com/DeleMike/AIpply)
[![Go Version](https://img.shields.io/github/go-mod/go-version/DeleMike/AIpply?style=for-the-badge&logo=go)](https://golang.org)
[![Go Tests](https://img.shields.io/github/actions/workflow/status/DeleMike/AIpply/go.yml?branch=main&label=Tests&style=for-the-badge&logo=go)](https://github.com/DeleMike/AIpply/actions)
[![codecov](https://img.shields.io/codecov/c/github/DeleMike/AIpply?style=for-the-badge&logo=codecov)](https://codecov.io/gh/DeleMike/AIpply)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

This repository contains the Go backend service for **AIpply**. It functions as a REST API that leverages the Google Gemini AI to generate tailored CVs, cover letters, and interview questions based on a job description and user-provided answers.

## 🛠️ Tech Stack

* **Core:** Go (Golang)
* **API/Routing:** Gin-Gonic
* **AI Model:** Google Gemini (via `google.golang.org/genai`), specifically using `gemini-2.0-flash` and `gemini-2.5-flash`
* **Database (Metrics):** Redis (for tracking generation counts)
* **Configuration:** From `config-example.yaml` and environment variables
* **Linting:** Revive

## ✨ How It Works: The Prompting Strategy

The core intelligence of this service lies in its detailed, role-based prompts sent to the Gemini model.

* **For CVs (`CVPrompt`)**: The prompt instructs the AI to act as an "elite career coach." It takes a JSON array of the user's Q&A and a job description, then synthesizes the narrative answers into "concise, powerful, action-oriented bullet points" associated with the correct job.
* **For Cover Letters (`CoverLetterPrompt`)**: The prompt sets the AI as an "expert career coach" that analyzes the job description for key skills and the user's answers for strong "stories." It then writes a letter that *explicitly connects* those stories to the job's requirements.
* **For Questions (`QuestionPrompt`)**: This prompt casts the AI as an "experienced hiring manager," generating 5-7 simple, conversational questions based on the job description and the candidate's stated experience level.

A critical requirement for both the CV and cover letter is that the AI **must return a clean HTML fragment** with no `<html>`, `<body>`, or `<style>` tags, allowing the frontend to handle all styling.

## 🔌 API Endpoints

All routes are versioned under the `/api/v1` prefix.

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | Health check route. |
| `POST` | `/api/v1/generate-questions` | Generates interview questions from a job description and experience level. |
| `POST` | `/api/v1/generate-cv` | Generates a full CV in HTML format based on a job description and a JSON array of Q&A pairs. |
| `POST` | `/api/v1/generate-cover-letter` | Generates a cover letter in HTML format based on a job description and a JSON array of Q&A pairs. |
| `GET` | `/api/v1/metrics` | Retrieves generation metrics (CV/Cover Letter counts) from Redis. |

### Request Body Examples

**`POST /generate-questions`**
```json
{
  "jobDescription": "Looking for a senior Go developer...",
  "experienceLevel": "Senior (5+ years)"
}
```

**`POST /generate-cv`** or **`POST /generate-cover-letter`**

```json
{
  "jobDescription": "Looking for a senior Go developer...",
  "answers": [
    {
      "question": "First, what is your full name, email address, and phone number?",
      "answer": "Jane Doe, jane.doe@email.com, 123-456-7890"
    },
    {
      "question": "Briefly describe your relevant experience.",
      "answer": "Senior Dev at TechCorp, 2020-2023. I led the backend team..."
    }
  ]
}
```

## ⚙️ Configuration

The application is configured using a `config.yaml` file and/or environment variables, managed by Viper.

| Environment Variable | `config.yaml` Key | Default | Description |
| :--- | :--- | :--- | :--- |
| `API_KEY` | `api_key` | - | **Required.** Your Google Gemini API key. |
| `REDIS_ADDR` | `redis.addr` | `localhost:6379` | Address for the Redis server. |
| `REDIS_PASSWORD` | `redis.password` | `""` | Password for the Redis server. |
| `REDIS_DB` | `redis.db` | `0` | Redis database index. |
| `SERVER_PORT` | `server.port` | - | Port for the Gin server (e.g., `8080`). |
| `GIN_MODE` | `server.gin_mode` | `debug` | Gin framework mode (`debug`, `release`, or `test`). |

## 🚀 Getting Started

### Prerequisites

  * Go (1.21 or later recommended)
  * A running Redis instance
  * A Google Gemini API Key

### Installation & Running

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/DeleMike/AIpply.git
    cd AIpply
    ```

2.  **Install dependencies:**

    ```sh
    go mod tidy
    ```

3.  **Configure the application:**

    Create a `config.yaml` file in the root directory, or set the environment variables listed in the table above.

    ```bash
    # Copy the example configuration file
    cp config-example.yaml config.yaml
    ```

    Then open `config.yaml` and fill in your own settings:

    ```yaml
    # config.yaml
    api_key: "your-gemini-api-key"
    
    server:
      port: 8080
      gin_mode: "debug"  # Use "release" in production
    
    redis:
      addr: "localhost:6379"
      password: ""
      db: 0
    ```

    > **💡 Tip:** Keep your API keys private and never commit them to version control. Add `config.yaml` to your `.gitignore` file. Environment variables can override any config value and take precedence over `config.yaml` settings.

4.  **Run the server:**

    ```sh
    go run .
    ```

    The server will start on the port you configured.

## 🧪 Testing & Linting

  * **Run all tests:**

    ```sh
    go test ./...
    ```

  * **Run the linter (Revive):**

    ```sh
    # Install revive if you haven't
    go install github.com/mgechev/revive@latest

    # Run from the root directory
    revive -config revive.toml ./...
    ```