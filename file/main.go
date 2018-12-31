package main

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
	"log"
	"hash/crc32"
	"flag"
)

type FileInfo struct {
	Path string
	FileInfo os.FileInfo
}

var (
	dirA = "/Volumes/WD_1.0T_01"
	bloom = make(map[uint32][]string)
)

func main() {
	// var dir string
	dir := flag.String("dir", "/", "dirname")
	flag.Parse()

	chFile := make(chan FileInfo)
	go walkDir(chFile, *dir)
	for info := range chFile {
		func() {
			fp, _ := os.Open(info.Path)
			defer fp.Close()

			buffer := make([]byte, 4096)
			{
				if _, err := fp.Read(buffer[:len(buffer) / 2]); nil != err {
					// log.Printf("read fatal, err: %+v, path:%s", err, path)
				}
				fp.Seek(int64(info.FileInfo.Size() / 2), os.SEEK_SET)
				if _, err := fp.Read(buffer[len(buffer) / 2:]); nil != err {
					// log.Printf("read fatal, err: %+v, path:%s", err, path)
				}
			}
			c := crc32.ChecksumIEEE(buffer)
			bloom[c] = append(bloom[c], info.Path)
			if len(bloom[c]) > 1 {
				// fmt.Println(buffer)
				// log.Printf("same file, sum: %d, %+v", c, bloom[c])
				fmt.Printf("same file, sum: %d\n", c)
				for _, name := range bloom[c] {
					fmt.Println(name)
				}
			}
		}()
	}
}

func walkDir(ch chan FileInfo, dir string) {
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		// log.Printf("dir: %s, fileinfo: %+v", path, f)
		if nil != err {
			log.Fatalf("%+v", err)
			close(ch)
			// return err
		}
		if f.IsDir() {
			return nil
		}
		for _, instr := range []string{".jpg", ".jpeg", ".com.url"} {
			if strings.Contains(f.Name(), instr) {
				return nil
			}
		}
		ch <- FileInfo{Path: path, FileInfo: f}
		return nil
	})
}
