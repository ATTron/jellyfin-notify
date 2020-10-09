package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var userKey = ""
var apiKey = ""
var endpoint = ""
var waitTime int

type Media struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

var lastItems []Media

func main() {
	epFlag := flag.String("endpoint", "https://localhost:8090", "the URL where your jellyfin instance lives")
	ukFlag := flag.String("user-key", "", "the user key for a user who has full library access")
	akFlag := flag.String("api-key", "", "the api key used to make the request")
	awsFlag := flag.String("aws-region", "us-east-1", "the aws region where your credentials have access to sns")
	snsFlag := flag.String("sns-arn", "", "the sns topic you will be publishing out to")
	envFlag := flag.Bool("env-file", false, "set if you are using an env file instead")
	timeFlag := flag.Int("wait-time", 168, "total time to wait before for new content")
	flag.Parse()

	if *envFlag {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		endpoint = os.Getenv("ENDPOINT")
		userKey = os.Getenv("USER-KEY")
		apiKey = os.Getenv("API-KEY")
		awsRegion = os.Getenv("AWS-REGION")
		snsArn = os.Getenv("SNS-ARN")
		waitTime, _ = strconv.Atoi(os.Getenv("WAIT-TIME"))
	} else {
		endpoint = *epFlag
		userKey = *ukFlag
		apiKey = *akFlag
		awsRegion = *awsFlag
		snsArn = *snsFlag
		waitTime = *timeFlag
	}

	checkVariables()

	var latestItems []Media
	for {
		resp, err := http.Get(join(endpoint, "emby/Users/", userKey, "/Items/Latest?IncludeItemTypes=Movie,Series&api_key=", apiKey))

		if err != nil {
			log.Fatalf("Failed to query the provided Jellyfin server at %s . . . exiting execution\n", endpoint)
			os.Exit(1)
		}

		respData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		jsonErr := json.Unmarshal(respData, &latestItems)
		if jsonErr != nil {
			log.Fatalln("failed to unmarshal json . . . exiting program")
		}

		if equal(lastItems, latestItems) {
			log.Printf("Nothing new has been added this week . . .\nSkipping until next week %+v", time.Now().Local().Add(time.Hour*time.Duration(168)))
		} else {
			log.Println("Not the same contents")
			lastItems = latestItems
			sendSMS(join("New movies and shows have been added to the jellyfin! See them at ", endpoint))
		}
		time.Sleep(time.Duration(waitTime) * time.Hour)
	}
}
