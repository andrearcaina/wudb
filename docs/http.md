# HTTP REST API

For now, an HTTP REST API is provided for basic operations.

Basic operations include:
- Set a key-value pair: `POST /set` with JSON body `{"key": "your_key", "value": "your_value"}`
- Get a value by key: `GET /get/{key}`
- Delete a key-value pair: `DELETE /del/{key}`

```bash
> curl -X POST -d '{"key": "name", "value": "John"}' http://localhost:8080/set
# {"key":"name","value":"John"}
> curl http://localhost:8080/get/name
# {"value":"John"}
> curl -X DELETE http://localhost:8080/del/name
# no response (204 No Content)
> curl http://localhost:8080/get/name
# key not found
```

Later on, I plan to use a different protocol, maybe similar to how Redis works.