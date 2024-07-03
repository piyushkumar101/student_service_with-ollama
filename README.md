# Student API Service

## Overview

This project is a simple REST API written in Go that performs basic CRUD (Create, Read, Update, Delete) operations on a list of students. Each student has the following attributes:
- ID (integer)
- Name (string)
- Age (integer)
- Email (string)

Additionally, the API integrates with [Ollama](https://www.ollama.com/) to generate AI-based summaries for student profiles.

## Features

- Create a new student
- Get all students
- Get a student by ID
- Update a student by ID
- Delete a student by ID
- Generate a summary of a student by ID using Ollama

## Requirements

- Go 1.21+
- Ollama installed and running on your local machine
- Llama3 language model for Ollama

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/piyushkumar101/student_service.git
    cd student_service
    ```

2. **Install dependencies**:
    ```sh
    go mod tidy
    ```

3. **Start Ollama**:
    ```sh
    ollama serve
    ```

## Running the Application

1. **Run the Go application**:
    ```sh
    go run main.go
    ```

    The application will start and listen on port `8080`.
