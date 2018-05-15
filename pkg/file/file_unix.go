// +build !windows

package file

import (
	"os"
	"syscall"
)

func SyncDir(dirName string) error {
	// fsync the dir to flush the rename
	dir, err := os.OpenFile(dirName, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return err
	}
	defer dir.Close()

	// while we're on unix, we may be running in a docker container that is
	// pointed at a windows volume over samba. that doesn't support fsyncs
	// on directories. this shows itself as an EINVAL, so we ignore that
	// error.
	err = dir.Sync()
	if pe, ok := err.(*os.PathError); ok && pe.Err == syscall.EINVAL {
		err = nil
	} else if err != nil {
		return err
	}

	return dir.Close()
}

// RenameFile will rename the source to target using os function.
func RenameFile(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}
