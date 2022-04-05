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
	"log"
	"os"
	"strings"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/cmd/cli/subcommands"
	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

// ******************************************************************//
//                               --add                              //
// ******************************************************************//.
func TestAddLink(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "link", "--add", "foo", "bar"}

	err = libdocuments.AddType(db, nil, "type")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "foo", "type")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "bar", "type")
	assert.NoError(t, err)

	subcommands.LinkMain(db, helpers.NOPExiter)
	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                             --remove                             //
// ******************************************************************//.
func TestRemoveLink(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "link", "--remove", "foo", "bar"}

	err = libdocuments.AddType(db, nil, "type")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "foo", "type")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "bar", "type")
	assert.NoError(t, err)

	subcommands.LinkMain(db, helpers.NOPExiter)
	assert.Empty(t, logInterceptBuffer.String())
}
