package main

import (
	"flag"
	"fmt"
	"github.com/architagr/workflow/models"
	routerVersion1 "github.com/architagr/workflow/router/version1"
	"github.com/gin-gonic/gin"
	"os"
)

var configuration models.Configuration

func main() {
	//get Environment for which the api will run
	//this is to help identifu with config file is to be read
	env := flag.String("env", "dev", "environment to be used")
	flag.Parse()

	fmt.Printf("Envirnment used is : %s\n", *env)

	//get config data according to falg been set
	configuration, err := configuration.Init(*env)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("config file issue for environment %s\n", *env)
		os.Exit(0)
	}

	//start gin server with default middleware
	r := gin.Default()
	//configure routes begin

	//confugure routes for api version1
	routerVersion1.RouterVersion1(r)

	//configure route end

	//run the api server on the mentioned id and port in the config file
	r.Run(configuration.Ip + ":" + configuration.Port)
}
