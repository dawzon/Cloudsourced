package main

import (
	"net/http"
)

//Database names
const dbname = "cloudsourced"

//Collection names
const collectionNameJobs = "jobs"

func main() {

	loadConfig()

	// clientOptions := options.Client().ApplyURI("mongodb://" + config.MongoUser + ":" + config.MongoPass + "@" + config.MongoHost + ":27017")
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Oh yeah yeah")
	// j := Job{0, "My First Job", "Hello world!", windows}
	// jobsCollection := client.Database(dbname).Collection(collectionNameJobs)

	// insertResult, err := jobsCollection.InsertOne(context.TODO(), j)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("result: ", insertResult.InsertedID)

	// err = client.Disconnect(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connection closed.")
}

func init() {

	//TODO Switch to HTTPS!
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	http.ListenAndServe(":80", nil)

	http.HandleFunc("/test", test)
}

func test(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("what hath God wrought!"))
}
