package cache_sdk

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_unzip(t *testing.T) {

	dir, err := ioutil.ReadDir("/Users/didi/go/src/awesomeProject/learn/types/")

	for _, fi := range dir {
		fmt.Println(fi.Name())
	}

	fmt.Println("-----------")
	fmt.Println(err)
	fmt.Println("-----------")
	//DeCompressTarGz("demo1.tar.gz", "Task")
}

// DeCompressTarGz 解压 tar.gz
func DeCompressTarGz(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		//logger.Logger.WithFields(logrus.Fields{"msg": "upzip tar", "tarFile": tarFile, "dest": dest, "date": hdr}).Info("_com_load_cache")
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

//解压 tar.gz
func DeCompress1(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + "/" + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}

//func Test_unzip(t *testing.T) {
//
//	// file read
//	fr, err := os.Open("./demo.tar.gz")
//	if err != nil {
//		panic(err)
//	}
//	defer fr.Close()
//
//	// gzip read
//	gr, err := gzip.NewReader(fr)
//	if err != nil {
//		panic(err)
//	}
//	defer gr.Close()
//
//	// tar read
//	tr := tar.NewReader(gr)
//
//	// 读取文件
//	for {
//		h, err := tr.Next()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			panic(err)
//		}
//
//		// 显示文件
//		fmt.Println(h.Name)
//
//		str, _ := os.Getwd()
//		fmt.Println(str)
//
//		// 打开文件
//		//panic: open /Users/didi/go/src/awesomeProject/learn/types/files/dir/: is a directory [recovered]
//		//panic: open /Users/didi/go/src/awesomeProject/learn/types/files/dir/: is a directory
//		//not a file
//
//		fw, err := os.OpenFile(h.Name, os.O_CREATE|os.O_WRONLY, 644 /*os.FileMode(h.Mode)*/)
//		if err != nil {
//			panic(err)
//		}
//		defer fw.Close()
//
//		// 写文件
//		_, err = io.Copy(fw, tr)
//		if err != nil {
//			panic(err)
//		}
//
//	}
//
//	fmt.Println("un tar.gz ok")
//}
