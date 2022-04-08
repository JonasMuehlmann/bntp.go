// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the"Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED"AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package subcommands_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/cmd/cli/subcommands"
	"github.com/JonasMuehlmann/bntp.go/internal/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

// ******************************************************************//
//                             --add-tag                            //
// ******************************************************************//.
func TestAddTagToDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag := "foo::bar::baz"
	document := path.Join(test.TestDataTempDir, t.Name())
	docType := "bar"

	os.Args = []string{"", "document", "--add-tag", document, tag}

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	_, err = file.WriteString("# Tags")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, tag)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document, docType)
	assert.NoError(t, err)

	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --remove-tag                           //
// ******************************************************************//.
func TestRemoveTagFromDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag := "foo::bar::baz"
	document := path.Join(test.TestDataTempDir, t.Name())
	docType := "bar"

	os.Args = []string{"", "document", "--remove-tag", document, tag}

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	_, err = file.WriteString("# Tags\n" + tag)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, tag)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document, docType)
	assert.NoError(t, err)

	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --rename-tag                           //
// ******************************************************************//.
func TestRenameTagInDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	oldTag := "foo::bar::baz"
	newTag := "foo::bar"
	document := path.Join(test.TestDataTempDir, t.Name())
	docType := "bar"

	os.Args = []string{"", "document", "--rename-tag", document, oldTag, newTag}

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	_, err = file.WriteString("# Tags\n" + oldTag)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, oldTag)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document, docType)
	assert.NoError(t, err)

	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                            --get-tags                            //
// ******************************************************************//.
func TestGetTagsFromDocument(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag1 := "foo::bar::baz"
	tag2 := "x::y::z"
	document := path.Join(test.TestDataTempDir, t.Name())
	docType := "bar"

	os.Args = []string{"", "document", "--get-tags", document}

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	_, err = file.WriteString(fmt.Sprintf("# Tags\n%v,%v", tag1, tag2))
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document, docType)
	assert.NoError(t, err)

	err = subcommands.DocumentMain(db)
	stdOutInterceptBuffer.Scan()
	assert.Equal(t, tag1, stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, tag2, stdOutInterceptBuffer.Text())

	assert.NoError(t, err)
}

// ******************************************************************//
//                         --find-tags-line                         //
// ******************************************************************//.
func TestFindTagsLine(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	document := path.Join(test.TestDataTempDir, t.Name())
	docType := "bar"

	os.Args = []string{"", "document", "--find-tags-line", document}

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	tag := "foo::bar"
	_, err = file.WriteString("\n# Tags\n" + tag)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document, docType)
	assert.NoError(t, err)

	err = subcommands.DocumentMain(db)
	stdOutInterceptBuffer.Scan()

	assert.Equal(t, "2 "+tag, stdOutInterceptBuffer.Text())
	assert.NoError(t, err)
}

// ******************************************************************//
//                            --has-tags                            //
// ******************************************************************//.
func TestHasTags(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	document := path.Join(test.TestDataTempDir, t.Name())
	docType := "bar"

	tag1 := "foo::bar::baz"
	tag2 := "foo::bar"
	os.Args = []string{"", "document", "--has-tags", document, tag1, tag2}

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	_, err = file.WriteString(fmt.Sprintf("# Tags\n%v,%v", tag1, tag2))
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document, docType)
	assert.NoError(t, err)

	err = subcommands.DocumentMain(db)
	stdOutInterceptBuffer.Scan()

	assert.Equal(t, "true", stdOutInterceptBuffer.Text())
	assert.NoError(t, err)
}

// ******************************************************************//
//                         --find-docs-with-tags                    //
// ******************************************************************//.
func TestFindDocumentsWithTags(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	document1 := path.Join(test.TestDataTempDir, t.Name()+"1")
	document2 := path.Join(test.TestDataTempDir, t.Name()+"2")
	docType := "bar"

	tag1 := "foo::bar::baz"
	tag2 := "foo::bar"

	err = libtags.AddTag(db, nil, tag1)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, tag2)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document1, docType)
	assert.NoError(t, err)

	err = libdocuments.AddTag(db, nil, document1, tag1)
	assert.NoError(t, err)

	err = libdocuments.AddTag(db, nil, document1, tag2)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document2, docType)
	assert.NoError(t, err)

	err = libdocuments.AddTag(db, nil, document2, tag1)
	assert.NoError(t, err)

	err = libdocuments.AddTag(db, nil, document2, tag2)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--find-docs-with-tags", tag1, tag2}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, document1, stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, document2, stdOutInterceptBuffer.Text())
}

// ******************************************************************//
//                         --find-links-line                        //
// ******************************************************************//.
func TestFindLinksLine(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	document := path.Join(test.TestDataTempDir, t.Name())
	docType := "bar"

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	link := "- (foo)[bar]"
	_, err = file.WriteString("\n# Links\n" + link)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--find-links-lines", document}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, "2 2", stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, link, stdOutInterceptBuffer.Text())
}

// ******************************************************************//
//                      --find-backlinks-lines                      //
// ******************************************************************//.
func TestFindBackLinksLine(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	document := path.Join(test.TestDataTempDir, t.Name())
	docType := "bar"

	os.Args = []string{"", "document", "--find-backlinks-lines", document}

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	link := "- [foo](bar)"
	_, err = file.WriteString("\n# Backlinks\n" + link)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, document, docType)
	assert.NoError(t, err)

	err = subcommands.DocumentMain(db)

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, "2 2", stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, link, stdOutInterceptBuffer.Text())

	assert.NoError(t, err)
}

// ******************************************************************//
//                            --add-link                            //
// ******************************************************************//.
func TestAddLinkToDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	sourceFileName := t.Name() + "Source"
	destFileName := t.Name() + "Dest"

	sourcePath := path.Join(test.TestDataTempDir, sourceFileName)
	destPath := path.Join(test.TestDataTempDir, destFileName)

	sourceFile, err := test.CreateTestTempFile(sourceFileName)
	assert.NoError(t, err)

	_, err = sourceFile.WriteString("# Links")
	assert.NoError(t, err)

	_, err = test.CreateTestTempFile(destFileName)
	assert.NoError(t, err)

	docType := "bar"
	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, sourcePath, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, destPath, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--add-link", sourcePath, sourcePath}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --remove-link                          //
// ******************************************************************//.
func TestRemoveLinkToDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	sourceFileName := t.Name() + "Source"
	destFileName := t.Name() + "Dest"

	sourcePath := path.Join(test.TestDataTempDir, sourceFileName)
	destPath := path.Join(test.TestDataTempDir, destFileName)

	sourceFile, err := test.CreateTestTempFile(sourceFileName)
	assert.NoError(t, err)

	_, err = sourceFile.WriteString(fmt.Sprintf("# Links\n- ()[%v]\n\n", destPath))
	assert.NoError(t, err)

	_, err = test.CreateTestTempFile(destFileName)
	assert.NoError(t, err)

	docType := "bar"
	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, sourcePath, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, destPath, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--remove-link", sourcePath, destPath}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                            --add-backlink                         //
// ******************************************************************//.
func TestAddBacklinkToDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	sourceFileName := t.Name() + "Source"
	destFileName := t.Name() + "Dest"

	sourcePath := path.Join(test.TestDataTempDir, sourceFileName)
	destPath := path.Join(test.TestDataTempDir, destFileName)

	sourceFile, err := test.CreateTestTempFile(sourceFileName)
	assert.NoError(t, err)

	_, err = sourceFile.WriteString("# Backlinks")
	assert.NoError(t, err)

	_, err = test.CreateTestTempFile(destFileName)
	assert.NoError(t, err)

	docType := "bar"
	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, sourcePath, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, destPath, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--add-backlink", sourcePath, sourcePath}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --remove-backlink                       //
// ******************************************************************//.
func TestRemoveBacklinkToDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	sourceFileName := t.Name() + "Source"
	destFileName := t.Name() + "Dest"

	sourcePath := path.Join(test.TestDataTempDir, sourceFileName)
	destPath := path.Join(test.TestDataTempDir, destFileName)

	sourceFile, err := test.CreateTestTempFile(sourceFileName)
	assert.NoError(t, err)

	_, err = sourceFile.WriteString(fmt.Sprintf("# Backlinks\n- ()[%v]\n\n", destPath))
	assert.NoError(t, err)

	_, err = test.CreateTestTempFile(destFileName)
	assert.NoError(t, err)

	docType := "bar"
	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, sourcePath, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, destPath, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--remove-backlink", sourcePath, destPath}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                             --add-doc                            //
// ******************************************************************//.
func TestAddDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	docPath := path.Join(test.TestDataTempDir, t.Name())

	_, err = test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	docType := "bar"
	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--add-doc", docPath, docType}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --remove-doc                           //
// ******************************************************************//.
func TestRemoveDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	docPath := path.Join(test.TestDataTempDir, t.Name())

	_, err = test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	docType := "bar"
	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, docPath, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--remove-doc", docPath, docType}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --rename-doc                           //
// ******************************************************************//.
func TestRenameDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	docPath := path.Join(test.TestDataTempDir, t.Name())

	_, err = test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	docType := "bar"
	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, docPath, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--rename-doc", docPath, "foo"}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                         --change-doc-type                        //
// ******************************************************************//.
func TestChangeDocumentType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	docPath := path.Join(test.TestDataTempDir, t.Name())

	_, err = test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	oldType := "bar"
	err = libdocuments.AddType(db, nil, oldType)
	assert.NoError(t, err)

	newType := "foo"
	err = libdocuments.AddType(db, nil, newType)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, docPath, oldType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--change-doc-type", docPath, newType}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                          --add-doc-type                          //
// ******************************************************************//.
func TestAddDocumentType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--add-doc-type", "foo"}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}

func TestRemoveDocumentType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	docType := "bar"
	err = libdocuments.AddType(db, nil, docType)
	assert.NoError(t, err)

	os.Args = []string{"", "document", "--remove-doc-type", docType}
	err = subcommands.DocumentMain(db)

	assert.NoError(t, err)
}
