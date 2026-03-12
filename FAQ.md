# FAQ

**Problem**: `go get -u github.com/marcpires/ecommerce-grpc-protos/golang/services/payment` fails with missing or invalid credentials.

**Solution:** Add the following line to you ~/.gitconfig
```
[url "ssh://git@github.com"]
  insteadOf = "https://github.com"
```

```sh
export GOPRIVATE=github.com/ecommerce-grpc
go get -u github.com/ecommerce-grpc/payment@v0.1.8
```
