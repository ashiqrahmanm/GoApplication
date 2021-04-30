package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/students", returnAllStudents)
	myRouter.HandleFunc("/student", createNewStudent).Methods("POST")
	myRouter.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")
	myRouter.HandleFunc("/student/{id}", updateStudent).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

type Student struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

var Students []Student

func returnAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStudents")
	json.NewEncoder(w).Encode(Students)
}

func createNewStudent(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	var student Student
	json.Unmarshal(reqBody, &student)
	Students = append(Students, student)
	json.NewEncoder(w).Encode(student)
}
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	id := vars["id"]

	// we then need to loop through all our articles
	for index, student := range Students {
		// if our id path parameter matches one of our
		// articles
		if student.Id == id {
			// updates our Students array to remove the student
			Students = append(Students[:index], Students[index+1:]...)
		}
	}

}
func updateStudent(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	for index, student := range Students {

		if student.Id == id {
			reqBody, _ := ioutil.ReadAll(r.Body)
			fmt.Fprintf(w, "%+v", string(reqBody))
			var student Student
			json.Unmarshal(reqBody, &student)
			Students = append(Students, student)
			Students = append(Students[:index], Students[index+1:]...)
			json.NewEncoder(w).Encode(student)
		}
	}
}
func main() {
	Students = []Student{
		Student{Id: "1", Name: "name1"},
		Student{Id: "2", Name: "name2"},
		Student{Id: "3", Name: "name3"},
	}
	handleRequests()
}
