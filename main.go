// main.go
package main

import (
	"fmt"
	"github.com/akhil/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	r := httprouter.New()
	session, err := getSession()
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}
	uc := controllers.NewUserController(session)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.PUT("/user/:id", uc.UpdateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:5000", r)
}

func getSession() (*mgo.Session, error) {
	s, err := mgo.Dial("mongodb://localhost:27017/Go-Crud")
	if err != nil {
		return nil, err
	}
	return s, nil
}
