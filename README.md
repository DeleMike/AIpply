# 🤖 AIpply

> AIpply for that role with confidence – let AI craft your perfect application!

[![Go Report Card](https://goreportcard.com/badge/github.com/DeleMike/AIpply?style=for-the-badge)](https://goreportcard.com/report/github.com/DeleMike/AIpply)
[![Go Version](https://img.shields.io/github/go-mod/go-version/DeleMike/AIpply?style=for-the-badge&logo=go)](https://golang.org)
[![Go Tests](https://img.shields.io/github/actions/workflow/status/DeleMike/AIpply/go.yml?branch=main&label=Tests&style=for-the-badge&logo=go)](https://github.com/DeleMike/AIpply/actions)
[![codecov](https://img.shields.io/codecov/c/github/DeleMike/AIpply?style=for-the-badge&logo=codecov)](https://codecov.io/gh/DeleMike/AIpply)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

AIpply is an open-source Go backend service designed to streamline job applications. It serves as a REST API that harnesses Google Gemini AI to create customized CVs, cover letters, and interview questions based on job descriptions and your personal responses. Whether you're a job seeker looking to save time or a developer interested in AI-driven tools, AIpply helps you stand out with professional, tailored documents.

## ✨ Features

- **Personalized Document Generation**: Automatically creates CVs and cover letters that align your experience with job requirements.
- **Interview Prep**: Generates relevant, conversational questions based on the job description and your experience level.
- **Powerful Prompts**: Uses role-based prompting strategies for high-quality, actionable outputs.
- **Metrics Tracking**: Monitors generation counts via Redis for usage insights.
- **Clean HTML Outputs**: Returns styled-ready HTML fragments for easy frontend integration.
- **Extensible API**: Built with Gin for fast, reliable REST endpoints.

## 🛠️ Tech Stack

* **Core:** Go (Golang)
* **API/Routing:** Gin-Gonic
* **AI Model:** Google Gemini (via `google.golang.org/genai`), specifically using `gemini-2.0-flash` and `gemini-2.5-flash`
* **Database (Metrics):** Redis (for tracking generation counts)
* **Configuration:** From `config-example.yaml`

## 🔍 How It Works

AIpply sends structured, role-based prompts to Google Gemini to generate personalized CVs, cover letters, and interview questions.
The model responses are returned as clean HTML fragments for the frontend to render.


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
  "jobDescription": "Looking for a junior Go developer...",
  "experienceLevel": "Just Getting Started"
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

## 📖 Usage

To interact with the API, you can use tools like curl, Postman, or integrate it into your application. Here are some curl examples assuming the server is running on `localhost:8080`:

- **Health Check**:
  ```bash
  curl http://localhost:5050/
  ```

- **Generate Interview Questions**:
  ```bash
  curl -X POST http://localhost:5050/api/v1/generate-questions \
    -H "Content-Type: application/json" \
    -d '{"job_description": "Senior Go Developer role focused on backend services.", "experience_level": "Senior (5+ years)"}'
  ```

- **Generate CV**:
  ```bash
  curl -X POST http://localhost:5050/api/v1/generate-cv \
    -H "Content-Type: application/json" \
    -d '{"job_description": "Senior Go Developer...", "answers": [{"question": "Your name?", "answer": "Jane Doe"}]}'
  ```

- **Get Metrics**:
  ```bash
  curl http://localhost:5050/api/v1/metrics
  ```

For more advanced usage, consider wrapping the API in a frontend or scripting it for batch processing.

## 🚀 Getting Started

### Prerequisites

  * Go (1.21 or later recommended)
  * A running Redis instance (check the [Docker](docker-compose.yaml) file)
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

## 🤝 Contributing

We welcome contributions from the community! Whether it's fixing bugs, adding features, improving documentation, or suggesting ideas, your help makes AIpply better for everyone.

### How to Contribute
1. **Fork the Repository**: Click the "Fork" button at the top of this page.
2. **Clone Your Fork**: `git clone https://github.com/YOUR_USERNAME/AIpply.git`
3. **Create a Branch**: Use a descriptive name, e.g., `git checkout -b feature/new-endpoint`
4. **Make Changes**: Follow the code style (use Go fmt) and add tests where applicable.
5. **Commit Your Changes**: Use clear commit messages, e.g., `git commit -m "Add support for additional AI models"`
6. **Push to Your Fork**: `git push origin feature/new-endpoint`
7. **Open a Pull Request**: Go to the original repository and submit a PR. Reference any related issues.

### Guidelines
- **Code Style**: Use Go's standard formatting (`go fmt`). Run the linter with `revive -config revive.toml ./...`.
- **Testing**: Ensure all tests pass (`go test ./...`). Aim for high coverage.
- **Issues**: Check existing issues before creating a new one. Use labels like "bug", "enhancement", or "documentation".
- **Code of Conduct**: Be respectful and inclusive. We follow the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct.html).
- **Questions?**: Open an issue or discussion for help.

For major changes, please open an issue first to discuss. Thanks for contributing!

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🛤️ Roadmap

- Integrate support for additional AI models (e.g., OpenAI).
- Expand metrics to include more analytics.
- Community-suggested features – submit yours via issues!