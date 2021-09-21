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

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0o644)
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

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0o644)
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

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0o644)
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
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
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

func HasTags(documentPath string, tags []string) bool {
	documentTags := GetTags(documentPath)

	for _, tag := range tags {
		for _, documentTag := range documentTags {
			if tag == documentTag {
				continue
			}
			return false
		}
	}
	return true
}

func FindDocumentsWithTags(rootDir string, tags []string) []string {
	filesWithTags := make([]string, 0, 100)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() {
			if !HasTags(path, tags) {
				return nil
			}
			filesWithTags = append(filesWithTags, path)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return filesWithTags
}

func FindLinksLines(documentPath string) (int, int, []string) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
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

	lineNumberFirstLink := -1
	lineNumberLastLink := -1
	links := make([]string, 0, 10)

	i := 0

	for scanner.Scan() {
		if scanner.Text() == "# Links" {
			lineNumberFirstLink = i + 1

			break
		}
		i++
	}

	for scanner.Scan() && strings.HasPrefix(scanner.Text(), "- ") {
		links[i-lineNumberFirstLink] = scanner.Text()
		i++
	}

	lineNumberLastLink = i

	return lineNumberFirstLink, lineNumberLastLink, links
}

func FindBacklLinksLines(documentPath string) (int, int, []string) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
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

	lineNumberFirstLink := -1
	lineNumberLastLink := -1
	links := make([]string, 0, 10)

	i := 0

	for scanner.Scan() {
		if scanner.Text() == "# Backlinks" {
			lineNumberFirstLink = i + 1

			break
		}
		i++
	}

	for scanner.Scan() && strings.HasPrefix(scanner.Text(), "- ") {
		links[i-lineNumberFirstLink] = scanner.Text()
		i++
	}

	lineNumberLastLink = i

	return lineNumberFirstLink, lineNumberLastLink, links
}

func AddLink(documentPathSource string, documentPathDestination string) {
	lineNumberFirstLink, lineNumberLastLink, links := FindLinksLines(documentPathSource)

	if lineNumberFirstLink == -1 || lineNumberLastLink == -1 {
		log.Fatal("Could not read Tag line of file")
	}

	links = append(links, documentPathDestination)

	file, err := os.OpenFile(documentPathSource, os.O_RDWR, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	offset, err := file.Seek(int64(lineNumberFirstLink), io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteAt([]byte(strings.Join(links, "\n")), offset)

	if err != nil {
		log.Fatal(err)
	}
}

func RemoveLink(documentPathSource string, documentPathDestination string) {
}
