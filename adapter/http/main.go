package main

import (
	"context"
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yeganebagheri/Smart-RealEstate/adapter/postgres"
	"github.com/yeganebagheri/Smart-RealEstate/di"

	_ "github.com/yeganebagheri/Smart-RealEstate/adapter/http/rest/docs"
	"github.com/yeganebagheri/Smart-RealEstate/adapter/http/rest/middleware"
)

func init() {
	viper.AddConfigPath("C:/Uni/Uni Projects/clean-go-1.1.5/Smart-RealEstate/")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	//viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Vinícius Boscardin
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:port
// @BasePath /
func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	//postgres.RunMigrations()
	UserService := di.ConfigUserDI(conn)
	PostService := di.ConfigPostDI(conn)

	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(middleware.Cors)

	jsonApiRouter.Handle("/registeruser", http.HandlerFunc(UserService.Create)).Methods("POST", "OPTIONS")
	jsonApiRouter.Handle("/user", http.HandlerFunc(UserService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET", "OPTIONS")
	jsonApiRouter.Handle("/loginuser", http.HandlerFunc(UserService.Login)).Methods("POST", "OPTIONS")
	jsonApiRouter.Handle("/createpost", http.HandlerFunc(PostService.CreatePost)).Methods("POST", "OPTIONS")
	jsonApiRouter.Handle("/getpost", http.HandlerFunc(PostService.Get)).Methods("GET", "OPTIONS")
	// router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	// 	result := schema.ExecuteQuery(r.URL.Query().Get("query"), graphQLRouter)
	// 	json.NewEncoder(w).Encode(result)
	// })

	port := viper.GetString("server.port")

	if port == "" {
		port = os.Getenv("PORT")
	}
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
