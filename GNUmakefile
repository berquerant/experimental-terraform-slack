default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

re-generate: clean generate

clean: remove-swagger-configure-go

generate: swagger go-generate

.PHONY: go-generate
go-generate:
	go generate ./...

SWAGGER_SPEC = swagger.yml
SWAGGER_DIR = swagger
SWAGGER_CMD = $(SWAGGER_DIR)/cmd
SWAGGER_RESTAPI = $(SWAGGER_DIR)/restapi
SWAGGER_CLIENT = $(SWAGGER_DIR)/client
SWAGGER_PRINCIPAL = models.Principal

SWAGGER_NAME = slack-api-for-terraform-provider
SWAGGER_SERVER_NAME = $(SWAGGER_NAME)-server
SWAGGER_CLIENT_NAME = $(SWAGGER_NAME)-client
SWAGGER_CONFIGURE_GO = configure_$(subst -,_,$(SWAGGER_NAME)).go
SWAGGER_SH = tools/swagger.sh

.PHONY: swagger
swagger:
	$(SWAGGER_SH) run $(SWAGGER_SPEC) $(SWAGGER_DIR) $(SWAGGER_PRINCIPAL)
	# for mac sed
	sed -i '' "s|go:generate swagger|go:generate ../../$(SWAGGER_SH) do|" $(SWAGGER_RESTAPI)/$(SWAGGER_CONFIGURE_GO)

.PHONY: remove-swagger-configure-go
remove-swagger-configure-go:
	rm -f $(SWAGGER_RESTAPI)/$(SWAGGER_CONFIGURE_GO)

all: build api client docker

.PHONY: build
build: tmp
	go build -o tmp/terraform-provider-slack ./cmd/provider/main.go

.PHONY: api
api: tmp
	go build -o tmp/api-server ./$(SWAGGER_CMD)/$(SWAGGER_SERVER_NAME)

.PHONY: client
client: tmp
	go build -o tmp/api-client ./$(SWAGGER_CMD)/$(SWAGGER_CLIENT_NAME)

tmp:
	mkdir -p $@

API_IMAGE = "terraform-slack-api"
API_DOCKERFILE = "docker/api.dockerfile"

.PHONY: docker
docker:
	docker build --tag $(API_IMAGE) -f $(API_DOCKERFILE) .
