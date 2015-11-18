package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"io/ioutil"
	"net"
	"swk/socket/tcp"
)

type FileManage struct {
	Package *tcp.PackageData
	Conn    net.Conn
}

func (f FileManage) Handle() {
	switch f.Package.Command {
	case 0:
		f.DownloadFile()
		break
	case 1:
		f.UploadFile()
		break
	case 2:
		f.GetFileInfo()
		break
	case 3:
		f.DeleteFile()
		break
	}
}

func (f FileManage) DeleteFile() {
	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: url, User: user})
	if err != nil {
		fmt.Println(err)
		f.notifyError(err.Error())
	}
	isSuccess, err := fs.Delete(gowfs.Path{Name: path + string(f.Package.Datas)}, true)
	if err != nil {
		f.notifyError(err.Error())
		fmt.Println(err)
		return
	}
	result := tcp.ReplyPackage(nil, true, 0, 0, f.Package.Identify, isSuccess)
	f.Conn.Write(result)
}

func (f FileManage) DownloadFile() {

	fileRequest := &FileRequest{}
	json.Unmarshal(f.Package.Datas, &fileRequest)
	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: url, User: user})
	if err != nil {
		f.notifyError(err.Error())
		fmt.Println(err)
		return
	}

	_, err = fs.GetFileStatus(gowfs.Path{Name: path + fileRequest.FileName})
	if err != nil {
		f.notifyError(err.Error())
		fmt.Println(err)
		return
	}

	body, err := fs.Open(gowfs.Path{Name: path + fileRequest.FileName}, fileRequest.Position, fileRequest.Size, 2048)
	if err != nil {
		f.notifyError(err.Error())
		fmt.Println(err)
       
	} else {
		rcvdData, _ := ioutil.ReadAll(body)
		result := tcp.ReplyPackage(rcvdData, true, 0, 0, f.Package.Identify, true)
		f.Conn.Write(result)
	}
    fmt.Println("下载完成")
}

func (f FileManage) notifyError(err string) {
	result := tcp.ReplyPackage(err, true, 0, 0, f.Package.Identify, false)
	f.Conn.Write(result)
}

func (f FileManage) UploadFile() {
	fmt.Println("上传文件请求")
	fileRequest := &FileRequest{}
	json.Unmarshal(f.Package.Datas, &fileRequest)
	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: url, User: user})
	if err != nil {
		fmt.Println(err)
		f.notifyError(err.Error())
		return
	}
	buffer := bytes.NewBuffer(fileRequest.Datas)
	if fileRequest.Position == 0 {
		ok, err := fs.Delete(gowfs.Path{Name: path + fileRequest.FileName}, true)

		if err != nil {
			fmt.Println(err)
			f.notifyError(err.Error())
			return
		}
		ok, err = fs.Create(buffer, gowfs.Path{Name: path + fileRequest.FileName}, false, 0, 0, 0777, 0)
		if err != nil {
			fmt.Println(err)
			f.notifyError(err.Error())
			return
		}
		if ok {
			result := tcp.ReplyPackage(nil, true, 0, 0, f.Package.Identify, true)
			f.Conn.Write(result)
		} else {
			f.notifyError(err.Error())
			fmt.Println(err)
		}
	} else {
		ok, err := fs.Append(buffer, gowfs.Path{Name: path + fileRequest.FileName}, buffer.Len())
		if err != nil {
			fmt.Println(err)
			f.notifyError(err.Error())
		}
		if ok {
			result := tcp.ReplyPackage(nil, true, 0, 0, f.Package.Identify, true)
			f.Conn.Write(result)
		} else {
			f.notifyError(err.Error())
		}
	}
fmt.Println("上传完成")
}

func (f FileManage) GetFileInfo() {
	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: url, User: user})
	if err != nil {
		f.notifyError(err.Error())
		fmt.Println(err)
		return
	}
    fileName:=string(f.Package.Datas)
	datas, err := fs.GetFileStatus(gowfs.Path{Name: path+fileName})
	if err != nil {
		f.notifyError(err.Error())
		fmt.Println(err)
		return
	}
    result := tcp.ReplyPackage(datas, true, 0, 0, f.Package.Identify, true)
    f.Conn.Write(result)
}
