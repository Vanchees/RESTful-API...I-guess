// RESTful API...I guess project main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var successResponseOfVersion []Success_responses_of_version
var errorResponseOfVersion []Error_responses_of_version

var successResponseOfInterfaces []Success_responses_of_interfaces
var errorResponseOfInterfaces []Error_responses_of_interfaces

var successResponseOfSingleInterface []Success_responses_of_single_interface
var errorResponseOfSingleInterface []Error_responses_of_single_interface

type Success_responses_of_version struct {
	Code    string `json:",omitempty"`
	Content string `json:"version: ,omitempty"`
}
type Error_responses_of_version struct {
	Code    string `json:",omitempty"`
	Content string `json:"error:,omitempty"`
}

type Success_responses_of_interfaces struct {
	Code    string `json:",omitempty"`
	Content string `json:"interfaces: ,omitempty"`
}
type Error_responses_of_interfaces struct {
	Code    string `json:",omitempty"`
	Content string `json:"error:,omitempty"`
}

type Content_of_single_interface_response struct {
	name      string `json:"name:,omitempty"`
	hw_addr   string `json:"w_addr:,omitempty"`
	inet_addr string `json:"inet_addr:,omitempty"`
	MTU       string `json:"MTU:,omitempty"`
}
type Success_responses_of_single_interface struct {
	Code    string                                `json:",omitempty"`
	Content *Content_of_single_interface_response `json:"interfaces: ,omitempty"`
}
type Error_responses_of_single_interface struct {
	Code    string `json:",omitempty"`
	Content string `json:"error:,omitempty"`
}

func GetSuccessResponseOfVersion(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(successResponseOfVersion)
}
func GetErrorResponseOfVersion(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(errorResponseOfVersion)
}

func GetSuccessResponseOfInterfaces(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(successResponseOfInterfaces)
}
func GetErrorResponseOfInterfaces(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(errorResponseOfInterfaces)
}

func GetSuccessResponseOfSingleInterface(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(successResponseOfSingleInterface)
}
func GetErrorResponseOfSingleInterface(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(errorResponseOfSingleInterface)
}

func main() {
	router := mux.NewRouter()

	successResponseOfVersion = append(successResponseOfVersion, Success_responses_of_version{Code: "200", Content: "v1"})
	errorResponseOfVersion = append(errorResponseOfVersion, Error_responses_of_version{Code: "500", Content: "error message"})

	successResponseOfInterfaces = append(successResponseOfInterfaces, Success_responses_of_interfaces{Code: "200", Content: "['lo0', 'en1']"})
	errorResponseOfInterfaces = append(errorResponseOfInterfaces, Error_responses_of_interfaces{Code: "500", Content: "error message"})

	successResponseOfSingleInterface = append(successResponseOfSingleInterface, Success_responses_of_single_interface{Code: "200", Content: &Content_of_single_interface_response{name: "eth0", hw_addr: "02:42:ac:12:00:02", inet_addr: "['172.17.128.27/24', '2001:db8::1/32']", MTU: "1500"}})
	errorResponseOfSingleInterface = append(errorResponseOfSingleInterface, Error_responses_of_single_interface{Code: "500", Content: "error message"})

	router.HandleFunc("/service/version", GetSuccessResponseOfVersion).Methods("GET")
	router.HandleFunc("/service/:v1/interfaces", GetSuccessResponseOfInterfaces).Methods("GET")
	router.HandleFunc("/service/:v1/interface/:eth0", GetSuccessResponseOfSingleInterface).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))
}
