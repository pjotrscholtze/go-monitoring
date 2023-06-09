package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gotify/go-api-client/v2/models"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/gotifyservice"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/checkmanager"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/config"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/controller"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/entity"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/informer"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/repo"
	"github.com/pjotrscholtze/go-monitoring/restapi"
	"github.com/pjotrscholtze/go-monitoring/restapi/operations"
)

// "github.com/pjotrscholtze/go-buildserver/cmd/go-buildserver/process"

//	func main() {
//		// process.StartProcessWithStdErrStdOutCallback("/bin/sh",
//		// 	[]string{path.Join("/home/pjotr/go/src/github.com/pjotrscholtze/go-monitoring", "boot.sh")},
//		// 	func(pt process.PipeType, t time.Time, s string) {
//		// 		println(s)
//		// 	})
//		cui := informer.NewCheckUpdateInformer()
//		tcr := repo.NewTargetCheckRepoInMemory(5)
//		cui.RegisterListenerFunc(func(result entity.Result, target config.Target, check config.Check) {
//			tcr.UpdateCheck(result, target, check)
//			fmt.Printf("via informer, we got %s, %s, %s\n", result.Message(), target.Name, check.Name)
//		})
//		cm := checkmanager.NewCheckManager(config.LoadMockConfig(), cui)
//		cm.ValidateConfig()
//		cm.Run()
//		// err := cm.PerformAllChecks()
//		// if err != nil {
//		// 	log.Fatal(err.Error())
//		// }
//		for {
//			time.Sleep(time.Second)
//			for _, tce := range tcr.List() {
//				println(" from main")
//				tce.Result.Log()
//			}
//		}
//	}
type mock struct {
	next http.Handler
	mux  *http.ServeMux
}

func (m *mock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix((*r).RequestURI, "/api/") ||
		strings.HasPrefix((*r).RequestURI, "/swagger.json") {
		m.next.ServeHTTP(w, r)
		return
	}
	m.mux.ServeHTTP(w, r)
}

func main() {
	if len(os.Args) != 2 {
		println("Usage: app <config-path.yaml>")
		return
	}
	path := os.Args[1]
	log.Println("Starting Go monitoring server")

	log.Printf("Loading config: %s\n", path)
	c := config.LoadConfigFromPath(path)
	// c := config.LoadConfigFromPath("../../example/example.yaml")
	// c := config.LoadMockConfig()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGoMonitoringAPI(swaggerSpec)
	server := restapi.NewServer(api)
	cui := informer.NewCheckUpdateInformer()
	tcr := repo.NewTargetCheckRepoInMemory(c.MaxHistoryInMemory)
	var gs gotifyservice.GotifyService
	if c.Gotify != nil {
		gs = gotifyservice.NewGotifyService(c.Gotify.GotifyURL, c.Gotify.ApplicationToken)
	}
	cui.RegisterListenerFunc(func(result entity.Result, target config.Target, check config.Check) {
		tcr.UpdateCheck(result, target, check)
		prio := 5
		succesfull := "Not ok"
		if result.Success() {
			prio = 1
			succesfull = "ok"
		}

		if gs != nil && check.DisableGotifyForSuccessfulCheck != true {
			err := gs.SendMessage(
				&models.MessageExternal{
					Title:    fmt.Sprintf("Target %s, check %s: %s\n", target.Name, check.Name, succesfull),
					Message:  result.Message(),
					Priority: prio,
				})

			if err != nil {
				log.Fatalf("Could not send message %v", err)
				return
			}
		}
	})
	controller.ConnectAPI(api, tcr)
	cm := checkmanager.NewCheckManager(c, cui)
	cm.ValidateConfig()
	cm.Run()

	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Go Monitoring server"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()
	server.Port = c.HTTPPort
	server.Host = c.HTTPHost

	t := &mock{
		next: api.Serve(nil),
		mux:  controller.RegisterUIController(),
	}
	server.SetHandler(t)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
