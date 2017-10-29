package main

import (
	"io"
	"os"
)

// setUpUserDirectory copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func setUpUserDirectory(uid string) {
	userDirLocation := PrjDir + uid
	ifErr(os.MkdirAll(userDirLocation, os.ModePerm))
	setDefaultProfilePic(uid)
}

func setDefaultProfilePic(uid string) {
	src := PrjDir + "profile.png"
	dst := PrjDir + uid + "/" + uid

	in, err := os.Open(src)
	ifErr(err)
	defer in.Close()

	out, err := os.Create(dst)
	ifErr(err)
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	// TODO: Error check that compares bytes written to original btyes of file
	_, err = io.Copy(out, in)
	ifErr(err)
	ifErr(out.Sync())
}
