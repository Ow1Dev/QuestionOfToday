# List available commands
default:
    @just --list

# Run solution for a giving day 
run:
    @go run cmd/main.go

update:
   nix flake update 

# Run test all solutions
test-all:
    @go test ./...


# format the code
fmt:
    @go fmt ./...
