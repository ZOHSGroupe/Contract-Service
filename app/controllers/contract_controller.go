package controllers

import (
	"Contract-Service/app/configs"
	"Contract-Service/app/models"
	"Contract-Service/app/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var contractCollection *mongo.Collection = configs.GetCollection(configs.DB, "contract")
var validate = validator.New()

func GetAllContracts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Fetch all contracts from the collection
		cursor, err := contractCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer cursor.Close(ctx)

		var contracts []models.Contract
		if err := cursor.All(ctx, &contracts); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ContractResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"contracts": contracts}})
	}
}

func CreateContract() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var contract models.Contract

		// Validate the request body
		if err := c.BindJSON(&contract); err != nil {
			c.JSON(http.StatusBadRequest, responses.ContractResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Use the validator library to validate required fields
		if validationErr := validate.Struct(&contract); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ContractResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Check if a valid contract with the same Vihecule IdentificationNumber exists
		existingContract, err := findValidContractByViheculeIdentificationNumber(ctx, contract.Vihecule.IdentificationNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if existingContract != nil {
			c.JSON(http.StatusBadRequest, responses.ContractResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "A valid contract with the same Vihecule IdentificationNumber already exists"}})
			return
		}

		// Check if the contract's EndDate is less than the current date
		endDate, err := time.Parse("2006-01-02", contract.EndDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		currentDate := time.Now().UTC()

		if endDate.Before(currentDate) {
			c.JSON(http.StatusBadRequest, responses.ContractResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Contract EndDate must be greater than or equal to the current date"}})
			return
		}
		// Set StartDate to the current date with the specified format
		contract.StartDate = time.Now().UTC().Format("2006-01-02")

		newContract := models.Contract{
			StartDate: contract.StartDate,
			EndDate:   contract.EndDate,
			Value:     contract.Value,
			Client: models.Client{
				Firstname:   contract.Client.Firstname,
				Lastname:    contract.Client.Lastname,
				Email:       contract.Client.Email,
				City:        contract.Client.City,
				Nationality: contract.Client.Nationality,
				Gender:      contract.Client.Gender,
				BirthDate:   contract.Client.BirthDate,
				NationalId:  contract.Client.NationalId,
			},
			Vihecule: models.Vihecule{
				CurrentValue:         contract.Vihecule.CurrentValue,
				CylinderCount:        contract.Vihecule.CylinderCount,
				EmptyWeight:          contract.Vihecule.EmptyWeight,
				FuelType:             contract.Vihecule.FuelType,
				Genre:                contract.Vihecule.Genre,
				GrossWeightRating:    contract.Vihecule.GrossWeightRating,
				ManufacturingDate:    contract.Vihecule.ManufacturingDate,
				Marque:               contract.Vihecule.Marque,
				TaxHorsePower:        contract.Vihecule.TaxHorsePower,
				Type:                 contract.Vihecule.Type,
				IdentificationNumber: contract.Vihecule.IdentificationNumber,
			},
			Permit: models.Permit{
				IssueDate:     contract.Permit.IssueDate,
				EndDate:       contract.Permit.EndDate,
				LicenceNumber: contract.Permit.LicenceNumber,
				Type:          contract.Permit.Type,
			},
		}

		// Insert the new contract into the database
		result, err := contractCollection.InsertOne(ctx, newContract)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Extract the inserted ID from the result
		insertedID, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Failed to extract inserted ID"}})
			return
		}

		// Include the ID in the response
		c.JSON(http.StatusCreated, responses.ContractResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"id": insertedID.Hex()}})
	}
}

// Helper function to find a valid contract by Vihecule IdentificationNumber
func findValidContractByViheculeIdentificationNumber(ctx context.Context, identificationNumber string) (*models.Contract, error) {
	cursor, err := contractCollection.Find(ctx, bson.M{"vihecule.identificationnumber": identificationNumber})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var contracts []models.Contract
	if err := cursor.All(ctx, &contracts); err != nil {
		return nil, err
	}

	// Check each contract individually
	for _, contract := range contracts {
		// Check if the contract is valid based on your criteria
		if isValidContract(&contract) {
			return &contract, nil
		}
	}

	// No valid contract found
	return nil, nil
}

// Function to check if a contract is valid based on your criteria
func isValidContract(contract *models.Contract) bool {
	// Add your validation logic here
	endDate, err := time.Parse("2006-01-02", contract.EndDate)
	if err != nil {
		return false
	}

	currentDate := time.Now().UTC()
	return endDate.After(currentDate)
}

func GetAContract() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get the contractID parameter from the URL
		contractID := c.Param("id")

		// Convert the contractID to an ObjectID
		objID, err := primitive.ObjectIDFromHex(contractID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ContractResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": "Invalid contract ID format"},
			})
			return
		}

		// Find the contract by its _id
		var contract models.Contract
		err = contractCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&contract)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Return the contract in the response
		c.JSON(http.StatusOK, responses.ContractResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": contract},
		})
	}
}

func DeleteAContract() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		contractID := c.Param("id")
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(contractID)

		result, err := contractCollection.DeleteOne(ctx, bson.M{"_id": objID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.ContractResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "contract with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.ContractResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "contract successfully deleted!"}},
		)
	}

}

func GetContractsByNationalID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		nationalID := c.Param("nationalId") // Assuming the parameter is part of the URL path

		// Validate nationalID if needed

		// Find all contracts for the given national ID
		cursor, err := contractCollection.Find(ctx, bson.M{"client.nationalid": nationalID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer cursor.Close(ctx)

		var contracts []models.Contract
		if err := cursor.All(ctx, &contracts); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Check if any contracts were found
		if len(contracts) == 0 {
			c.JSON(http.StatusNotFound, responses.ContractResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No contracts found with the given national ID"}})
			return
		}

		c.JSON(http.StatusOK, responses.ContractResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"contracts": contracts}})
	}
}

func GetContractsByVehicleIdentificationNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		identificationNumber := c.Param("identificationNumber") // Assuming the parameter is part of the URL path

		// Validate identificationNumber if needed

		// Find all contracts for the given vehicle identification number
		cursor, err := contractCollection.Find(ctx, bson.M{"vihecule.identificationnumber": identificationNumber})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer cursor.Close(ctx)

		var contracts []models.Contract
		if err := cursor.All(ctx, &contracts); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Check if any contracts were found
		if len(contracts) == 0 {
			c.JSON(http.StatusNotFound, responses.ContractResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No contracts found with the given identification number"}})
			return
		}

		c.JSON(http.StatusOK, responses.ContractResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"contracts": contracts}})
	}
}

// CheckContractValidityByViheculeIdentificationNumber checks the validity of a contract based on the Vihecule IdentificationNumber
func CheckContractValidityByViheculeIdentificationNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		identificationNumber := c.Param("identificationNumber") // Assuming the parameter is part of the URL path

		// Validate identificationNumber if needed

		// Helper function to check contract validity
		isValid, err := findValidContractByViheculeIdentificationNumber(c.Request.Context(), identificationNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Response with validity status
		condition := isValid != nil
		c.JSON(http.StatusOK, responses.ContractResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"validity": condition},
		})
	}
}

func GetValiditContractOfViheculeByIdentificationNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		identificationNumber := c.Param("identificationNumber") // Assuming the parameter is part of the URL path

		// Validate identificationNumber if needed

		// Helper function to check contract validity
		isValid, err := findValidContractByViheculeIdentificationNumber(c.Request.Context(), identificationNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ContractResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// Response with validity status
		if isValid == nil {
			c.JSON(http.StatusOK, responses.ContractResponse{
				Status:  http.StatusNotFound,
				Message: "success",
				Data:    map[string]interface{}{"error": "this vihecule dont have any valid contract"},
			})
			return
		}
		c.JSON(http.StatusOK, responses.ContractResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"contract": isValid},
		})
	}
}
