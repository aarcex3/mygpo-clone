test:
	poetry run pytest tests/test_*.py -vv -s --showlocals

run:
	poetry run python -m src.main