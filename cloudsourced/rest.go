package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// func status() {
// 	//TODO
// }

// func listJobs() {
// 	//TODO
// }

func checkJSONErr(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

//Dashboard-specifc (mostly) APIs

func handleQueuedJobs(rw http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(getJobsByStatus(waiting))
	checkJSONErr(err)
	rw.Write(data)
}

func handleRunningJobs(rw http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(jobsRunning)
	checkJSONErr(err)
	rw.Write(data)
}

func handleFailedJobs(rw http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(getJobsByStatus(failed))
	checkJSONErr(err)
	rw.Write(data)
}

func handleJobsByOwner(rw http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var object map[string]interface{}
	json.Unmarshal(body, &object)
	var owner = object["owner"]

	data, err := json.Marshal(getJobsByOwner(owner.(string)))
	checkJSONErr(err)
	rw.Write(data)
}

func handleUploadFile(rw http.ResponseWriter, r *http.Request) {
	//TODO
}

func handleDeleteFile(rw http.ResponseWriter, r *http.Request) {
	//TODO
}

func handleSubmitJob(rw http.ResponseWriter, r *http.Request) {
	//TODO
}

func handleCancelJob(rw http.ResponseWriter, r *http.Request) {
	//TODO
}

//Worker-specific APIs

func handleConnect(rw http.ResponseWriter, r *http.Request) {
	//TODO
}

func handleDisconnect(rw http.ResponseWriter, r *http.Request) {
	//TODO
}

func handleGetWork(rw http.ResponseWriter, r *http.Request) {
	//TODO
}

func handleSubmitWork(rw http.ResponseWriter, r *http.Request) {
	//TODO
}

func handleGetFile(rw http.ResponseWriter, r *http.Request) {
	//TODO
}
