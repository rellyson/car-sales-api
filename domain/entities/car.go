package entities

type CarType string

const (
	Hatch CarType = "HATCH"
	Sedan CarType = "SEDAN"
	SUV   CarType = "SUV"
	Truck CarType = "Truck"
)

type Car struct {
	BaseEntity
	Brand           string         `json:"brand"`
	Model           string         `json:"model"`
	Type            CarType        `json:"type"`
	ManufactureYear any            `json:"manufacture_year"`
	ModelYear       any            `json:"model_year"`
	Accessories     map[string]any `json:"accessories"`
}
