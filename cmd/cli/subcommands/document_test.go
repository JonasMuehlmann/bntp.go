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
	"log"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/cmd/cli/subcommands"
	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

// ******************************************************************//
//                             --add-tag                            //
// ******************************************************************//.
func TestAddTagToDocument(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

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

	subcommands.DocumentMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                           --remove-tag                           //
// ******************************************************************//.
func TestRemoveTagFromDocument(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

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

	subcommands.DocumentMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                           --rename-tag                           //
// ******************************************************************//.
func TestRenameTagInDocument(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

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

	subcommands.DocumentMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                            --get-tags                            //
// ******************************************************************//.
func TestGetTagsFromDocument(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

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

	subcommands.DocumentMain(db, helpers.NOPExiter)
	stdOutInterceptBuffer.Scan()
	assert.Equal(t, tag1, stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, tag2, stdOutInterceptBuffer.Text())

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                         --find-tags-line                         //
// ******************************************************************//.
func TestFindTagsLine(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

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

	subcommands.DocumentMain(db, helpers.NOPExiter)
	stdOutInterceptBuffer.Scan()

	assert.Equal(t, "2 "+tag, stdOutInterceptBuffer.Text())
	assert.Empty(t, logInterceptBuffer.String())
}

func TestHasTags(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

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

	subcommands.DocumentMain(db, helpers.NOPExiter)
	stdOutInterceptBuffer.Scan()

	assert.Equal(t, "true", stdOutInterceptBuffer.Text())
	assert.Empty(t, logInterceptBuffer.String())
}
