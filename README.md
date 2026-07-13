
# Fiber + Pgx Example

A small Fiber example that shows a layered handler/business/model structure with pgx.

- http://gofiber.io
- postgresql + pgadmin4
- redis
- pgx
- docker
- docker-compose
- goose, go-migrate

## Folder Structure

- model - handle all persistence, and 3rd party calls
- business (controller) -  handle all business logic
- presentation (view) - handle all parsing and rendering

## Maintenance checklist

- [x] Go runtime updated to 1.26.5.
- [x] Direct Go dependencies refreshed and module files tidied.
- [x] Added a sanity test for article validation without requiring PostgreSQL.
- [x] `make test` compiles and tests all Go packages.
- [x] `make verify-dependency-security` and `make vulncheck` check dependency security.
