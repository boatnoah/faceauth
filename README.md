# Faceauth

Faceauth is toy project that authenticates users using their face.

# Endpoints

| Endpoint        | Method | Desc.             |
| --------------- | ------ | ----------------- |
| /v1/auth/login  | POST   | Log in an user    |
| /v1/auth/signup | POST   | Register new user |

This is inherently not safe.

- Very easily spoofable

        (REST)          (grpc)

  client <-> Go gateway <-> Python service

  faceauth/
  ├── proto/
  │ └── face/v1/face.proto # source of truth for the wire format
  ├── gen/
  │ ├── go/face/v1/ # generated Go stubs (committed or built)
  │ └── python/face/v1/ # generated Python stubs
  ├── cmd/api/ # Go gateway (already there)
  ├── internal/
  │ ├── client/grpc.go # thin wrapper around generated client
  │ └── ... # auth, store, etc.
  ├── services/
  │ └── face/ # Python gRPC server
  │ ├── pyproject.toml
  │ ├── Dockerfile
  │ └── face/server.py # implements FaceServiceServicer
  └── docker-compose.yml # add `face` service + grpc port
