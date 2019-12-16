package services

import (
	"log"

	"github.com/vladfin1/rest/data"
)

type emplDAO interface {
	Get() []*data.Employee
	GetByID(id string) *data.Employee
	Create(name string)
	Update(u data.Employee)
	Delete(id string)
}

type EmplService struct {
	dao emplDAO
}

func NewEmplService(dao emplDAO) *EmplService {
	return &EmplService{dao}
}

func (s *EmplService) Get() []*data.Employee {
	db := data.GetConnection()
	rows, err := db.Query("SELECT* FROM employee limit 11")
	if err != nil {
		log.Fatal(err)
	}
	var empls []*data.Employee
	defer rows.Close()
	for rows.Next() {
		employee := new(data.Employee)
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Lastname, &employee.UnitID)
		if err != nil {
			panic(err)
		}
		empls = append(empls, employee)
	}
	defer db.Close()
	return empls
}

func (s *EmplService) GetByID(id string) data.Employee {
	db := data.GetConnection()
	rows, err := db.Query("SELECT* FROM employee WHERE emp_id  = " + id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		employee := data.Employee{}
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Lastname, &employee.UnitID)
		if err != nil {
			panic(err)
		}
		return employee
	}
	defer rows.Close()
	defer db.Close()
	return data.Employee{}
}

func (s *EmplService) Upadate(emp data.Employee) {
	db := data.GetConnection()
	db.Exec("UPDATE employee SET name = '" + emp.Name + "', last_name = '" + emp.Lastname + "', uid = '" + emp.UnitID + "' WHERE emp_id = " + emp.ID)
	defer db.Close()
}

func (s *EmplService) Delete(id string) {
	db := data.GetConnection()
	db.Exec("DELETE FROM employee WHERE emp_id  = " + id)
	defer db.Close()
}

func (s *EmplService) Create(name string, lname string, uid string) {
	db := data.GetConnection()
	db.Exec("INSERT INTO employee(name, last_name, uid) VALUES('" + name + "' , '" + lname + "', '" + uid + "')")
	defer db.Close()
}
