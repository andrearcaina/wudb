## wuDB

simple lightweight in-memory key-value store built in Go

testing basic features rn (basic REST API and persistence)

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

### todo:
no particular order 
- [X] persistence
- [ ] advanced querying
- [ ] ACID compliance
- [ ] distributed support (use raft consensus algorithm)