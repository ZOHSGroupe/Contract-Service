package models

type Client struct {
	//ID          primitive.ObjectID `json:"id,omitempty"`
	Firstname   string `json:"firstname,omitempty" validate:"required"`
	Lastname    string `json:"lastname,omitempty" validate:"required"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	City        string `json:"city,omitempty" validate:"required"`
	StartDate   string `json:"start_date,omitempty" validate:"required,datetime=2006-01-02"`
	BirthDate   string `json:"birthDate,omitempty" validate:"required"`
	Gender      string `json:"gender,omitempty" validate:"required"`
	Nationality string `json:"nationality,omitempty" validate:"required"`
	Address     string `json:"address,omitempty" validate:"required"`
	NationalId  string `json:"nationalId,omitempty" validate:"required"`
}

type Vihecule struct {
	//ID                   primitive.ObjectID `json:"id,omitempty"`
	CurrentValue         string `json:"currentValue,omitempty" validate:"required"`
	CylinderCount        string `json:"cylinderCount,omitempty" validate:"required"`
	EmptyWeight          string `json:"emptyWeight,omitempty" validate:"required"`
	FuelType             string `json:"fuelType,omitempty" validate:"required"`
	Genre                string `json:"genre,omitempty" validate:"required"`
	GrossWeightRating    string `json:"grossWeightRating,omitempty" validate:"required"`
	ManufacturingDate    string `json:"manufacturingDate,omitempty" validate:"required"`
	Marque               string `json:"marque,omitempty" validate:"required"`
	TaxHorsePower        string `json:"taxHorsePower,omitempty" validate:"required"`
	Type                 string `json:"type,omitempty" validate:"required"`
	IdentificationNumber string `json:"IdentificationNumber,omitempty" validate:"required"`
}

type Permit struct {
	//ID            primitive.ObjectID `json:"id,omitempty"`
	IssueDate     string `json:"start_date,omitempty" validate:"required,datetime=2006-01-02"`
	EndDate       string `json:"end_date,omitempty" validate:"required,datetime=2006-01-02"`
	LicenceNumber string `json:"licenceNumber,omitempty" validate:"required"`
	Type          string `json:"type,omitempty" validate:"required"`
}

type Contract struct {
	//ID        primitive.ObjectID `json:"id,omitempty"`
	StartDate string   `json:"start_date,omitempty" `
	EndDate   string   `json:"end_date,omitempty" validate:"required,datetime=2006-01-02"`
	Value     string   `json:"value,omitempty" validate:"required"`
	Client    Client   `json:"client,omitempty"`
	Vihecule  Vihecule `json:"vihecule,omitempty"`
	Permit    Permit   `json:"permit,omitempty"`
}
