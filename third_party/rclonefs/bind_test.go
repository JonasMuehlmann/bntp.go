package rclonefs

import (
	"fmt"
	"github.com/spf13/afero"
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func walkPrintFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	var size int64
	if !info.IsDir() {
		size = info.Size()
	}
	fmt.Println(path, info.Name(), size, info.IsDir(), err)
	return nil
}

func TestNewRCloneFs(t *testing.T) {
	usr, _ := user.Current()
	_ = SetConfigPath(filepath.Join(usr.HomeDir, ".config/rclone/rclone.conf"))
	fs := NewRCloneFs("ibm")

	_ = afero.Walk(fs, "/", walkPrintFn)
}

func TestNewBindPathFs(t *testing.T) {
	fsRoot := afero.NewMemMapFs()
	fsHome := afero.NewMemMapFs()

	_ = fsRoot.MkdirAll("/tmp/test", os.ModeDir)
	_ = afero.WriteFile(fsRoot, "/tmp/test/hello.go", []byte("fsRoot"), os.ModePerm)

	bindFs := NewBindPathFs(map[string]afero.Fs{
		"/":          fsRoot,
		"/home":      fsHome,
		"/home/root": fsRoot,
	})

	_ = afero.WriteFile(bindFs, "/home/world.go", []byte("bindFs to fsHome"), os.ModePerm)
	content, _ := afero.ReadFile(fsHome, "/world.go")
	t.Logf("write test: '%s' == '%s' ", []byte("bindFs to fsHome"), content)

	t.Logf("list all files inside bind filesystem")
	_ = afero.Walk(bindFs, "/", walkPrintFn)

	oldContent, _ := afero.ReadFile(bindFs, "/home/root/tmp/test/hello.go")
	t.Logf("load test: '%s' == '%s' ", []byte("fsRoot"), oldContent)

	_ = afero.WriteFile(bindFs, "/home/root/tmp/test/hello.go", []byte("bindFs /home/root"), os.ModePerm)
	newContent, _ := afero.ReadFile(bindFs, "/tmp/test/hello.go")
	t.Logf("write test: '%s' == '%s' ", []byte("bindFs /home/root"), newContent)

	t.Logf("list all files inside root filesystem")
	_ = afero.Walk(fsRoot, "/", walkPrintFn)

	t.Logf("list all files inside home filesystem")
	_ = afero.Walk(fsHome, "/", walkPrintFn)

	osFs := afero.NewOsFs()
	bindOSFs := NewBindPathFs(map[string]afero.Fs{
		"/":     osFs,
		"/home": osFs,
	})

	if tmpFile, err := bindOSFs.Open("/tmp/xx"); err == nil {
		defer tmpFile.Close()
		t.Logf("root bind file name %s", tmpFile.Name())

		if osTmpFile, ok := tmpFile.(*os.File); ok {
			t.Logf("success with real os file name %s", osTmpFile.Name())
		}
	}

	if tmpFile, err := bindOSFs.Open("/home/tmp/xx"); err == nil {
		defer tmpFile.Close()
		t.Logf("bindOSFs file name %s", tmpFile.Name())

		if osTmpFile, ok := tmpFile.(*os.File); ok {
			t.Errorf("should not happen %s", osTmpFile.Name())
		}
	}
}
