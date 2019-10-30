package main

//OSType operating system type
type OSType int

const (
	windows OSType = iota
	linux   OSType = iota
	other   OSType = iota
)
