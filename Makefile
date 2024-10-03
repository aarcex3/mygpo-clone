# Variables
GO := go
FMT := gofmt -w
FIND := find

# Directories
SRC_DIR := ./cmd
TEST_DIR := ./test

# Targets
.PHONY: clean test format run docker init_db

# Init database (assuming there's an init.sql for Go-related DB setup)
init_db:
	rm -f mygpo-clone.sqlite
	sqlite3 mygpo-clone.sqlite < init.sql
	
# Format the Go source files
format:
	$(FMT) ./

# Clean target to delete Go build and temporary files
clean:
	$(GO) clean
	docker-compose down --volumes --rmi all

# Test target to run Go tests with specified options
test:
	rm -f testing.sqlite
	sqlite3 testing.sqlite < init.sql
	$(GO) test $(TEST_DIR) -v

# Run target to start the Go application
run:
	$(GO) run $(SRC_DIR)/api/main.go

# Build and run Docker container
docker:
	docker-compose up --build
