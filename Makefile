# Directory where the main program resides
PROGRAM_DIR := cmd/api

# Run the Go application located in PROGRAM_DIR
run:
	go run ./$(PROGRAM_DIR)

# Command to manage the air tool
AIR_CMD := air
air:
	# Check if the 'air' command is available
	@if ! command -v $(AIR_CMD) > /dev/null 2>&1; then \
		echo "air not found, installing..."; \
		go install github.com/air-verse/air@latest; \
	fi
	# Run the 'air' tool
	@echo "Running air..."; \
	$(AIR_CMD)

# Configuration for PostgreSQL container
POSTGRES_CONTAINER_NAME := greenlight
POSTGRES_USER := admin
POSTGRES_PASSWORD := It123456@
POSTGRES_DB := greenlight
POSTGRES_PORT := 5432

# Create and start a PostgreSQL container
start_postgres:
	@echo "Creating and starting PostgreSQL container..."
	@if docker ps -a --filter "name=$(POSTGRES_CONTAINER_NAME)" --format "{{.Names}}" | grep -q '^$(POSTGRES_CONTAINER_NAME)$$'; then \
		echo "Container '$(POSTGRES_CONTAINER_NAME)' exists."; \
		docker start $(POSTGRES_CONTAINER_NAME); \
	else \
		echo "Container '$(POSTGRES_CONTAINER_NAME)' does not exist."; \
		echo "Creating and starting a new PostgreSQL container ..."; \
		docker run --name $(POSTGRES_CONTAINER_NAME) \
				-e POSTGRES_USER=$(POSTGRES_USER) \
				-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
				-e POSTGRES_DB=$(POSTGRES_DB) \
				-p $(POSTGRES_PORT):5432 \
				-d postgres:latest; \
		echo "PostgreSQL container $(POSTGRES_CONTAINER_NAME) started."; \
	fi

# Stop the PostgreSQL container
stop_postgres:
	@echo "Stopping PostgreSQL container..."
	# Check if the PostgreSQL container is running
	@if [ "$(shell docker ps -q -f name=$(POSTGRES_CONTAINER_NAME))" ]; then \
		docker stop $(POSTGRES_CONTAINER_NAME); \
		echo "PostgreSQL container $(POSTGRES_CONTAINER_NAME) stopped."; \
	else \
		echo "Container $(POSTGRES_CONTAINER_NAME) is not running."; \
	fi

.PHONY: start_postgres stop_postgres
