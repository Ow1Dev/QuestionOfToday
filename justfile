set dotenv-load

# List available commands
default:
    @just --list

# Run solution for a giving day 
run ARGS="":
    @go run cmd/web/main.go {{ARGS}}

compile-style:
    @tailwindcss -i ./public/style/input.css -o ./public/style/output.css

update:
   nix flake update 

# Run test all solutions
test-all:
    @go test ./...

# format the code
fmt:
    @go fmt ./...

up:
  docker compose up -d

down:
  docker compose down --remove-orphans --volumes

migrate-up:
  migrate -database $DATABASE_URL -path ./migrations up
