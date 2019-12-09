package main

import (
	"time"
)

//Job information related to a compute job
type Job struct {

	//Metadata
	ID          int       `json:"id"`
	Status      status    `json:"status"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Owner       string    `json:"owner"`
	SubmitDate  time.Time `json:"submitdate"`
	RunDate     time.Time `json:"rundate"`
	FinishDate  time.Time `json:"finishdate"`

	//Env specs
	//os osType
	//TODO
	//platform Platformtype
	//perfTier int
	//minRAM   int
	//hasGPU   bool
}

type status int

const (
	waiting status = iota
	running
	failed
	finished
)

type osType int

const (
	windows osType = iota
	linux   osType = iota
	other   osType = iota
)
