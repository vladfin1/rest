package apis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vladfin1/rest/data"

	"github.com/gorilla/mux"
	"github.com/vladfin1/rest/services"
)

type (
	unitResource struct {
		service services.UnitService
	}
	emplResource struct {
		service services.EmplService
	}
)

func ServeResource(router *mux.Router, service services.UnitService, service2 services.EmplService) {
	r := &unitResource{service}
	r2 := &emplResource{service2}
	router.HandleFunc("/units", r.units).Methods("GET")
	router.HandleFunc("/units/{id}", r.unit).Methods("GET")
	router.HandleFunc("/unit/{id}/{id2}", r.employeeFromUnit).Methods("GET")
	router.HandleFunc("/unit/{id}", r.employeesFromUnit).Methods("GET")
	router.HandleFunc("/units", r.createUnit).Methods("POST")
	router.HandleFunc("/units", r.updateUnit).Methods("PUT")
	router.HandleFunc("/units/{id}", r.deleteUnit).Methods("DELETE")
	router.HandleFunc("/employees", r2.employees).Methods("GET")
	router.HandleFunc("/employees/{id}", r2.employee).Methods("GET")
	router.HandleFunc("/employees", r2.createEmployee).Methods("POST")
	router.HandleFunc("/employees", r2.updateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", r2.deleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (r *emplResource) employees(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r.service.Get())
}

func (r *unitResource) units(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r.service.Get())
}

func (r *emplResource) employee(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(re)
	json.NewEncoder(w).Encode(r.service.GetByID(params["id"]))
}

func (r *unitResource) unit(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(re)
	json.NewEncoder(w).Encode(r.service.GetByID(params["id"]))
}

func (r *unitResource) employeeFromUnit(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(re)
	json.NewEncoder(w).Encode(r.service.GetEmpl(params["id"], params["id2"]))
}

func (r *emplResource) createEmployee(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empl data.Employee
	reqBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &empl)
	r.service.Create(empl.Name, empl.Lastname, empl.UnitID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(empl)
}

func (r *emplResource) updateEmployee(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empl data.Employee
	reqBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &empl)
	r.service.Upadate(empl)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(empl)
}

func (r *emplResource) deleteEmployee(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(re)
	r.service.Delete(params["id"])
	w.WriteHeader(http.StatusOK)
}

func (r *unitResource) createUnit(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var unit data.Unit
	reqBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &unit)
	r.service.Create(unit.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(unit)
}

func (r *unitResource) updateUnit(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var unit data.Unit
	reqBody, err := ioutil.ReadAll(re.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &unit)
	r.service.Upadate(unit)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(unit)
}

func (r *unitResource) deleteUnit(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(re)
	r.service.Delete(params["id"])
	w.WriteHeader(http.StatusOK)
}

func (r *unitResource) employeesFromUnit(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(re)
	json.NewEncoder(w).Encode(r.service.GetEmpls(params["id"]))
}
