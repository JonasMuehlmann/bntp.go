package libdocuments

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func AddTag(documentPath string, tag string) {

}

func RemoveTag(documentPath string, tag string) {

}

func RenameTag(documentPath string, oldTag string, newTag string) {
	lineNumber, tags := FindTagsLine(documentPath)

	if lineNumber == -1 {
		log.Fatal("Could not read Tag line of file")
	}

	tags = strings.Replace(tags, oldTag, newTag, -1)

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

func FindDocumentsWithTags(tags []string) []string {

}

func AddLink(documentPathSource string, documentPathDestination string) {

}

func ListLinks(documentPathLinkSource string) []string {

}

func ListBacklinks(documentPathLinkDestination string) []string {

}
