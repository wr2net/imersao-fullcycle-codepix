# CodePix 

### Run Services
```bash
docker-compose up -d
```

### Create Protocol Buffer
```bash
docker exec -it codepix-app-1 protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
```

### Restart Services
```bash
docker-compose restart
```

### Access App
```bash
docker exec -it codepix-app-1 bash
```

### Run App
```bash
docker exec -it codepix-app-1 go run main.go
```

### Run Evans
```bash
docker exec -it codepix-app-1 evans -r repl
```

* Register Pix Key
  * ```bash call RegisterPixKey ```
* Find Pix Key
  * ```bash call Find ```


### Run Tests
```bash
docker exec -it codepix-app-1 go test ./...
```

### Stop Services
```bash
docker-compose down
```