package services

import (
	"log"

	"github.com/vladfin1/rest/data"
)

type unitDAO interface {
	Get() []*data.Unit
	GetByID(id string) *data.Unit
	Create(name string)
	Update(u data.Unit)
	Delete(id string)
	GetEmpl(id string, id2 string) *data.Unit
	GetEmpls(id string) []*data.Unit
}

type UnitService struct {
	dao unitDAO
}

func NewUnitService(dao unitDAO) *UnitService {
	return &UnitService{dao}
}

func (s *UnitService) Get() []*data.Unit {
	db := data.GetConnection()
	rows, err := db.Query("SELECT* FROM unit limit 11")
	if err != nil {
		log.Fatal(err)
	}
	var units []*data.Unit
	defer rows.Close()
	for rows.Next() {
		unit := new(data.Unit)
		err = rows.Scan(&unit.ID, &unit.Title)
		if err != nil {
			panic(err)
		}
		units = append(units, unit)
	}
	defer db.Close()
	return units
}

func (s *UnitService) GetByID(id string) data.Unit {
	db := data.GetConnection()
	rows, err := db.Query("SELECT* FROM unit WHERE unit_id  = " + id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		unit := data.Unit{}
		err = rows.Scan(&unit.ID, &unit.Title)
		if err != nil {
			panic(err)
		}
		return unit
	}
	defer rows.Close()
	defer db.Close()
	return data.Unit{}
}

func (s *UnitService) Upadate(unit data.Unit) {
	db := data.GetConnection()
	db.Exec("UPDATE unit SET name = '" + unit.Title + "' WHERE unit_id = " + unit.ID)
	defer db.Close()
}

func (s *UnitService) Delete(id string) {
	db := data.GetConnection()
	db.Exec("DELETE FROM unit WHERE emp_id  = " + id)
	defer db.Close()
}

func (s *UnitService) Create(name string) {
	db := data.GetConnection()
	db.Exec("INSERT INTO unit(name) VALUES('" + name + "')")
	defer db.Close()
}

func (s *UnitService) GetEmpls(id string) []*data.Employee {
	db := data.GetConnection()
	rows, err := db.Query("SELECT* FROM employee WHERE uid = " + id)
	if err != nil {
		log.Fatal(err)
	}
	var empls []*data.Employee
	defer rows.Close()
	for rows.Next() {
		emp := new(data.Employee)
		err = rows.Scan(&emp.ID, &emp.Lastname, &emp.Name, &emp.UnitID)
		if err != nil {
			panic(err)
		}
		empls = append(empls, emp)
	}
	defer db.Close()
	return empls
}

func (s *UnitService) GetEmpl(id string, id2 string) *data.Employee {
	db := data.GetConnection()
	rows, err := db.Query("SELECT* FROM employee WHERE uid = " + id + " AND emp_id = " + id2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	emp := new(data.Employee)
	for rows.Next() {
		err = rows.Scan(&emp.ID, &emp.Lastname, &emp.Name, &emp.UnitID)
		if err != nil {
			log.Fatal(err)
		}
		return emp
	}
	defer db.Close()
	return emp
}
