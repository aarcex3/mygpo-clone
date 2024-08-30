# Variables
PYTHON := poetry run python
PYTEST := poetry run pytest
BLACK := poetry run black
FIND := find

# Directories
SRC_DIR := src
TEST_DIR := tests

# Targets
.PHONY: clean test run


# Init database
init_db:
	rm -f mygpo-clone.sqlite
	sqlite3 mygpo-clone.sqlite < init.sql
	
#Format the files
format:
	$(BLACK) $(SRC_DIR)

# Clean target to delete __pycache__ directories
clean:
	$(FIND) . -type d -name "__pycache__" -exec rm -rf {} +
	docker-compose down --volumes --rmi all

# Test target to run pytest with specified options
test:
	rm -f testing.sqlite
	sqlite3 testing.sqlite < init.sql
	$(PYTEST) $(TEST_DIR)/test_*.py -vv -s --showlocals

# Run target to start the application
run:
	$(PYTHON) -m $(SRC_DIR).main


docker:
	docker-compose up --build

