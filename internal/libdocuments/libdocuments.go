package libdocuments

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func AddTag(documentPath string, tag string) {
	lineNumber, tags := FindTagsLine(documentPath)

	if lineNumber == -1 {
		log.Fatal("Could not read Tag line of file")
	}

	tags += tag

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	offset, err := file.Seek(int64(lineNumber), io.SeekStart)

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteAt([]byte(tags), offset)

	if err != nil {
		log.Fatal(err)
	}
}

func RemoveTag(documentPath string, tag string) {
	lineNumber, tags := FindTagsLine(documentPath)

	if lineNumber == -1 {
		log.Fatal("Could not read Tag line of file")
	}

	tags = strings.Replace(tags, tag, "", -1)
	tags = strings.Replace(tags, ",,", ",", -1)

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	offset, err := file.Seek(int64(lineNumber), io.SeekStart)

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteAt([]byte(tags), offset)

	if err != nil {
		log.Fatal(err)
	}
}

func RenameTag(documentPath string, oldTag string, newTag string) {
	lineNumber, tags := FindTagsLine(documentPath)

	if lineNumber == -1 {
		log.Fatal("Could not read Tag line of file")
	}

	tags = strings.Replace(tags, oldTag, newTag, -1)

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	offset, err := file.Seek(int64(lineNumber), io.SeekStart)

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteAt([]byte(tags), offset)

	if err != nil {
		log.Fatal(err)
	}
}

func GetTags(documentPath string) []string {
	lineNumber, tags := FindTagsLine(documentPath)

	if lineNumber == -1 {
		log.Fatal("Could not read Tag line of file")
	}

	return strings.Split(tags, ",")
}

func FindTagsLine(documentPath string) (int, string) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		if scanner.Text() == "# Tags" {
			scanner.Scan()
			return i + 1, scanner.Text()
		}
		i++
	}

	return -1, ""
}
func HasTag(documentPath string, tag string) bool {
	tags := GetTags(documentPath)
	for _, tag_ := range tags {

		if tag_ == tag {
			return true
		}
	}
	return false

}
func FindDocumentsWithTags(rootDir string, tags []string) []string {
	filesWithTags := make([]string, 0, 100)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() {
			for _, tag := range tags {
				if HasTag(path, tag) {
					filesWithTags = append(filesWithTags, tag)
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return filesWithTags
}

func AddLink(documentPathSource string, documentPathDestination string) {

}

func ListLinks(documentPathLinkSource string) []string {

}

func ListBacklinks(documentPathLinkDestination string) []string {

}
