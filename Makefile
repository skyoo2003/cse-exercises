.PHONY: all setup clean test lint

all: test lint

setup:
	pre-commit install
	set -e && \
		rm -rf .venv && \
		python -m venv .venv && \
		source .venv/bin/activate && \
		pip install -r requirements.txt

clean:
	find . -name '*.pyc' -exec rm -f {} +
	find . -name '*.pyo' -exec rm -f {} +
	find . -name '*~' -exec rm -f {} +
	find . -name '__pycache__' -exec rm -fr {} +

test:
	go test -v ./...
	set -e && \
		source .venv/bin/activate && \
		pytest -v -s

lint:
	golangci-lint run ./...
	set -e && \
		source .venv/bin/activate && \
		flake8 .
