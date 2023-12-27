package controllers

import (
	"Contract-Service/app/configs"
	"Contract-Service/app/models"
	"Contract-Service/app/responses"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var contractCollection *mongo.Collection = configs.GetCollection(configs.DB, "contract")
var validate = validator.New()

func CreateContract() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var contract models.Contract
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&contract); err != nil {
			c.JSON(http.StatusBadRequest, responses.ContractResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&contract); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ContractResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

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
