package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aidapedia/go-routeros"
	"github.com/aidapedia/go-routeros/driver"
	"github.com/aidapedia/go-routeros/model"
	"github.com/aidapedia/go-routeros/module"
	"github.com/julienschmidt/httprouter"
)

type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}

func main() {
	routerBuilder := routeros.NewRouterOS(&routeros.Options{
		Address:  os.Getenv("ROUTEROS_ADDRESS"),
		Username: os.Getenv("ROUTEROS_USERNAME"),
		Password: os.Getenv("ROUTEROS_PASSWORD"),
	})
	errs := routerBuilder.Connect()
	if errs != nil {
		log.Fatal(errs)
	}
	defer routerBuilder.Close()

	active, errs := driver.New(routerBuilder, module.HotspotActiveModule)
	if errs != nil {
		log.Fatal(errs)
	}

	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		response := JsonErrorResponse{Error: &ApiError{Status: 500, Title: "Internal Server Error"}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}
	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "pong")
	})
	router.POST("/reset", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var body struct {
			User string `json:"user"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := JsonErrorResponse{Error: &ApiError{Status: 400, Title: "Bad Request"}}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				panic(err)
			}
		}
		activeRes, err := active.Print(context.Background(), model.PrintRequest{
			Where: []model.Where{
				{
					Field:    "user",
					Value:    body.User,
					Operator: model.OperatorEqual,
				},
			},
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := JsonErrorResponse{Error: &ApiError{Status: 500, Title: "Internal Server Error"}}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				panic(err)
			}
		}

		if len(activeRes) == 0 {
			w.WriteHeader(http.StatusNotFound)
			response := JsonErrorResponse{Error: &ApiError{Status: 404, Title: "Record Not Found"}}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				panic(err)
			}
		}

		for _, record := range activeRes {
			req, ok := record.(*module.HotspotActiveData)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				response := JsonErrorResponse{Error: &ApiError{Status: 500, Title: "Internal Server Error"}}
				if err := json.NewEncoder(w).Encode(response); err != nil {
					log.Print()
				}
			}
			err := active.Remove(context.Background(), req.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				response := JsonErrorResponse{Error: &ApiError{Status: 500, Title: "Internal Server Error"}}
				if err := json.NewEncoder(w).Encode(response); err != nil {
					panic(err)
				}
			}
		}

	})

	log.Println("Server On : Listening")
	errs = http.ListenAndServe(":38000", router)
	if errs != nil {
		log.Fatal(errs)
	}
}
