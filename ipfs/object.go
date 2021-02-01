package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type IPFSLink struct {
	Name string
	// name or alias of this link
	Hash []byte
	// cryptographic hash of target
	Size int
	// total size of target
}
type IPFSObject struct {
	links []IPFSLink
	// array of links
	data []byte
	// opaque content data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// CreateIPFSTree converts normal directory into IPFS tree
func CreateIPFSTree(dirpath string) {
	filepath.Walk(dirpath, func(path string, f os.FileInfo, err error) error {
		switch mode := f.Mode(); {
		case mode.IsDir():
			DirToIPFSObject(path)
		case mode.isRegular():
			FileToIPFSObject(path)
		}

	})
}

func FileToIPFSObject(filepath string) {
	dat, err := ioutil.ReadFile(filepath)
	check(err)

	h := sha1.New()
	h.Write(dat)
	fileHash := h.Sum(nil)
	fmt.Printf("%x\n", fileHash)

	// links := IPFSLinks {

	// }

	// return IPFSObject {

	// }
	return
}

func DirToIPFSObject(dirpath string) {
	return
}
