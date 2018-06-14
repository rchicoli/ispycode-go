package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	filename := "/tmp/myfile.txt"

	fileinfo, _ := os.Stat(filename)

	stat := fileinfo.Sys().(*syscall.Stat_t)

	// atime is the file access time.
	fmt.Println(time.Unix(stat.Atim.Sec, stat.Atim.Nsec))

	// mtime is the file modify time.
	fmt.Println(time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec))

	// ctime is the inode or file change time.
	fmt.Println(time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec))
}
