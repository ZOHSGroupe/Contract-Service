## Description

Contract-Service is an Application Programming Interface (API) to handle add,update,delete and get contract of ZOHS company into NoSQL database.
## Installation :
```bash
# install requirements
$ go get -u github.com/gin-gonic/gin go.mongodb.org/mongo-driver/mongo github.com/joho/godotenv github.com/go-playground/validator/v10 github.com/klauspost/compress@v1.16.3 github.com/bytedance/sonic github.com/dgrijalva/jwt-go
``` 
## Running the app : 
```bash
# Run application
$ go run main.go
```
## Build Docker image : 
```bash
# build a docker image
$ docker build -t contract-service .
```
## Running the app in the Docker : 
```bash
# Run docker image
$ docker run -p 5000:5000 contract-service
```
## Running with Docker compose :
```bash
# Run docker compose
$ docker compose up
```

## Models

### Client

- `Firstname: string` - First name of the client.
- `Lastname: string` - Last name of the client.
- `Email: string` - Email address of the client.
- `City: string` - City of the client.
- `BirthDate: string` - Birth date of the client.
- `Gender: string` - Gender of the client.
- `Nationality: string` - Nationality of the client.
- `Address: string` - Address of the client.
- `NationalId: string` - National ID of the client.

### Vihecule

- `CurrentValue: string` (required) - Current value of the vehicle.
- `CylinderCount: string` (required) - Cylinder count of the vehicle.
- `EmptyWeight: string` (required) - Empty weight of the vehicle.
- `FuelType: string` (required) - Fuel type of the vehicle.
- `Genre: string` (required) - Genre of the vehicle.
- `GrossWeightRating: string` (required) - Gross weight rating of the vehicle.
- `ManufacturingDate: string` (required) - Manufacturing date of the vehicle.
- `Marque: string` (required) - Marque of the vehicle.
- `TaxHorsePower: string` (required) - Tax horse power of the vehicle.
- `Type: string` (required) - Type of the vehicle.
- `IdentificationNumber: string` (required) - Identification number of the vehicle.

### Permit

- `IssueDate: string` (required) - Issue date of the permit.
- `EndDate: string` (required) - End date of the permit.
- `LicenceNumber: string` (required) - License number of the permit.
- `Type: string` (required) - Type of the permit.

### Contract

- `StartDate: string` (required) - Start date of the contract.
- `EndDate: string` (required) - End date of the contract.
- `Value: string` (required) - Value of the contract.
- `Client: Client` (required) - Client information.
- `Vihecule: Vihecule` (required) - Vehicle information.
- `Permit: Permit` (required) - Permit information.

## Available Endpoints

### 1.CreateContract
- **Endpoint:** `POST /contract`
- **Description:** Create a new contract.
- **Request Body:**
  - `StartDate: string` (required) - Start date of the contract.
  - `EndDate: string` (required) - End date of the contract.
  - `Value: string` (required) - Value of the contract.
  - `Client: Client` (required) - Client information.
  - `Vihecule: Vihecule` (required) - Vehicle information.
  - `Permit: Permit` (required) - Permit information.
- **Response:**
  - `201`: Contract created successfully with the ID.
  - `400`: Bad Request (error in request or validation).
  - `500`: Internal Server Error.

### 2. GetAllContracts

- **Endpoint**: `GET /contract`
- **Description**: Get all contracts.
- **Response**:
  - `200`: Successful retrieval with a list of contracts.
  - `500`: Internal Server Error.

### 3.GetAContract

- **Endpoint:** `GET /contract/:id`
- **Description:** Get details of a contract by ID.
- **Response:**
  - `200`: Successful retrieval with contract details.
  - `400`: Bad Request (invalid ID format).
  - `404`: Contract not found.
  - `500`: Internal Server Error.

### 4.DeleteAContract

- **Endpoint:** `DELETE /contract/:id`
- **Description:** Delete a contract by ID.
- **Response:**
  - `200`: Contract deleted successfully.
  - `404`: Contract not found.
  - `500`: Internal Server Error.

### 5. GetContractsByNationalID

- **Endpoint**: `GET /contract/client/:nationalId`
- **Description**: Get contracts by client's national ID.
- **Response**:
  - `200`: Successful retrieval with a list of contracts.
  - `404`: No contracts found with the given national ID.
  - `500`: Internal Server Error.

### 6. GetContractsByVehicleIdentificationNumber

- **Endpoint**: `GET /contract/vihecule/:identificationNumber`
- **Description**: Get contracts by vehicle identification number.
- **Response**:
  - `200`: Successful retrieval with a list of contracts.
  - `404`: No contracts found with the given vehicle identification number.
  - `500`: Internal Server Error.

### 7. CheckContractValidityByViheculeIdentificationNumber

- **Endpoint**: `GET /contract/validity/:identificationNumber`
- **Description**: Check the validity of contracts by vehicle identification number.
- **Response**:
  - `200`: Successful retrieval with the validity status.
  - `500`: Internal

  ### 8. GetValiditContractOfViheculeByIdentificationNumber

- **Endpoint**: `GET /contract/valid/:identificationNumber`
- **Description**: Get a valid contract of vehicle by identification number.
- **Response**:
  - `200`: Successful retrieval with the contract.
  - `404`: No contracts valid found with the given vehicle identification number.
  - `500`: Internal


## Stay in touch :
- Author - [Ouail Laamiri](https://www.linkedin.com/in/ouaillaamiri/) 
- Test - [Postman](https://www.postman.com/avionics-meteorologist-32935362/workspace/postman-api-fundamentals-student-expert/collection/29141176-d922c605-2315-488b-850b-e47edeccdaf1?action=share&creator=29141176)
- Documentation - [Postman](https://documenter.getpostman.com/view/29141176/2s9YsDkamW)

## License

Contract-Service is [GPL licensed]().