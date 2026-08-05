# Pantry Search

A personal pantry inventory tracker. Search what's in your pantry, log purchases, track expiry dates, and match recipes against what you already have on hand.

## Status

Early development. The backend API is in progress (Go + Gin + SQLite). Frontend and recipe-matching are not started yet.

## Tech Stack

| Layer    | Choice                                     |
| -------- | ------------------------------------------ |
| Backend  | Go, [Gin](https://github.com/gin-gonic/gin) |
| Database | SQLite (`mattn/go-sqlite3`)                |
| Frontend | Not yet started                            |

## Getting Started

Requires Go 1.25+.

```bash
git clone https://github.com/afan104/pantry-search.git
cd pantry-search
go mod tidy
go run ./backend/cmd/api/main.go
```

The API serves on `:3000`. A SQLite database file is created at the repo root on first run.

## API Endpoints

| Method | Path                            | Description                                                                          |
| ------ | ------------------------------- | ------------------------------------------------------------------------------------ |
| GET    | `/getIngredients`               | List the entire pantry                                                               |
| GET    | `/getIngredient/:ingredient`    | Look up a specific ingredient                                                        |
| POST   | `/postIngredient/:ingredient`   | Add quantity to an existing entry, or create it if it doesn't exist                  |
| PUT    | `/putIngredient/:ingredient`    | Set an ingredient's quantity to an exact value (e.g. after a manual inventory count) |
| DELETE | `/deleteIngredient/:ingredient` | Remove an ingredient from the pantry                                                 |

## Database Schema

**Pantry**

| Field          | Notes    |
| -------------- | -------- |
| ingredient     |          |
| ingredientType | optional |
| quantity       |          |
| units          |          |
| dateUpdated    |          |
| expectedExpiry |          |

**Users**

| Field         | Notes                               |
| ------------- | ----------------------------------- |
| id            |                                     |
| email         | unique, enforced by SQL constraints |
| password_hash | salted and hashed before storage    |

Relational (SQLite) vs. NoSQL: unique-email enforcement and race-safe writes need ACID guarantees.

## Roadmap

### Phase 1: Basic Functionality

- [ ] View entire pantry
- [ ] Search pantry for a specific ingredient
- [ ] Add to an existing entry or create a new one
- [ ] Correct an entry to an exact value after a manual inventory check
- [ ] Remove an ingredient
- [ ] Flag expiring/expired ingredients: expired items move to a separate section at the bottom in red; items expiring soon are flagged in a box alongside them
- [ ] Password salting/hashing for user auth
- [ ] Unit tests
- [ ] CI/CD
- [ ] Swap to log/slog using Gin's built-in logging for requests

### Phase 2: Advanced Functionality

- [ ] Submit a full recipe and deduct its ingredients from the pantry, with a review step before committing and an "undo" to put ingredients back
- [ ] Pull ingredients directly from a recipe URL (scraping)
- [ ] Standardized units with a conversion system; unconvertible inputs are shown in original units for manual comparison
- [ ] Automatic backups of pantry state before each action, so any change can be undone
- [ ] Rate-limiting for authentication security
- [ ] Prometheus `/metrics` endpoint
- [ ] Grafana dashboard for visualization, reading from the `/metrics` endpoint
- [ ] Move to cloud

### Future Considerations

- Multi-user support with per-user auth
- Bulk input via file upload
- Mobile compatibility
- API versioning, once there's a client (e.g. a mobile app) that can't update in lockstep with the backend
- Waste tracking (sustainability)
