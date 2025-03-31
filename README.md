# tarantool_crud
Реализация CRUD API с использование tarantool

# Usage
1. docker compose up
2. go run cmd/app/main.go

## Examples

Примеры запросов
- [Create](#create)
- [Read](#read)
- [Update](#update)
- [Delete](#delete)

### Create <a name="create"></a>
Request:
```curl
curl -i -X POST 'http://localhost:8081/kv' \
-H 'Content-Type: application/json' \
-d '{
  "key": "mobile_app_v1",
  "value": {
    "dark_mode": { "enabled": true, "default": false },
    "experimental": {
      "ai_recommendations": true,
      "custom_themes": false
    }
  }
}'
```
Response:
```bash
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sat, 29 Mar 2025 16:09:01 GMT
Content-Length: 144

{"key":"mobile_app_v1","value":{"dark_mode":{"enabled":true,"default":false},"experimental":{"ai_recommendations":true,"custom_themes":false}}}
```

### Read <a name="read"></a>
Request:
```curl
curl -i -X GET 'http://localhost:8081/kv/mobile_app_v1'
```

Response:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 29 Mar 2025 16:53:36 GMT
Content-Length: 144

{"key":"mobile_app_v1","value":{"dark_mode":{"enabled":true,"default":false},"experimental":{"ai_recommendations":true,"custom_themes":false}}}
```

### Update <a name="update"></a>
Request:
```curl
curl -i -X PUT 'http://localhost:8081/kv/mobile_app_v1' \
-H 'Content-Type: application/json' \
-d '{
  "value": {
    "dark_mode": { "enabled": true, "default": false },
    "experimental": {
      "ai_recommendations": true,
      "custom_themes": false,
      "new_feature": false
    }
  }
}'
```

Response:
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 29 Mar 2025 17:02:30 GMT
Content-Length: 164

{"key":"mobile_app_v1","value":{"dark_mode":{"enabled":true,"default":false},"experimental":{"ai_recommendations":true,"custom_themes":false,"new_feature":false}}}
```

### Delete <a name="delete"></a>
Request:
```curl
curl -i -X DELETE 'http://localhost:8081/kv/mobile_app_v1'
```

Response:
```bash
HTTP/1.1 200 OK
Date: Sat, 29 Mar 2025 17:05:02 GMT
Content-Length: 0
```