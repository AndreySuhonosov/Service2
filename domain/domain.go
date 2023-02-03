package domain

type CellRepository interface {
}

type DosimeterRepository interface {
	GetUsedDosimetersList(DispenserId string) []Dosimeter
}

type UserRepository interface {
	GetPersonData(PersonId string) Person
	SaveIssue(PersonId string) Person
}

// Описание ячейки УВД
type Cell struct {
	// Идентификатор ячейки
	CellID string `json:"Cell_ID,omitempty"`
	// Идентификатор УВД
	DispeneserID string `json:"dispeneserID,omitempty"`
	// Состояние ячейки
	State string `json:"state,omitempty"`
}

// Описание данных дозиметра
type Dosimeter struct {
	// Идентификатор дозиметра
	Id string `json:"id,omitempty"`
	// Название дозиметра
	Title string `json:"title,omitempty"`
	// Описание дозиметра
	Description string `json:"description,omitempty"`
	// Идентификатор УВД
	DispenserID string `json:"dispenserID,omitempty"`
	// Описание УВД - Место нахождения
	DispenserDesc string `json:"dispenserDesc,omitempty"`
}

// Описание работника
type Person struct {
	// Номер пропуска работника
	PassID string `json:"passID,omitempty"`
	// ФИО Работника
	Name string `json:"name,omitempty"`
	// Роль работника если есть: Администратор, инженер
	Role string `json:"role,omitempty"`
}
