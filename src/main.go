package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
var sess, err = session.NewSession(&aws.Config{
	Region: aws.String("us-east-1"),
})
var db = dynamodb.New(sess)

func init() {
	if err != nil {
		fmt.Println("Error creating DynamoDB session:", err)
		return
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	godotenv.Load(fmt.Sprintf("%s.env", env))
}

func generateShortURL() string {
	const caracteres = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = caracteres[rng.Intn(len(caracteres))]
	}
	return string(shortURL)
}

func createShortURL(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	longURL := request.QueryStringParameters["url"]
	if longURL == "" {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Missing URL parameter"}, nil
	}
	shortURL := generateShortURL()
	item := map[string]*dynamodb.AttributeValue{
		"shortUrl": {
			S: aws.String(shortURL),
		},
		"longUrl": {
			S: aws.String(longURL),
		},
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String("UrlShortener"),
		Item:      item,
	}
	_, err := db.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: "Internal Server Error - Cant Write"}, err
	}
	baseURL := os.Getenv("BASE_URL")
	fmt.Println("baseURL:", baseURL)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("Short URL: %s/%s\n", baseURL, shortURL),
	}, nil
}

func resolveURL(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	shortURL := request.PathParameters["shortURL"]

	input := &dynamodb.GetItemInput{
		TableName: aws.String("UrlShortener"),
		Key: map[string]*dynamodb.AttributeValue{
			"shortUrl": {
				S: aws.String(shortURL),
			},
		},
	}

	result, err := db.GetItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: "Internal Server Error - Cant Resolve"}, err
	}

	if result.Item == nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "Short URL not found to resolve"}, nil
	}

	longURL := *result.Item["longUrl"].S

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusSeeOther,
		Headers:    map[string]string{"Location": longURL},
	}, nil
}

func deleteURL(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	shortURL := request.QueryStringParameters["shortURL"]
	if shortURL == "" {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Missing shortURL parameter"}, nil
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("UrlShortener"),
		Key: map[string]*dynamodb.AttributeValue{
			"shortUrl": {
				S: aws.String(shortURL),
			},
		},
	}

	_, err := db.DeleteItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: "Internal Server Error - Cant delete"}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("Short URL %s deleted\n", shortURL),
	}, nil
}

func main() {
	lambda.Start(handler)
}
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case http.MethodPost:
		if request.Resource == "/create" {
			return createShortURL(request)
		}
	case http.MethodGet:
		if request.Resource == "/resolve/{shortURL}" {
			return resolveURL(request)
		}
	case http.MethodDelete:
		if request.Resource == "/delete/" {
			return deleteURL(request)
		}
	default:
		return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed, Body: "Method not allowed"}, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "Resource not found"}, nil
}
