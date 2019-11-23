package cmd

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/cli"
)

//StartApp command
type StartApp struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *StartApp) Run(args []string) (rs int) {
	opts := struct {
		Host string
		Port string
	}{}
	flagset := flag.NewFlagSet("web", flag.ExitOnError)
	flagset.StringVar(&opts.Host, "address", "127.0.0.1", "web serve port")
	flagset.StringVar(&opts.Port, "port", "8008", "web serve port")
	flagset.Parse(args)
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, OK("test version"))
	})
	router.Run(opts.Host + ":" + opts.Port)
	return 0
}

//Synopsis Synopsis information
func (c *StartApp) Synopsis() string {
	return "Start web app"
}

//Help Help information
func (c *StartApp) Help() string {
	return "${app} web"
}
