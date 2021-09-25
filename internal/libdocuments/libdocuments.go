package libdocuments

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func AddTag(documentPath string, tag string) error {
	lineNumber, tags, err := FindTagsLine(documentPath)
	if err != nil {
		return err
	}

	tags += tag

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}

	defer file.Close()

	offset, err := file.Seek(int64(lineNumber), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(tags), offset)

	if err != nil {
		return err
	}

	return nil
}

func RemoveTag(documentPath string, tag string) error {
	lineNumber, tags, err := FindTagsLine(documentPath)
	if err != nil {
		return err
	}

	tags = strings.Replace(tags, tag, "", -1)
	tags = strings.Replace(tags, ",,", ",", -1)

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}

	defer file.Close()

	offset, err := file.Seek(int64(lineNumber), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(tags), offset)
	if err != nil {
		return err
	}

	return nil
}

func RenameTag(documentPath string, oldTag string, newTag string) error {
	lineNumber, tags, err := FindTagsLine(documentPath)

	if err != nil {
		return err
	}

	tags = strings.Replace(tags, oldTag, newTag, -1)

	file, err := os.OpenFile(documentPath, os.O_RDWR, 0o644)

	if err != nil {
		return err
	}

	defer file.Close()

	offset, err := file.Seek(int64(lineNumber), io.SeekStart)

	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(tags), offset)

	if err != nil {
		return err
	}

	return nil
}

func GetTags(documentPath string) ([]string, error) {
	_, tags, err := FindTagsLine(documentPath)

	if err != nil {
		return nil, err
	}

	return strings.Split(tags, ","), nil
}

func FindTagsLine(documentPath string) (int, string, error) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)

	if err != nil {
		return 0, "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0

	for scanner.Scan() {
		if scanner.Text() == "# Tags" {
			scanner.Scan()

			return i + 1, scanner.Text(), nil
		}
		i++
	}

	return 0, "", errors.New("Could not find tags line")
}

func HasTags(documentPath string, tags []string) (bool, error) {
	documentTags, err := GetTags(documentPath)

	if err != nil {
		return false, err
	}

	for _, tag := range tags {
		for _, documentTag := range documentTags {
			if tag == documentTag {
				continue
			}

			return false, nil
		}
	}

	return true, nil
}

func FindDocumentsWithTags(rootDir string, tags []string) ([]string, error) {
	filesWithTags := make([]string, 0, 100)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			hasTags, err := HasTags(path, tags)

			if err != nil {
				return err
			}
			if !hasTags {
				return nil
			}
			filesWithTags = append(filesWithTags, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return filesWithTags, nil
}

func FindLinksLines(documentPath string) (int, int, []string, error) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
	if err != nil {
		return 0, 0, nil, err
	}

	defer file.Close()

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

	return lineNumberFirstLink, lineNumberLastLink, links, nil
}

func FindBacklinksLines(documentPath string) (int, int, []string, error) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
	if err != nil {
		return 0, 0, nil, err
	}

	defer file.Close()

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

	return lineNumberFirstLink, lineNumberLastLink, links, nil
}

func AddLink(documentPathSource string, documentPathDestination string) error {
	lineNumberFirstLink, lineNumberLastLink, links, err := FindLinksLines(documentPathSource)

	if err != nil {
		return err
	}

	links = append(links, documentPathDestination)

	file, err := os.OpenFile(documentPathSource, os.O_RDWR, 0o644)

	if err != nil {
		return err
	}

	defer file.Close()

	offset, err := file.Seek(int64(lineNumberFirstLink), io.SeekStart)

	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(strings.Join(links, "\n")), offset)

	if err != nil {
		return err
	}

	return nil
}

func AddBacklink(documentPathDestination string, documentPathSource string) error {
	lineNumberFirstLink, lineNumberLastLink, links, err := FindBacklinksLines(documentPathSource)

	if err != nil {
		return err
	}

	links = append(links, documentPathSource)

	file, err := os.OpenFile(documentPathDestination, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}

	defer file.Close()

	offset, err := file.Seek(int64(lineNumberFirstLink), io.SeekStart)

	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(strings.Join(links, "\n")), offset)

	if err != nil {
		return err
	}

	return nil
}

func RemoveBacklink(documentPathDestination string, documentPathSource string) error {
	lineNumberFirstLink, lineNumberLastLink, linksOrig, err := FindBacklinksLines(documentPathSource)

	if err != nil {
		return err
	}

	iLinkToDelete := -1

	for i, link := range linksOrig {
		if link == documentPathSource {
			iLinkToDelete = i
		}
	}

	links := make([]string, 0, 10)

	links = append(links, linksOrig[:iLinkToDelete]...)
	links = append(links, linksOrig[iLinkToDelete+1:]...)

	file, err := os.OpenFile(documentPathDestination, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	offset, err := file.Seek(int64(lineNumberFirstLink), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(strings.Join(links, "\n")), offset)

	if err != nil {
		return err
	}

	return err
}
