package main

//Job information related to a compute job
type Job struct {
	//Metadata
	id          int
	name        string
	description string
	//Env specs
	os OSType
	//platform Platformtype
	perfTier int
	minRAM   int
	hasGPU   bool
}
