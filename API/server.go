package main

import (
	getcollection "API/Collection"
	database "API/databases"
	model "API/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func handleRequest(i context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var DB = database.ConnectDB()
	var recruiterCollection = getcollection.GetCollection(DB, "Recruiter")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var recruiter model.Recruiter

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	fmt.Println("upto here")

	var ijson map[string]interface{}

	json.Unmarshal([]byte(request.Body), &ijson)
	recruiter = model.Recruiter{}
	if FirstName, ok := ijson["FirstName"].(string); ok {
		recruiter.FirstName = FirstName
	}

	result, err := recruiterCollection.InsertOne(ctx, recruiter)
	fmt.Println("upto there")
	fmt.Println(result)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "failed!!", StatusCode: 200}, nil
	}

	data := map[string]interface{}{"data": result}
	jsonStr, _ := json.Marshal(data)

	return events.APIGatewayProxyResponse{Body: string(jsonStr), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
