package cmd

import (
	"github.com/gorilla/mux"
	"github.com/JunboLi-Ray/go-demo/app/action"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"net/http"
	"github.com/JunboLi-Ray/go-demo/app/config"
	"log"
	"github.com/JunboLi-Ray/go-demo/app/service"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		os.Exit(0)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE)
	go func() {
		for {
			sig := <-c
			fmt.Println("Received ", sig)
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				service.CloseDb()
				os.Exit(0)
			}
		}
	}()
	r := mux.NewRouter()
	r.HandleFunc("/users", action.AllNormalUserFunc).Methods("GET")
	r.HandleFunc("/users", action.AddNormalUserFunc).Methods("POST")
	r.HandleFunc("/users/{user_id}/relationships", action.UserAllLikeRelaFunc).Methods("GET")
	r.HandleFunc("/users/{user_id}/relationships/{other_user_id}", action.UpdateLikeRelaFunc).Methods("PUT")
	log.Fatal(http.ListenAndServe(config.SysConfs.ServerHost, r))
}
