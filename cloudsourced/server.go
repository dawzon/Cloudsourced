package main

import (
	"net/http"
	"sync"
	"time"
)

//REST endpoints
const apiEndpoint = "/api"

//Init This starts the server
func Init() {

	loadConfig()
	dbConnect()

	http.HandleFunc(apiEndpoint+"/queued_jobs", handleQueuedJobs)
	http.HandleFunc(apiEndpoint+"/running_jobs", handleRunningJobs)
	http.HandleFunc(apiEndpoint+"/failed_jobs", handleFailedJobs)
	http.HandleFunc(apiEndpoint+"/finished_jobs", handleFinishedJobs)
	http.HandleFunc(apiEndpoint+"/jobs_by_owner", handleJobsByOwner)
	//http.HandleFunc(apiEndpoint+"/upload_file", handleUploadFile)
	//http.HandleFunc(apiEndpoint+"/delete_file", handleDeleteFile)
	http.HandleFunc(apiEndpoint+"/submit_job", handleSubmitJob)
	//http.HandleFunc(apiEndpoint+"/cancel_job", handleCancelJob)
	http.HandleFunc(apiEndpoint+"/active_nodes", handleActiveNodes)

	//http.HandleFunc(apiEndpoint+"/connect", handleConnect)
	//http.HandleFunc(apiEndpoint+"/disconnect", handleDisconnect)
	http.HandleFunc(apiEndpoint+"/get_work", handleGetWork)
	http.HandleFunc(apiEndpoint+"/submit_work", handleSubmitWork)

	//TODO Switch to HTTPS!
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)

	http.Handle("/", http.FileServer(http.Dir("../dashboard/public")))
	http.ListenAndServe(":8080", nil)
	//the above function blocks, don't put anything here
}

var nextID int = 0
var idMutex sync.Mutex

func getNextID() int {
	idMutex.Lock()
	n := nextID
	nextID++
	defer idMutex.Unlock()
	return n
}

var nodes []Node
var nodeMutex sync.Mutex

func addNode(new Node) {

	nodeMutex.Lock()
	defer nodeMutex.Unlock()
	for _, n := range nodes {
		if new.alias == n.alias {
			//No need to add, alias is already in the system
			return
		}
	}
	nodes = append(nodes, new)
}

var jobTimeouts = make(map[int]*time.Timer)

func resetAbandonedJob(id int) {

	delete(jobTimeouts, id)
	setJobStatus(id, waiting)
}
