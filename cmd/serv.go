package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transwarp/tos-app-market/pkg/setting"
	"transwarp/tos-app-market/pkg/task"
	"transwarp/tos-app-market/router"
	"transwarp/tos-app-market/router/middleware"
)

const servDesc = `
This command enable a TOS Application Market.

you can sp  the listen host and pot like :

    $ tos-app-market serv -a addr -p port

`

type ServCmd struct {
	configFile string
}

func (sc *ServCmd) initConfig() {
	logrus.Infof("loading configuration from [%s]", sc.configFile)
	setting.InitConfig(sc.configFile)
	settingConfig, err := json.Marshal(setting.Config)
	if err != nil {
		logrus.Fatal("failed to marshal setting config")
	}
	logrus.Infof("finished loading configuration: %s", string(settingConfig))
}

func NewServCmd() *cobra.Command {
	inst := &ServCmd{}

	cmd := &cobra.Command{
		Use:   "serv",
		Short: "enable a Walm Web Server",
		Long:  servDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return inst.run()
		},
	}
	cmd.PersistentFlags().StringVar(&inst.configFile, "config", "app-market.yaml", "config file (default is app-market.yaml)")

	return cmd
}

func (sc *ServCmd) run() error {
	sc.initConfig()
	stopChan := make(chan struct{})
	task.GetDefaultTaskManager().StartWorker()
	initRestApi()

	server := &http.Server{Addr: fmt.Sprintf(":%d", setting.Config.HttpConfig.HTTPPort), Handler: restful.DefaultContainer}
	go func() {
		logrus.Error(server.ListenAndServe())
	}()
	logrus.Info("walm server started")
	// server.ListenAndServeTLS()

	//shut down gracefully
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	err := server.Shutdown(context.Background())
	if err != nil {
		logrus.Error(err.Error())
	}
	task.GetDefaultTaskManager().StopWorker()
	close(stopChan)
	logrus.Info("waiting for workqueue stopping")
	time.Sleep(2 * time.Second)
	logrus.Info("walm server stopped gracefully")
	return nil
}

func initRestApi() {
	// accept and respond in JSON unless told otherwise
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	// gzip if accepted
	restful.DefaultContainer.EnableContentEncoding(true)
	// faster router
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.Filter(middleware.ServerStatsFilter)
	restful.Filter(middleware.RouteLogging)

	logrus.Infoln("Adding Route...")
	restful.Add(router.InitRootRouter())
	restful.Add(router.DockerImageRouter())
	restful.Add(router.ApplicationDevTemplateRouter())
	restful.Add(router.PackageChartRouter())
	restful.Add(router.WapperWalmChartRouter())
	logrus.Infoln("Add Route Success")
	config := restfulspec.Config{
		// You control what services are visible
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("swagger-ui/dist"))))
	http.Handle("/swagger/", http.RedirectHandler("/swagger-ui/?url=/apidocs.json", http.StatusFound))
	logrus.Infof("ready to serve on port %d", setting.Config.HttpConfig.HTTPPort)
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Tos App Market",
			Description: "TOS APP Market Server",
			Version:     "0.0.1",
		},
	}
}
