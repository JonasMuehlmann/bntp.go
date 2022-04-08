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
	"os"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/cmd/cli/subcommands"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

// ******************************************************************//
//                             --import                             //
// ******************************************************************//.
func TestImport(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	yml := `
tags:
- bar
- baz
    `
	_, err = file.WriteString(yml)
	assert.NoError(t, err)

	os.Args = []string{"", "tag", "--import", file.Name()}

	err = subcommands.TagMain(db)
	assert.NoError(t, err)
}

func TestImportFileDoesNotExist(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "tag", "--import", "foo"}

	err = subcommands.TagMain(db)
	assert.Error(t, err)
}

// ******************************************************************//
//                             --export                             //
// ******************************************************************//.
func TestExport(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	os.Args = []string{"", "tag", "--export", file.Name()}

	err = libtags.AddTag(db, nil, "foo")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "bar")
	assert.NoError(t, err)

	err = subcommands.TagMain(db)
}

// ******************************************************************//
//                            --ambiguous                           //
// ******************************************************************//.
func TestAmbiguousNotAmbiguous(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag := "foo::bar::baz"
	os.Args = []string{"", "tag", "--ambiguous", tag}

	err = libtags.AddTag(db, nil, tag)
	assert.NoError(t, err)

	err = subcommands.TagMain(db)
	stdOutInterceptBuffer.Scan()

	assert.Equal(t, "false", stdOutInterceptBuffer.Text())
	assert.NoError(t, err)
}

func TestAmbiguousAmbiguous(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag1 := "foo::bar::baz"
	tag2 := "foo::foo::baz"
	os.Args = []string{"", "tag", "--ambiguous", tag1}

	err = libtags.AddTag(db, nil, tag1)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, tag2)
	assert.NoError(t, err)

	err = subcommands.TagMain(db)
	stdOutInterceptBuffer.Scan()

	assert.Equal(t, "true", stdOutInterceptBuffer.Text())
	assert.NoError(t, err)
}

// ******************************************************************//
//                            --comonent                            //
// ******************************************************************//.
func TestAmbiguousComponent(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag1 := "foo::bar::baz"
	tag2 := "foo::foo::baz"
	os.Args = []string{"", "tag", "--component", tag1}

	err = libtags.AddTag(db, nil, tag1)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, tag2)
	assert.NoError(t, err)

	err = subcommands.TagMain(db)
	stdOutInterceptBuffer.Scan()

	assert.Equal(t, "1 bar", stdOutInterceptBuffer.Text())
	assert.NoError(t, err)
}

// ******************************************************************//
//                               --add                              //
// ******************************************************************//.
func TestAddTag(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "tag", "--add", "foo"}

	err = subcommands.TagMain(db)
	assert.NoError(t, err)
}

// ******************************************************************//
//                             --remove                             //
// ******************************************************************//.
func TestRemove(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "tag", "--remove", "foo"}

	err = subcommands.TagMain(db)
	assert.NoError(t, err)
}

// ******************************************************************//
//                             --rename                             //
// ******************************************************************//.
func TestRename(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "tag", "--rename", "foo", "bar"}

	err = subcommands.TagMain(db)
	assert.NoError(t, err)
}

// ******************************************************************//
//                             --shorten                            //
// ******************************************************************//.
func TestShorten(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag1 := "foo::foo::baz"
	os.Args = []string{"", "tag", "--shorten", tag1}

	err = libtags.AddTag(db, nil, tag1)
	assert.NoError(t, err)

	err = subcommands.TagMain(db)
	stdOutInterceptBuffer.Scan()

	assert.NoError(t, err)
	assert.Equal(t, "baz", stdOutInterceptBuffer.Text())
}

// ******************************************************************//
//                              --list                              //
// ******************************************************************//.
func TestList(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag1 := "foo::foo::baz"
	tag2 := "foo::bar::baz"
	os.Args = []string{"", "tag", "--list"}

	err = libtags.AddTag(db, nil, tag1)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, tag2)
	assert.NoError(t, err)

	err = subcommands.TagMain(db)
	stdOutInterceptBuffer.Scan()

	assert.NoError(t, err)
	assert.Equal(t, tag1, stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()

	assert.NoError(t, err)
	assert.Equal(t, tag2, stdOutInterceptBuffer.Text())
}

// ******************************************************************//
//                           --list-short                           //
// ******************************************************************//.
func TestListShort(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	tag1 := "foo::foo::baz"
	tag2 := "foo::bar::baz"
	os.Args = []string{"", "tag", "--list-short"}

	err = libtags.AddTag(db, nil, tag1)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, tag2)
	assert.NoError(t, err)

	err = subcommands.TagMain(db)
	stdOutInterceptBuffer.Scan()

	assert.NoError(t, err)
	assert.Equal(t, "foo::baz", stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()

	assert.NoError(t, err)
	assert.Equal(t, "bar::baz", stdOutInterceptBuffer.Text())
}
