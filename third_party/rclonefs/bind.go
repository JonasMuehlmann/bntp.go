package rclonefs

import (
	"context"
	"fmt"
	"github.com/rclone/rclone/fs"
	"github.com/spf13/afero"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"
)

func NewRCloneFs(section string) afero.Fs {
	if newFs, err := fs.NewFs(context.Background(), section+":"); err == nil {
		return New(newFs)
	} else {
		panic(err)
	}
}

// BindPathFs works very much like Linux mount bind
// bindPointLayer is read only filesystem which only contain bind directory hierarchical
type BindPathFs struct {
	bindPointLayer afero.Fs
	pathPrefix []string
	bindFs []afero.Fs
}

type BindPathFile struct {
	afero.File
	path string
}

func (f *BindPathFile) Name() string {
	return f.path
}

// NewBindPathFile return a file with a modified file path if necessary
func NewBindPathFile(f afero.File, path string) afero.File {
	if filepath.Clean(path)==filepath.Clean(f.Name()){
		return f
	}else{
		return &BindPathFile{File: f, path: filepath.Clean(path)}
	}
}

// NewBindPathFs serve the same functionality like Linux `/etc/fstab` by bind-mounting different afero filesystem into single one
// Like `afero.BasePathFs` in reverse way, it remove a mount point prefix to access the underlining file
func NewBindPathFs(binds map[string]afero.Fs) afero.Fs {
	pathPrefix := make([]string, 0, len(binds))
	bindMap := make(map[string]string)

	for k := range binds{
		cleanPath := filepath.Clean(k)
		bindMap[cleanPath] = k
		if len(bindMap) == len(pathPrefix){
			panic(fmt.Errorf("path %s is a duplicate of existing %s", k, cleanPath))
		}

		pathPrefix = append(pathPrefix, cleanPath)
	}

	sort.Slice(pathPrefix, func(i, j int) bool {
		return len(pathPrefix[i]) > len(pathPrefix[j])
	})

	bindFs := make([]afero.Fs, 0, len(pathPrefix))

	rootLayer := afero.NewMemMapFs()

	for _, k := range pathPrefix{
		bindFs = append(bindFs, binds[bindMap[k]])
		_ = rootLayer.MkdirAll(k, os.ModeDir)
	}

	return &BindPathFs{
		bindPointLayer: afero.NewReadOnlyFs(rootLayer),
		pathPrefix: pathPrefix,
		bindFs: bindFs,
	}
}

func (b *BindPathFs) realPath(name string) (fsIndex int, path string, err error) {
	cleanPath := filepath.Clean(name)
	for i, prefix := range b.pathPrefix{
		if strings.HasPrefix(cleanPath, prefix){
			newPath := cleanPath[len(prefix):]
			if prefix=="/"{
				newPath = cleanPath
			}
			if newPath==""{
				newPath = "/"
			}
			return i, newPath, nil
		}
	}

	return len(b.pathPrefix), "", os.ErrNotExist
}

func (b *BindPathFs) Chtimes(name string, atime, mtime time.Time) (err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return &os.PathError{Op: "chtimes", Path: name, Err: err}
	}else{
		// we may lost the exact filename if error return
		return b.bindFs[bindFs].Chtimes(n, atime, mtime)
	}
}

func (b *BindPathFs) Chmod(name string, mode os.FileMode) (err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return &os.PathError{Op: "chmod", Path: name, Err: err}
	}else{
		return b.bindFs[bindFs].Chmod(n, mode)
	}
}

func (b *BindPathFs) Chown(name string, uid, gid int) (err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return &os.PathError{Op: "chown", Path: name, Err: err}
	}else{
		return b.bindFs[bindFs].Chown(n, uid, gid)
	}
}

func (b *BindPathFs) Name() string {
	return "BindPathFs"
}

func (b *BindPathFs) Stat(name string) (fi os.FileInfo, err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return nil, &os.PathError{Op: "stat", Path: name, Err: err}
	}else{
		return b.bindFs[bindFs].Stat(n)
	}
}

func (b *BindPathFs) Rename(oldName, newName string) (err error) {
	if bindFsOld, oldPath, err := b.realPath(oldName); err != nil {
		return &os.PathError{Op: "rename", Path: oldName, Err: err}
	}else if bindFsNew, newPath, err := b.realPath(newName); err != nil {
		return &os.PathError{Op: "rename", Path: newName, Err: err}
	}else if bindFsOld==bindFsNew{
		return b.bindFs[bindFsOld].Rename(oldPath, newPath)
	}else{
		bindFsOld := b.bindFs[bindFsOld]
		bindFsNew := b.bindFs[bindFsNew]

		oldFile, err := bindFsOld.Open(oldPath)
		if err != nil {
			return &os.PathError{Op: "rename", Path: oldName, Err: err}
		}
		defer oldFile.Close()

		newFile, err := bindFsNew.Create(newPath)
		if err != nil {
			return &os.PathError{Op: "rename", Path: newName, Err: err}
		}
		defer newFile.Close()

		n, err := io.Copy(newFile, oldFile)
		if err != nil {
			_ = bindFsNew.Remove(newPath)
			return &os.PathError{Op: "rename", Path: newName, Err: err}
		}

		oldFileMeta, err := oldFile.Stat()
		if err != nil || oldFileMeta.Size() != n {
			_ = bindFsNew.Remove(newPath)
			return syscall.EIO
		}

		err = bindFsOld.Remove(oldPath)
		if err != nil {
			return &os.PathError{Op: "rename", Path: oldName, Err: err}
		}

		return bindFsNew.Chtimes(newPath, time.Now(), oldFileMeta.ModTime())
	}
}

func (b *BindPathFs) RemoveAll(name string) (err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return &os.PathError{Op: "remove_all", Path: name, Err: err}
	}else{
		return b.bindFs[bindFs].RemoveAll(n)
	}
}

func (b *BindPathFs) Remove(name string) (err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return &os.PathError{Op: "remove", Path: name, Err: err}
	}else{
		return b.bindFs[bindFs].Remove(n)
	}
}

func (b *BindPathFs) OpenFile(name string, flag int, mode os.FileMode) (f afero.File, err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return nil, &os.PathError{Op: "openfile", Path: name, Err: err}
	}else{
		file, err := b.bindFs[bindFs].OpenFile(n, flag, mode)
		if err != nil {
			return nil, err
		}
		return NewBindPathFile(file, name), nil
	}
}

func (b *BindPathFs) Open(name string) (f afero.File, err error) {
	var unionFile afero.File = nil
	if ok, err := afero.Exists(b.bindPointLayer, name); err==nil && ok{
		unionFile, err = b.bindPointLayer.Open(name)
	}
	if bindFs, n, err := b.realPath(name); err != nil {
		if unionFile!=nil{
			return unionFile, nil
		}else{
			return nil, &os.PathError{Op: "open", Path: name, Err: err}
		}
	}else{
		file, err := b.bindFs[bindFs].Open(n)
		if err != nil {
			return nil, err
		}
		if unionFile!=nil{
			// todo: we need a merger for UnionFile to handle Layer file overwrite Base directory case
			return &afero.UnionFile{
				Base: unionFile,
				Layer: NewBindPathFile(file, name),
			}, nil
		}else{
			return NewBindPathFile(file, name), nil
		}
	}
}

func (b *BindPathFs) Mkdir(name string, mode os.FileMode) (err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return &os.PathError{Op: "mkdir", Path: name, Err: err}
	}else{
		return b.bindFs[bindFs].Mkdir(n, mode)
	}
}

func (b *BindPathFs) MkdirAll(name string, mode os.FileMode) (err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return &os.PathError{Op: "mkdir", Path: name, Err: err}
	}else{
		return b.bindFs[bindFs].MkdirAll(n, mode)
	}
}

func (b *BindPathFs) Create(name string) (f afero.File, err error) {
	if bindFs, n, err := b.realPath(name); err != nil {
		return nil, &os.PathError{Op: "create", Path: name, Err: err}
	}else{
		file, err := b.bindFs[bindFs].Create(n)
		if err != nil {
			return nil, err
		}
		return NewBindPathFile(file, name), nil
	}
}