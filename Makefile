LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=auth user=auth-user password=auth-password sslmode=disable"

.PHONY: build
build: install-go-deps generate vendor-proto
	CGO_ENABLED=0 GOOS=linux go build -o bin/auth-api cmd/server/main.go

install-go-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@latest

swagger:
	mkdir -p pkg/swagger

generate: swagger generate-user-api generate-auth-api generate-access-api
	statik -src=pkg/swagger/ -include='*.css,*html,*.json,*.png,*.js'

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 --proto_path vendor.protogen \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
    --plugin=protoc-gen-validate=bin/protoc-gen-validate \
    --grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
    --plugin=protoc-gen-go-grpc-gateway=bin/protoc-gen-go-grpc-gateway \
    --openapiv2_out=allow_merge=true,merge_file_name=user:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/user_v1/user.proto

generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 --proto_path vendor.protogen \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/auth_v1 --validate_opt=paths=source_relative \
    --plugin=protoc-gen-validate=bin/protoc-gen-validate \
    --grpc-gateway_out=pkg/auth_v1 --grpc-gateway_opt=paths=source_relative \
    --plugin=protoc-gen-go-grpc-gateway=bin/protoc-gen-go-grpc-gateway \
    --openapiv2_out=allow_merge=true,merge_file_name=auth:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/auth_v1/auth.proto

generate-access-api:
	mkdir -p pkg/access_v1
	protoc --proto_path api/access_v1 --proto_path vendor.protogen \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go_out=pkg/access_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/access_v1 --validate_opt=paths=source_relative \
    --plugin=protoc-gen-validate=bin/protoc-gen-validate \
    --grpc-gateway_out=pkg/access_v1 --grpc-gateway_opt=paths=source_relative \
    --plugin=protoc-gen-go-grpc-gateway=bin/protoc-gen-go-grpc-gateway \
    --openapiv2_out=allow_merge=true,merge_file_name=access:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/access_v1/access.proto

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

lint:
	golangci-lint run

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

gen:
	mockgen -source=internal/repository/user/repository.go \
	-destination=internal/repository/mocks/mock_repository.go
	mockgen -source=internal/repository/access/repository.go \
	-destination=internal/repository/mocks/mock_repository.go
	mockgen -source=internal/repository/auth/repository.go \
	-destination=internal/repository/mocks/mock_repository.go

cert:
	openssl genrsa -out ca.key 4096
	openssl req -new -x509 -key ca.key -sha256 -subj "/C=US/ST=NJ/O=Test, Inc." -days 365 -out ca.cert
	openssl genrsa -out service.key 4096
	openssl req -new -key service.key -out service.csr -config certificate.conf
	openssl x509 -req -in service.csr -CA ca.cert -CAkey ca.key -CAcreateserial \
    		-out service.pem -days 365 -sha256 -extfile certificate.conf -extensions req_ext

grpc-load-test:
	ghz \
		--proto api/user_v1/user.proto \
		--import-paths=vendor.protogen \
		--call user_v1.UserV1.Get \
		--data '' \
		--rps 100 \
		--total 1000 \
		--cacert=service.pem \
		localhost:50051