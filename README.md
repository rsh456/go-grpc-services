Go gRPC Services 
=================
1.Rocket Service Mocks
> go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/rsh456/go-grpc-services/internal/rocket Store
as a comment where Store indicates the interface to be mocked

> go generate ./....

2.Test Service
> go test ./... -v
[alt text](https://github.com/rsh456/go-grpc-services/blob/master/images/test_1.PNG)
