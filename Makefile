# Variables
PYTHON := poetry run python
PYTEST := poetry run pytest
FIND := find

# Directories
SRC_DIR := src
TEST_DIR := tests

# Targets
.PHONY: clean test run

# Clean target to delete __pycache__ directories
clean:
	$(FIND) . -type d -name "__pycache__" -exec rm -rf {} +

# Test target to run pytest with specified options
test:
	$(PYTEST) $(TEST_DIR)/test_*.py -vv -s --showlocals

# Run target to start the application
run:
	$(PYTHON) -m $(SRC_DIR).main
