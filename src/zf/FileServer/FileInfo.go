package main

type FileInfo struct {
	Type             string
	Length           int
	AccessTime       int
	ModificationTime int
}

type FileStatus struct {
	FileStatus FileInfo
}
