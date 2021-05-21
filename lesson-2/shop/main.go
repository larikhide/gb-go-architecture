package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"shop/pkg/email"
	"shop/pkg/tgbot"
	"shop/repository"
	"shop/service"
)

func main() {

	flagConfigFile := flag.String("config", "./config-with-passwords.yaml", "File config in yaml format")
	flag.Parse()

	config, err := ReadConfig(*flagConfigFile)
	if err != nil {
		panic(fmt.Sprintf("Not read config file. %s", err))
	}

	em, err := email.NewSMTPClient(config.Host, config.Username, config.Password)
	if err != nil {
		log.Fatal("Unable to init smtp client")
	}

	db := repository.NewMapDB()

	tg, err := tgbot.NewTelegramAPI(config.Token, config.ChatID)
	if err != nil {
		log.Fatal("Unable to init telegram bot")
	}

	service := service.NewService(em, tg, db)
	handler := &shopHandler{
		service: service,
		db:      db,
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	router.HandleFunc("/order", handler.createOrderHandler).Methods("POST")
	router.HandleFunc("/order/{id}", handler.getOrderHandler).Methods("GET")

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
