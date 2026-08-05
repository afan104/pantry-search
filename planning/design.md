# Design Decisions

Rationale behind choices made in this project. The README covers what the app does and how to run it; this doc covers why it's built this way.

## Auth store: SQL, not NoSQL

The `users` table (id, email, password_hash) lives in the same SQLite database as the pantry data, not a separate NoSQL store.

Emails must be unique, and account creation needs to reject a duplicate email even under concurrent signups. A relational database enforces that with a unique constraint and a transaction, for free. A NoSQL store would need that uniqueness check built by hand, with the same race risk described below unless the store offers a matching guarantee.

## Ingredient quantity updates: PUT to replace, POST to increment

Two different operations on the same resource, split across two endpoints instead of one:

- `PUT /putIngredient/:ingredient` sets the quantity to an exact value. Used after a manual inventory count ("there's actually 250g left"). Idempotent: sending the same request twice leaves the same result, so it's safe to retry.
- `POST /postIngredient/:ingredient` adds a delta to the current quantity. Used when logging a purchase ("bought 100g today"). Not idempotent: repeating it changes the result each time (adds again). Handled as one atomic read-modify-write on the server, not a client-side GET-then-PUT.

**Why not just GET the current value, add client-side, then PUT the result?** Two concurrent requests (e.g. two purchases logged close together) would both read the same starting value, compute the same new total, and the second write would silently overwrite the first, a lost update. Doing the add as a single atomic server-side operation (`UPDATE ... SET quantity = quantity + ? WHERE ...`) closes that race window, since there's no gap between the read and the write for another request to land in.

**Tradeoff accepted**: POST-to-increment isn't safe against duplicate delivery. If a client retries a timed-out request, the add can apply twice. Not worth solving with idempotency keys for a single-user personal app; worth revisiting if this ever supports concurrent multi-device use for the same account.

## Endpoint split: create vs. replace vs. increment

Three distinct intents, kept as three distinct endpoints rather than overloading one handler with conditional logic:

| Endpoint | Verb | Intent |
|---|---|---|
| `POST /postIngredient/:ingredient` | create-or-add | Log a purchase; creates the row if missing, adds to it if present |
| `PUT /putIngredient/:ingredient` | replace | Set the exact quantity after a manual count |
| `DELETE /deleteIngredient/:ingredient` | remove | Remove an ingredient entirely |

## Observability: instrument now, defer infrastructure

Structured logging (`log/slog` in place of Gin's default request logger) and a Prometheus `/metrics` endpoint are Phase 1 items, buildable and testable locally with no cloud dependency.

Distributed tracing, centralized log aggregation, and hosted alerting are deferred until the app is actually deployed somewhere with multiple instances or services. They exist to solve problems (correlating requests across services, collecting logs from many machines) that don't exist yet for a single local binary talking to one SQLite file.

Reasoning for doing the logging/metrics work early rather than after a cloud move: it's a code-level change (what the app emits), separable from the infrastructure-level change (where that output gets shipped and displayed). Building it now, in standard formats (JSON logs, Prometheus exposition format), means the later cloud/k8s move is pointing existing output at new tooling, not retrofitting instrumentation into a live app.

A Grafana dashboard is a heavier lift than the logging swap alone: standing up Grafana and Prometheus locally via Docker is quick, but a useful dashboard needs the `/metrics` endpoint instrumented first, then scrape config, then panel design. Treated as its own item, downstream of the metrics endpoint, not bundled with the `log/slog` swap.
