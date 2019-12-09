package main

import (
	"time"
)

//Job information related to a compute job
type Job struct {

	//Server only
	ID         int       `json:"id"`
	Owner      string    `json:"owner"`
	Status     status    `json:"status"`
	SubmitDate time.Time `json:"submitdate"`
	RunDate    time.Time `json:"rundate"`
	FinishDate time.Time `json:"finishdate"`

	//Labels
	Name        string `json:"name"`
	Description string `json:"description"`

	//Env specs
	Platform platformType `json:"platform"`
	MinRAM   int          `json:"minram"`
	Files    []string     `json:"files"`
	script   string       `json:"files"`
}

type status int

const (
	waiting status = iota
	running
	failed
	finished
)

type platformType int

const (
	windows platformType = iota
	linux
	other
)
