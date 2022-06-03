// Copyright © 2021 Nicolas Wang <cqwang@uw.edu>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rclonefs

import (
	"sort"
	"syscall"

	"os"
	"time"

	_ "github.com/rclone/rclone/backend/all"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config"
	"github.com/rclone/rclone/fs/config/configfile"
	"github.com/rclone/rclone/vfs"

	"github.com/spf13/afero"
)

type Fs struct {
	*vfs.VFS
}

func New(client fs.Fs) afero.Fs {
	return &Fs{vfs.New(client, nil)}
}

func (s Fs) Name() string { return "RCloneFs" }

func (s Fs) Create(name string) (afero.File, error) {
	return s.VFS.Create(name)
}

func (s Fs) MkdirAll(path string, perm os.FileMode) error {
	// Fast path: if we can tell whether path is a directory or file, stop with success or error.
	dir, err := s.Stat(path)
	if err == nil {
		if dir.IsDir() {
			return nil
		}
		return nil
	}

	// Slow path: make sure parent exists and then call Mkdir for path.
	i := len(path)
	for i > 0 && os.IsPathSeparator(path[i-1]) { // Skip trailing path separator.
		i--
	}

	j := i
	for j > 0 && !os.IsPathSeparator(path[j-1]) { // Scan backward over element.
		j--
	}

	if j > 1 {
		// Create parent
		if err = s.MkdirAll(path[0:j-1], perm); err!=nil{
			return err
		}
	}

	// Parent now exists; invoke Mkdir and use its result.
	if err = s.Mkdir(path, perm); err != nil {
		// Handle arguments like "foo/." by
		// double-checking that directory doesn't exist.
		if dir, err1 := s.Stat(path); err1 == nil && dir.IsDir() {
			return nil
		}else{
			return err
		}
	}
	return nil
}

func (s Fs) Open(name string) (afero.File, error) {
	return s.VFS.Open(name)
}

func (s Fs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	return s.VFS.OpenFile(name, flag, perm)
}

func (s Fs) Remove(name string) error {
	return s.VFS.Remove(name)
}

func (s Fs) RemoveAll(path string) error {
	// Follow the same way as os.RemoveAll
	if path == "" {
		// fail silently to retain compatibility with previous behavior
		// of RemoveAll. See issue 28830.
		return nil
	}

	// The rmdir system call does not permit removing ".",
	// so we don't permit it either.
	if path == "." || (len(path) >= 2 && path[len(path)-1] == '.' && os.IsPathSeparator(path[len(path)-2])) {
		return &os.PathError{Op: "RemoveAll", Path: path, Err: syscall.EINVAL}
	}

	// Simple case: if Remove works, we're done.
	err := s.VFS.Remove(path)
	if err == nil || os.IsNotExist(err) {
		return nil
	}

	fileList := make([]string, 0)
	dirList := make([]string, 0)

	listWalkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			dirList = append(dirList, path)
		}else{
			fileList = append(fileList, path)
		}
		return nil
	}

	if err := afero.Walk(s, path, listWalkFn); err==nil{
		for _, f := range fileList{
			if e := s.Remove(f); e!=nil && !os.IsNotExist(e){
				return e
			}
		}

		sort.Slice(dirList, func(i, j int) bool {
			return len(dirList[i]) > len(dirList[j])
		})

		for _, d := range dirList{
			if e := s.Remove(d); e!=nil && !os.IsNotExist(e){
				return e
			}
		}
	}else {
		return err
	}

	return nil
}

func (s Fs) Rename(oldname, newname string) error {
	return s.VFS.Rename(oldname, newname)
}

func (s Fs) Stat(name string) (os.FileInfo, error) {
	return s.VFS.Stat(name)
}

func (s Fs) Chmod(name string, mode os.FileMode) error {
	// we do not support Chmod, silently return
	return nil
}

func (s Fs) Chown(name string, uid, gid int) error {
	// we do not support Chmod, silently return
	return nil
}

func (s Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return s.VFS.Chtimes(name, atime, mtime)
}

func SetConfigPath(path string) error {
	err := config.SetConfigPath(path)
	if err!=nil{
		return err
	}
	//config.ClearConfigPassword()
	configfile.Install()

	return nil
}
