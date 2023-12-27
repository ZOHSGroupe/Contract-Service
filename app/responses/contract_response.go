package responses

/*
type ContractResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Location  string             `json:"location"`
	Title     string             `json:"title"`
	StartDate string             `json:"start_date"`
	EndDate   string             `json:"end_date"`
	Value     string             `json:"value"`
	Client    struct {
		Firstname   string `json:"firstname"`
		Lastname    string `json:"lastname"`
		Email       string `json:"email"`
		City        string `json:"city"`
		Nationality string `json:"nationality"`
		Gender      string `json:"gender"`
		DateOfBirth string `json:"date_of_birth"`
	} `json:"client"`
	Vihecule struct {
		CurrentValue         string `json:"current_value"`
		CylinderCount        string `json:"cylinder_count"`
		EmptyWeight          string `json:"empty_weight"`
		FuelType             string `json:"fuel_type"`
		Genre                string `json:"genre"`
		GrossWeightRating    string `json:"gross_weight_rating"`
		ManufacturingDate    string `json:"manufacturing_date"`
		Marque               string `json:"marque"`
		TaxHorsePower        string `json:"tax_horse_power"`
		Type                 string `json:"type"`
		IdentificationNumber string `json:"identification_number"`
	} `json:"vihecule"`
	Permit struct {
		IssueDate     string `json:"start_date"`
		EndDate       string `json:"end_date"`
		LicenceNumber string `json:"licence_number"`
		Type          string `json:"type"`
	} `json:"permit"`
}
*/

type ContractResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
