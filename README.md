Go gRPC Services Course
========================
1.Rocket Service Mocks
> go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/rsh456/go-grpc-services/internal/rocket Store
as a comment where Store indicates the interface to be mocked

> go generate ./....

