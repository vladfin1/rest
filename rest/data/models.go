package data

type Unit struct {
	ID    string `json:"unit_id"`
	Title string `json:"name"`
}

type Employee struct {
	ID       string `json:"emp_id"`
	Name     string `json:"name"`
	Lastname string `json:"last_name"`
	UnitID   string `json:"unit_id"`
}
