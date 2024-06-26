REGION = eu-west-1
ENV ?= test
SERVICE_NAME = worker-report-document-linker

lambda: clean
	@env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/bootstrap ./cmd/lambda/main.go
	@cd ./bin && zip $(SERVICE_NAME).zip bootstrap

deploy: lambda
	$(eval VERSION = $(shell aws lambda update-function-code --function-name ${SERVICE_NAME} --region ${REGION} --zip-file fileb://bin/$(SERVICE_NAME).zip --publish |  jq .Version))
	@aws lambda update-alias --function-name ${SERVICE_NAME} --region ${REGION} --name ${ENV} --function-version $(VERSION)

clean:
	@echo "clean ${SERVICE_NAME}"
	@rm -rf bin

