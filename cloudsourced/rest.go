package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	data, err := json.Marshal(getJobsByStatus(running))
	checkJSONErr(err)
	rw.Write(data)
}

// func handleFailedJobs(rw http.ResponseWriter, r *http.Request) {

// 	data, err := json.Marshal(getJobsByStatus(failed))
// 	checkJSONErr(err)
// 	rw.Write(data)
// }

func handleFinishedJobs(rw http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(getJobsByStatus(finished))
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

// func handleUploadFile(rw http.ResponseWriter, r *http.Request) {
// 	//TODO
// }

// func handleDeleteFile(rw http.ResponseWriter, r *http.Request) {
// 	//TODO
// }

func handleSubmitJob(rw http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var object map[string]interface{}
	json.Unmarshal(body, &object)
	var name = object["name"].(string)
	var description = object["description"].(string)
	var platform = object["description"].(string)
	var script = object["script"].(string)
	var owner = object["owner"].(string)
	var timeout = int(object["timeout"].(float64))

	newJob := Job{
		ID:          getNextID(),
		Owner:       owner,
		Status:      waiting,
		Timeout:     time.Duration(timeout) * time.Second,
		SubmitDate:  time.Now(),
		Name:        name,
		Description: description,
		Platform:    platformStringToInt(platform),
		Script:      script,
	}

	storeJob(newJob)
}

// func handleCancelJob(rw http.ResponseWriter, r *http.Request) {
// 	//TODO
// }

func handleActiveNodes(rw http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(nodes)
	checkJSONErr(err)
	rw.Write(data)
}

// func handleConnect(rw http.ResponseWriter, r *http.Request) {

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var object map[string]interface{}
// 	json.Unmarshal(body, &object)
// 	var nodeName = object["name"].(string)
// 	var nodeOwner = object["owner"].(string)

// 	addNode(Node{nodeName, nodeOwner})
// }

// func handleDisconnect(rw http.ResponseWriter, r *http.Request) {
// 	//TODO
// }

func handleGetWork(rw http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var object map[string]interface{}
	json.Unmarshal(body, &object)
	var alias = object["alias"].(string)

	addNode(Node{alias}) //TODO this might not be the right way to do this

	job, available := getJobAndRun(alias) //This marks the job as running in the database
	if !available {
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	jobJSON, _ := json.Marshal(job)
	rw.Write([]byte(jobJSON))

	jobTimeouts[job.ID] = time.AfterFunc(job.Timeout, func() {
		resetAbandonedJob(job.ID)
	})
}

func handleSubmitWork(rw http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var object map[string]interface{}
	json.Unmarshal(body, &object)
	var id = int(object["id"].(float64))
	var alias = object["alias"].(string)
	var output = object["output"].(string)

	addNode(Node{alias}) //TODO

	job := getJobByID(id)
	if job.Status != running {
		//TODO return an error
		return
	}

	jobTimeouts[id].Stop()
	delete(jobTimeouts, id)

	finalizeJob(id, output)
}

// func handleGetFile(rw http.ResponseWriter, r *http.Request) {
// 	//TODO
// }
