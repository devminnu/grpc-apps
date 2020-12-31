module github.com/devminnu/weatherapp/location

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/urfave/negroni v1.0.0
	go.mongodb.org/mongo-driver v1.4.4
	go.uber.org/zap v1.10.0
	google.golang.org/grpc v1.21.1
)

replace (
	github.com/devminnu/weatherapp/pb/location => /Users/minhaj/go/src/grpc-apps/weatherapp/pb/location
	github.com/devminnu/weatherapp/location/config => /Users/minhaj/go/src/grpc-apps/weatherapp/location/config
	github.com/devminnu/weatherapp/location/db => /Users/minhaj/go/src/grpc-apps/weatherapp/location/db
	github.com/devminnu/weatherapp/location/logger => /Users/minhaj/go/src/grpc-apps/weatherapp/location/logger
	github.com/devminnu/weatherapp/location/service => /Users/minhaj/go/src/grpc-apps/weatherapp/location/service
	github.com/devminnu/weatherapp/pb/location => /Users/minhaj/go/src/grpc-apps/weatherapp/pb/location
)
