package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoPort = 27017

var client *mongo.Client
var database *mongo.Database
var collectionJobs *mongo.Collection
var collectionFiles *mongo.Collection
var collectionUsers *mongo.Collection

//Database names
const dbname = "cloudsourced"

//Collection names
const collectionNameJobs = "jobs"
const collectionNameFiles = "files"
const collectionNameUsers = "users"

func dbConnect() {

	clientOptions := options.Client().ApplyURI("mongodb://" + config.MongoUser + ":" + config.MongoPass + "@" + config.MongoHost + ":" + strconv.Itoa(mongoPort))
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	client = c

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//TODO do error checking
	database = client.Database(dbname)
	collectionJobs = database.Collection(collectionNameJobs)
	collectionFiles = database.Collection(collectionNameFiles)
	collectionUsers = database.Collection(collectionNameUsers)
}

func dbDisconnect() {

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connection closed")
}

func storeJob(job Job) {

	result, err := collectionJobs.InsertOne(context.TODO(), job)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func getJobByID(id int) Job {

	result := collectionFiles.FindOne(context.TODO(), bson.M{"ID": id})
	if result == nil {
		log.Fatal("No job matching ID " + string(id) + " was found")
	}

	var job Job
	err := result.Decode(&job)
	if err != nil {
		log.Fatal(err)
	}

	return job
}

func getJobsByAttribute(b map[string]interface{}) []Job {

	cursor, err := collectionJobs.Find(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}

	var jobs []Job

	for cursor.Next(context.TODO()) {
		fmt.Println("a thing")
		var j Job
		err := cursor.Decode(&j)
		if err != nil {
			log.Fatal(err)
		}
		jobs = append(jobs, j)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.TODO())

	return jobs
}

func getJobsByStatus(s status) []Job {

	return getJobsByAttribute(bson.M{"status": s})
}

func getJobsByOwner(owner string) []Job {

	return getJobsByAttribute(bson.M{"owner": owner})
}
