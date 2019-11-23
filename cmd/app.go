package cmd

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/cli"
	"github.com/smarkm/k8s-nms/api"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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

	kconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kconfig)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return 0
	}

	router := gin.Default()
	router.Static("/app/", os.Getenv("K8s_NMS_Home")+"web/")
	apig := router.Group("/api")
	apig.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, api.OK("test version"))
	})
	apig.GET("/network/ns", func(c *gin.Context) {
		ns := c.Query("ns")
		data, err := api.PodNetwork(clientset, ns)
		if err != nil {
			c.JSON(http.StatusOK, api.Error(1, "err", err.Error()))
			return
		}

		c.JSON(http.StatusOK, api.OK(data))
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
