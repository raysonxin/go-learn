package main

type FileInfoRequest struct {
	FileName string
}

type FileRequest struct {
	FileName string
	Position int64
	Size     int64
	Datas    []byte
}
