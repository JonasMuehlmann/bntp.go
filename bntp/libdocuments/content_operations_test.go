package libdocuments_test

import (
	"context"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/bntp/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

func TestAddTags(t *testing.T) {
	tests := []struct {
		name               string
		tagsToAdd          []string
		content            string
		expectedNewContent string
		err                error
		errorMatcher       testCommon.OutputValidator
	}{
		{
			name:      "empty document",
			tagsToAdd: []string{"foo"},
			content:   "",
			err:       helper.IneffectiveOperationError{},
		},
		{
			name:      "no tag line",
			tagsToAdd: []string{"foo"},
			content:   "# Links\n# Backlinks\n",
			err:       libdocuments.DocumentSyntaxError{},
		},
		{
			name:      "no tags",
			tagsToAdd: []string{},
			content:   "# Tags",
			err:       helper.IneffectiveOperationError{},
		},
		{
			name:      "empty tags",
			tagsToAdd: []string{"", ""},
			content:   "# Tags",
			err:       helper.NilInputError{},
		},
		{
			name:               "add first tags",
			tagsToAdd:          []string{"foo", "bar"},
			content:            "# Tags",
			expectedNewContent: "# Tags\nfoo,bar",
		},
		{
			name:               "add with one existing",
			tagsToAdd:          []string{"bar", "baz"},
			content:            "# Tags\nfoo",
			expectedNewContent: "# Tags\nfoo,bar,baz",
		},
		{
			name:               "add with two existing",
			tagsToAdd:          []string{"baz"},
			content:            "# Tags\nfoo,bar",
			expectedNewContent: "# Tags\nfoo,bar,baz",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			newContent, err := libdocuments.AddTags(context.Background(), test.content, test.tagsToAdd)
			assert.Equal(t, test.expectedNewContent, newContent, test.name+", assert new content matches expected")

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestRemoveTags(t *testing.T) {
	tests := []struct {
		name               string
		tagsToRemove       []string
		content            string
		expectedNewContent string
		err                error
		errorMatcher       testCommon.OutputValidator
	}{
		{
			name:         "empty document",
			tagsToRemove: []string{"foo"},
			content:      "",
			err:          helper.IneffectiveOperationError{},
		},
		{
			name:         "no tag line",
			tagsToRemove: []string{"foo"},
			content:      "# Links\n# Backlinks\n",
			err:          libdocuments.DocumentSyntaxError{},
		},
		{
			name:         "no tags",
			tagsToRemove: []string{},
			content:      "# Tags",
			err:          helper.IneffectiveOperationError{},
		},
		{
			name:         "empty tags",
			tagsToRemove: []string{"", ""},
			content:      "# Tags",
			err:          helper.NilInputError{},
		},
		{
			name:               "remove first tags",
			tagsToRemove:       []string{"foo"},
			content:            "# Tags\nfoo",
			expectedNewContent: "# Tags\n",
		},
		{
			name:               "remove second",
			tagsToRemove:       []string{"bar"},
			content:            "# Tags\nfoo,bar",
			expectedNewContent: "# Tags\nfoo",
		},
		{
			name:               "remove all",
			tagsToRemove:       []string{"foo", "bar"},
			content:            "# Tags\nfoo,bar",
			expectedNewContent: "# Tags\n",
		},
		{
			name:         "remove non-existent",
			tagsToRemove: []string{"foo"},
			content:      "# Tags",
			err:          helper.IneffectiveOperationError{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			newContent, err := libdocuments.RemoveTags(context.Background(), test.content, test.tagsToRemove)
			assert.Equal(t, test.expectedNewContent, newContent, test.name+", assert new content matches expected")

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestAddLinks(t *testing.T) {
	tests := []struct {
		name               string
		linksToAdd         []string
		content            string
		expectedNewContent string
		err                error
		errorMatcher       testCommon.OutputValidator
	}{
		{
			name:       "empty document",
			linksToAdd: []string{"foo"},
			content:    "",
			err:        helper.IneffectiveOperationError{},
		},
		{
			name:       "no links line",
			linksToAdd: []string{"foo"},
			content:    "#Tags\n# Backlinks\n",
			err:        libdocuments.DocumentSyntaxError{},
		},
		{
			name:       "no links",
			linksToAdd: []string{},
			content:    "# Links",
			err:        helper.IneffectiveOperationError{},
		},
		{
			name:       "empty links",
			linksToAdd: []string{"", ""},
			content:    "# Links",
			err:        helper.NilInputError{},
		},
		{
			name:               "add first links",
			linksToAdd:         []string{"foo", "bar"},
			content:            "# Links",
			expectedNewContent: "# Links\n- (foo)[foo]\n- (bar)[bar]",
		},
		{
			name:               "add with one existing",
			linksToAdd:         []string{"bar", "baz"},
			content:            "# Links\n- (foo)[foo]",
			expectedNewContent: "# Links\n- (foo)[foo]\n- (bar)[bar]\n- (baz)[baz]",
		},
		{
			name:               "add with two existing",
			linksToAdd:         []string{"baz"},
			content:            "# Links\n- (foo)[foo]\n- (bar)[bar]\n- (baz)[bar]",
			expectedNewContent: "# Links\n- (foo)[foo]\n- (bar)[bar]\n- (baz)[bar]\n- (baz)[baz]",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			newContent, err := libdocuments.AddLinks(context.Background(), test.content, test.linksToAdd)
			assert.Equal(t, test.expectedNewContent, newContent, test.name+", assert new content matches expected")

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestRemoveLinks(t *testing.T) {
	tests := []struct {
		name               string
		linksToRemove      []string
		content            string
		expectedNewContent string
		err                error
		errorMatcher       testCommon.OutputValidator
	}{
		{
			name:          "empty document",
			linksToRemove: []string{"foo"},
			content:       "",
			err:           helper.IneffectiveOperationError{},
		},
		{
			name:          "no links line",
			linksToRemove: []string{"foo"},
			content:       "#Tags\n# Backlinks\n",
			err:           libdocuments.DocumentSyntaxError{},
		},
		{
			name:          "no links",
			linksToRemove: []string{},
			content:       "# Links",
			err:           helper.IneffectiveOperationError{},
		},
		{
			name:          "empty links",
			linksToRemove: []string{"", ""},
			content:       "# Links",
			err:           helper.NilInputError{},
		},
		{
			name:               "remove first links",
			linksToRemove:      []string{"foo"},
			content:            "# Links\n- (foo)[foo]",
			expectedNewContent: "# Links\n",
		},
		{
			name:               "remove second link",
			linksToRemove:      []string{"bar"},
			content:            "# Links\n- (foo)[foo]\n- (bar)[bar]",
			expectedNewContent: "# Links\n- (foo)[foo]",
		},
		{
			name:               "remove all",
			linksToRemove:      []string{"foo", "bar"},
			content:            "# Links\n- (foo)[foo]\n- (bar)[bar]",
			expectedNewContent: "# Links\n",
		},
		{
			name:          "remove non-existent",
			linksToRemove: []string{"foo"},
			content:       "# Links",
			err:           helper.IneffectiveOperationError{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			newContent, err := libdocuments.RemoveLinks(context.Background(), test.content, test.linksToRemove)
			assert.Equal(t, test.expectedNewContent, newContent, test.name+", assert new content matches expected")

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestAddBacklinks(t *testing.T) {
	tests := []struct {
		name               string
		linksToAdd         []string
		content            string
		expectedNewContent string
		err                error
		errorMatcher       testCommon.OutputValidator
	}{
		{
			name:       "empty document",
			linksToAdd: []string{"foo"},
			content:    "",
			err:        helper.IneffectiveOperationError{},
		},
		{
			name:       "no links line",
			linksToAdd: []string{"foo"},
			content:    "#Tags\n# Links\n",
			err:        libdocuments.DocumentSyntaxError{},
		},
		{
			name:       "no links",
			linksToAdd: []string{},
			content:    "# Backlinks",
			err:        helper.IneffectiveOperationError{},
		},
		{
			name:       "empty links",
			linksToAdd: []string{"", ""},
			content:    "# Backlinks",
			err:        helper.NilInputError{},
		},
		{
			name:               "add first links",
			linksToAdd:         []string{"foo", "bar"},
			content:            "# Backlinks",
			expectedNewContent: "# Backlinks\n- (foo)[foo]\n- (bar)[bar]",
		},
		{
			name:               "add with one existing",
			linksToAdd:         []string{"bar", "baz"},
			content:            "# Backlinks\n- (foo)[foo]",
			expectedNewContent: "# Backlinks\n- (foo)[foo]\n- (bar)[bar]\n- (baz)[baz]",
		},
		{
			name:               "add with two existing",
			linksToAdd:         []string{"baz"},
			content:            "# Backlinks\n- (foo)[foo]\n- (bar)[bar]\n- (baz)[bar]",
			expectedNewContent: "# Backlinks\n- (foo)[foo]\n- (bar)[bar]\n- (baz)[bar]\n- (baz)[baz]",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			newContent, err := libdocuments.AddBacklinks(context.Background(), test.content, test.linksToAdd)
			assert.Equal(t, test.expectedNewContent, newContent, test.name+", assert new content matches expected")

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestRemoveBacklinks(t *testing.T) {
	tests := []struct {
		name               string
		linksToRemove      []string
		content            string
		expectedNewContent string
		err                error
		errorMatcher       testCommon.OutputValidator
	}{
		{
			name:          "empty document",
			linksToRemove: []string{"foo"},
			content:       "",
			err:           helper.IneffectiveOperationError{},
		},
		{
			name:          "no links line",
			linksToRemove: []string{"foo"},
			content:       "#Tags\n# Links\n",
			err:           libdocuments.DocumentSyntaxError{},
		},
		{
			name:          "no links",
			linksToRemove: []string{},
			content:       "# Backlinks",
			err:           helper.IneffectiveOperationError{},
		},
		{
			name:          "empty links",
			linksToRemove: []string{"", ""},
			content:       "# Backlinks",
			err:           helper.NilInputError{},
		},
		{
			name:               "remove first links",
			linksToRemove:      []string{"foo"},
			content:            "# Backlinks\n- (foo)[foo]",
			expectedNewContent: "# Backlinks\n",
		},
		{
			name:               "remove second link",
			linksToRemove:      []string{"bar"},
			content:            "# Backlinks\n- (foo)[foo]\n- (bar)[bar]",
			expectedNewContent: "# Backlinks\n- (foo)[foo]",
		},
		{
			name:               "remove all",
			linksToRemove:      []string{"foo", "bar"},
			content:            "# Backlinks\n- (foo)[foo]\n- (bar)[bar]",
			expectedNewContent: "# Backlinks\n",
		},
		{
			name:          "remove non-existent",
			linksToRemove: []string{"foo"},
			content:       "# Backlinks",
			err:           helper.IneffectiveOperationError{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			newContent, err := libdocuments.RemoveBacklinks(context.Background(), test.content, test.linksToRemove)
			assert.Equal(t, test.expectedNewContent, newContent, test.name+", assert new content matches expected")

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}
