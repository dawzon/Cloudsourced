package main

import (
	"time"
)

//Job information related to a compute job
type Job struct {

	//Server only
	ID         int           `json:"id"`
	Owner      string        `json:"owner"`
	Status     status        `json:"status"`
	Timeout    time.Duration `json:"timeout"`
	SubmitDate time.Time     `json:"submitdate"`
	RunDate    time.Time     `json:"rundate"`
	FinishDate time.Time     `json:"finishdate"`
	Output     string        `json:"output"`

	//Labels
	Name        string `json:"name"`
	Description string `json:"description"`

	//Env specs
	Platform platformType `json:"platform"`
	//MinRAM   int          `json:"minram"`
	//Files    []string     `json:"files"`
	Script string `json:"script"`
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
	web
)

func platformStringToInt(p string) platformType {
	//TODO web is the only implemented platform
	return web
}
