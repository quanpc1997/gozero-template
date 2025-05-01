# Swagger

```shell
# Install swagger
go install github.com/zeromicro/goctl-swagger@latest

# Generate swagger.json
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api src.api -dir .

# Run swagger using docker
docker run --rm -p 8083:8080 \
  -e SWAGGER_JSON=/foo/swagger.json \
  -v $PWD:/foo \
  swaggerapi/swagger-ui
```
