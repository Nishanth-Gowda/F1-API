# Formula 1 Teams API

![Formula 1 Logo](assets/f1_logo.png)
> A RESTful API for managing Formula 1 teams.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)

## Introduction

The Formula 1 Teams API is a Go and Gin-based application that provides CRUD (Create, Read, Update, Delete) operations for managing Formula 1 teams. Whether you're a Formula 1 enthusiast or a developer, this API allows you to interact with data related to Formula 1 teams, including team information, drivers, and championships won.

## Features

- Create, Read, Update, and Delete Formula 1 teams.
- Get a list of all Formula 1 teams.
- Retrieve details of a specific Formula 1 team.
- Add or update drivers for a team.
- Track championships won by teams.
- Secure API endpoints with authentication.

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (at least Go 1.13)
- Gin framework
- MySQL or your preferred database

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/f1-teams-api.git
   ```

2. Install dependencies:

    ``` sh
    go get ./...
    ```


3. Build and run the application:

    ``` sh
    go run main.go
    ```

## Usage
    **Endpoints**

    Here are some of the available endpoints:

        ```
        GET /teams: Get a list of all Formula 1 teams.
        GET /teams/{id}: Retrieve details of a specific Formula 1 team.
        POST /teams: Create a new Formula 1 team.
        PUT /teams/{id}: Update details of a specific Formula 1 team.
        DELETE /teams/{id}: Delete a Formula 1 team.
        PUT /teams/{id}/drivers: Update the list of drivers for a team.
    ```



