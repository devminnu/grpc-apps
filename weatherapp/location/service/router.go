package service

import (
	"fmt"
	"strconv"

	"github.com/devminnu/weatherapp/location/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
	appName       = "usermgr"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter(deps Dependencies) (router *mux.Router) {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	// authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("/login", loginEndpoint)
	// 	authorized.POST("/submit", submitEndpoint)
	// 	authorized.POST("/read", readEndpoint)

	// 	// nested group
	// 	testing := authorized.Group("testing")
	// 	testing.GET("/analytics", analyticsEndpoint)
	// }

	// Listen and serve on 0.0.0.0:8080
	//

	/* 	// mux router
	   	router := service.InitRouter(deps)

	   	// init web server
	   	server := negroni.Classic()
	   	server.UseHandler(router)
	*/
	port := config.AppPort() // This should be changed to the service port number via argument or environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	r.Run(addr)

	location := router.Group("/v1")
	{
		location.POST("/getLocation", getLocation)
		// location.POST("/submit", submitEndpoint)
		// location.POST("/read", readEndpoint)
	}

	/* 	root := mux.NewRouter()

	   	// Make a path prefix that will be assigned to this microservice
	   	router = root.PathPrefix("/usermgr").Subrouter()

	   	// Version 1 API management
	   	v1 := fmt.Sprintf("application/vnd.%s.v1", appName)

	   	router.HandleFunc("/user", createUserHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1) // create users

	   	// Note: Get user(s) will be supported only via gRPC

	   	return */
}
