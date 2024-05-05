dev:
	docker compose -f "docker-compose.dev.yaml" -p go-monolithic-boilerplate up 
swagger:
	swag init -g main.go --output ./docs --quiet --parseDependency --parseInternal
gorm:
	cd pkg/generate && go run gorm_gen.go
