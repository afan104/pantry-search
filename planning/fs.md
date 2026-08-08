File structure:

```
myapp/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go          # app entrypoint
│   │
│   ├── internal/
│   │   └── handlers/            # HTTP handlers (controllers)
│   │       ├── deleteIngredients.go
│   │       ├── getIngredient.go
│   │       ├── getIngredients.go
│   │       ├── postIngredient.go
│   │       └── putIngredient.go
├── frontend/
│   ├── src/
│   ├── public/
│   └── package.json
│
├── db/
│   ├── schema.sql               # initial schema
│   └── seeds.sql                # for testing with fake users
│
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
