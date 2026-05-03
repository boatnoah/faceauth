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
