FRONTEND_DIR = './web/app'

setup: setup-frontend setup-backend

setup-frontend:
	@cd $(FRONTEND_DIR) && yarn

setup-backend: mod

mod:
	@go mod download

env: env-local

env-local:
	@cp .env .env.local

clean:
	@rm -rf $(FRONTEND_DIR)/build

generate: generate-backend generate-frontend

generate-frontend:
	@docker compose up --remove-orphans --build -d openapi-client
	@rm -rf web/app/src/api/generated && true
	@docker compose cp openapi-client:/out web/app/src/api/generated
	@docker compose stop openapi-client

generate-backend:
	@docker compose up --remove-orphans --build -d openapi-server
	@rm -rf internal/app/generated && true
	@docker compose cp openapi-server:/out internal/app/generated
	@docker compose stop openapi-server

build: build-backend build-frontend

build-backend:
	@mkdir -p ./build
	@go build -v -o ./build/pinman main.go

build-frontend:
	@cd $(FRONTEND_DIR) && yarn build

run: clean generate build-frontend
	@SPA_PATH=./web/app/dist ENV_FILE=.env.local go run main.go

run-backend:
	@ENV_FILE=.env.local go run main.go

run-frontend:
	@cd $(FRONTEND_DIR) && VITE_API_HOST=http://localhost:8080/api yarn start

test: test-backend test-frontend

test-setup:
	@go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo

test-frontend:
	@cd $(FRONTEND_DIR) && yarn test
	@rm -rf ./reports/ts && true
	@mkdir -p ./reports/ts
	@mv -f $(FRONTEND_DIR)/coverage $(FRONTEND_DIR)/reports ./reports/ts/

test-backend: test-setup
	@ginkgo	./...

test-backend-cov: test-setup
	@ginkgo --cover \
		--race \
		--json-report=report.json \
		--output-dir=reports/go \
		--skip-package generated \
		--skip-file mock \
		./...
