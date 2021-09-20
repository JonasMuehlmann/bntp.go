package libdocuments

import (
	"bufio"
	"log"
	"os"
)

func AddTag(documentPath string, tag string) {

}

func RemoveTag(documentPath string, tag string) {

}

func RenameTag(documentPath string, oldTag string, newTag string) {

}

func GetTags(documentPath string) []string {

}

func FindTagsLine(documentPath string) (int, string) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		if scanner.Text() == "# Tags" {
			return i, scanner.Text()
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
