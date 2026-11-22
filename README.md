# Go + GORM PostgreSQL — User Management (2025)

Educational Go console application demonstrating CRUD operations using **GORM** and **PostgreSQL**.

> **Project status:** original source recovered from this repository. Requires a PostgreSQL database and local configuration. The application has **not yet been verified by an automated CI build or live database test**. It is a console program, not a web frontend or HTTP API.

## What the program does

When you run `main.go`, it:
1. Loads database settings from a local `.env` file.
2. Connects to PostgreSQL and applies the `User` table migration.
3. Creates four sample users.
4. Queries users older than 18.
5. Changes the name and age of the first matching user.
6. Deletes the second matching user.
7. Prints the remaining records to the terminal.

> **Important:** Running the program modifies the configured database (inserts, updates, deletes). Use a disposable development database, **not production**. Sample names are hard-coded; a second run may encounter uniqueness conflicts.

## Requirements

- Go **1.22** or newer
- PostgreSQL (local or in Docker)
- Internet connection once to download Go modules

## Run locally

1. Clone this repository:

   ```bash
   git clone https://github.com/Zhenis455/GolandProject1.git
   cd GolandProject1
   ```

2. Create a dedicated PostgreSQL database, e.g. `go_users_dev`.

3. Copy the example environment file into `.env`:

   **Windows CMD:** `copy .env.example .env`

   **macOS/Linux:** `cp .env.example .env`

   Edit `.env` and provide your **own local** PostgreSQL username and password.

4. Download modules and start:

   ```bash
   go mod download
   go run .
   ```

The program prints created users, updated records and database status in the terminal. There is no web URL to open.

## Code layout

- `main.go` — demonstration workflow
- `controller/userController.go` — GORM CRUD operations
- `initializer/config.go` — environment file loading
- `initializer/db.go` — PostgreSQL connection
- `migrate/main.go` — automatic table migration
- `model/user.go` — user model

## Quality and security notes

- No automated unit tests have been found in this repository yet.
- `.env` must **never** be committed; use `.env.example`.
- If database credentials were pushed in an earlier commit, removing the file does not erase Git history: **rotate affected credentials**.
- This example does not implement authentication, production migrations or HTTP endpoints.

## История проекта

Учебный пример на Go + GORM + PostgreSQL. Реализует CRUD для пользователей через консоль. Для запуска нужен PostgreSQL и собственный файл `.env`. Проверка реального запуска ещё не завершена.
