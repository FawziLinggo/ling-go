package main

import (
	"main/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	userController := controllers.NewUserController(getSession())
	r.GET("/user/:id", userController.GetUser)
	r.POST("/user", userController.CreateUser)
	r.DELETE("/user/:id", userController.DeleteUser)
	http.ListenAndServe("fawzi.linggo.id:30451", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://master.k8s.alldataint.com:30002/streamhatchet?&authSource=admin&replicaSet=devops_adi")
	if err != nil {
		panic(err)
	}
	return s
}
