# ogen

## 生成

```go
ogen --target ogen -package ogen --clean openapi.yml
```

## リクエスト

```bash
$ curl -X GET http://localhost:8080/pets
[]
```

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id":1, "name":"sensuikan1973"}' localhost:8080/pets
```

```bash
$ curl -X GET http://localhost:8080/pets/1
{"code":1,"message":"error"}
```
