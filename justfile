# List available commands
default:
    @just --list

# Run solution for a giving day 
run:
    @go run cmd/web/main.go

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
