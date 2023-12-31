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
## Stay in touch :
- Author - [Ouail Laamiri](https://www.linkedin.com/in/ouaillaamiri/) 
- Test - [Postman](https://www.postman.com/avionics-meteorologist-32935362/workspace/postman-api-fundamentals-student-expert/collection/29141176-d922c605-2315-488b-850b-e47edeccdaf1?action=share&creator=29141176)
- Documentation - [Postman](https://documenter.getpostman.com/view/29141176/2s9YsDkamW)

## License

Contract-Service is [MIT licensed](LICENSE).