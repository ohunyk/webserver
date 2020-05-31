package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ohunyk/webserver/datastore"
	"github.com/ohunyk/webserver/registry"
	"github.com/ohunyk/webserver/router"
	"github.com/spf13/viper"

	"log"
)

func init(){
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service run on Debug mode")
	}
}

func main(){

	db := datastore.NewMongo()
	r := registry.NewRegistry(db)
	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	// Run Server
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
