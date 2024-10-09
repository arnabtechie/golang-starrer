# Go/Gin PostgreSQL Starter Application

This is a starter application built with Go, the Gin web framework, and PostgreSQL. This project serves as a template for building RESTful APIs with a focus on best practices.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Installation & Usage](#installation)
- [License](#license)

## Prerequisites

Make sure you have the following installed on your machine:

- Go Version >= 1.20
- PostgreSQL
- Git
- (Optional) Docker

Change the .env file on your requirements


## Installation & Usage

1. Clone this repository:
   ```bash
   git clone https://github.com/arnabtechie/golang-starrer.git
   cd golang-starrer
   go mod tidy
   go run main.go
2. For building
    ```bash
    go build
3. Use air for auto detection of file changes
    ```bash
    go install github.com/cosmtrek/air@latest
    air