// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cmd

import (
	"context"
	"fmt"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var Format, FilterRaw, UpdaterRaw string

var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Manage bntp bookmarks",
	Long:  `A longer description`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		return nil
	},
}

var bookmarkAddCmd = &cobra.Command{
	Use:   "add MODEL...",
	Short: "Add a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		bookmarks := make([]*domain.Bookmark, len(args))

		for i, arg := range args {
			bookmarks[i] = new(domain.Bookmark)

			err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarks[i], arg)
			if err != nil {
				return err
			}
		}

		err := BNTPBackend.BookmarkManager.Add(context.Background(), bookmarks)
		if err != nil {
			return err
		}

		return nil
	},
}

var bookmarkReplaceCmd = &cobra.Command{
	Use:   "replace MODEL...",
	Short: "Replace a bntp bookmark with an updated version",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		bookmarks := make([]*domain.Bookmark, 0, len(args))

		for i, bookmarkOut := range bookmarks {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
			if err != nil {
				return err
			}
		}

		err := BNTPBackend.BookmarkManager.Replace(context.Background(), bookmarks)
		if err != nil {
			return err
		}

		return nil
	},
}

var bookmarkUpsertCmd = &cobra.Command{
	Use:   "upsert MODEL...",
	Short: "Add or replace a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		bookmarks := make([]*domain.Bookmark, 0, len(args))

		for i, bookmarkOut := range bookmarks {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
			if err != nil {
				return err
			}
		}

		err := BNTPBackend.BookmarkManager.Upsert(context.Background(), bookmarks)
		if err != nil {
			return err
		}

		return nil
	},
}

var bookmarkEditCmd = &cobra.Command{
	Use:   "edit MODEL...",
	Short: "Edit a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var err error
		var filter *domain.BookmarkFilter
		var updater *domain.BookmarkUpdater
		var numAffectedRecords int64

		bookmarks := make([]*domain.Bookmark, 0, len(args))

		for i, bookmarkOut := range bookmarks {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
			if err != nil {
				return err
			}
		}

		err = BNTPBackend.Unmarshallers[Format].Unmarshall(updater, UpdaterRaw)
		if err != nil {
			return err
		}
		if FilterRaw == "" {
			err := BNTPBackend.BookmarkManager.Update(context.Background(), bookmarks, updater)
			if err != nil {
				return err
			}

			numAffectedRecords = int64(len(args))
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return err
			}

			numAffectedRecords, err = BNTPBackend.BookmarkManager.UpdateWhere(context.Background(), filter, updater)
			if err != nil {
				return err
			}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), numAffectedRecords)

		return nil
	},
}

var bookmarkListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var bookmarks []*domain.Bookmark
		var filter *domain.BookmarkFilter
		var output string
		var err error

		if FilterRaw == "" {
			bookmarks, err = BNTPBackend.BookmarkManager.GetAll(context.Background())
			if err != nil {
				return err
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return err
			}

			bookmarks, err = BNTPBackend.BookmarkManager.GetWhere(context.Background(), filter)
			if err != nil {
				return err
			}
		}

		output, err = BNTPBackend.Marshallers[Format].Marshall(bookmarks)
		if err != nil {
			return err
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), output)

		return nil
	},
}

var bookmarkRemoveCmd = &cobra.Command{
	Use:   "remove [MODEL...]",
	Short: "Remove a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var filter *domain.BookmarkFilter
		var err error
		var numAffectedRecords int64

		if FilterRaw == "" {
			bookmarks := make([]*domain.Bookmark, 0, len(args))

			for i, bookmarkOut := range bookmarks {
				err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
				if err != nil {
					return err
				}
			}

			err = BNTPBackend.BookmarkManager.Delete(context.Background(), bookmarks)
			if err != nil {
				return err
			}

			numAffectedRecords = int64(len(args))
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return err
			}

			numAffectedRecords, err = BNTPBackend.BookmarkManager.DeleteWhere(context.Background(), filter)
			if err != nil {
				return err
			}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), numAffectedRecords)

		return nil
	},
}

var bookmarkFindCmd = &cobra.Command{
	Use:   "find-first",
	Short: "Find the first bookmark matching a filter",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var filter *domain.BookmarkFilter
		var err error
		var result *domain.Bookmark
		var output string

		err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
		if err != nil {
			return err
		}

		result, err = BNTPBackend.BookmarkManager.GetFirstWhere(context.Background(), filter)
		if err != nil {
			return err
		}

		output, err = BNTPBackend.Marshallers[Format].Marshall(result)
		if err != nil {
			return err
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), output)

		return nil
	},
}
var bookmarkCountCmd = &cobra.Command{
	Use:   "count",
	Short: "Manage bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var filter *domain.BookmarkFilter
		var count int64
		var err error

		if FilterRaw == "" {
			count, err = BNTPBackend.BookmarkManager.CountAll(context.Background())
			if err != nil {
				return err
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return err
			}

			count, err = BNTPBackend.BookmarkManager.CountWhere(context.Background(), filter)
			if err != nil {
				return err
			}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), count)

		return nil
	},
}
var bookmarkDoesExistCmd = &cobra.Command{
	Use:   "does-exist [MODEL]",
	Short: "Manage bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var filter *domain.BookmarkFilter
		var err error
		var bookmark *domain.Bookmark
		var doesExist bool

		if FilterRaw == "" {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(bookmark, args[0])
			if err != nil {
				return err
			}

			doesExist, err = BNTPBackend.BookmarkManager.DoesExist(context.Background(), bookmark)
			if err != nil {
				return err
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return err
			}

			doesExist, err = BNTPBackend.BookmarkManager.DoesExistWhere(context.Background(), filter)
			if err != nil {
				return err
			}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), doesExist)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(bookmarkCmd)
	bookmarkCmd.AddCommand(bookmarkListCmd)
	bookmarkCmd.AddCommand(bookmarkReplaceCmd)
	bookmarkCmd.AddCommand(bookmarkEditCmd)
	bookmarkCmd.AddCommand(bookmarkAddCmd)
	bookmarkCmd.AddCommand(bookmarkRemoveCmd)
	bookmarkCmd.AddCommand(bookmarkCountCmd)
	bookmarkCmd.AddCommand(bookmarkDoesExistCmd)
	bookmarkCmd.AddCommand(bookmarkFindCmd)
	bookmarkCmd.AddCommand(bookmarkUpsertCmd)

	bookmarkFindCmd.MarkFlagRequired("filter")
	bookmarkEditCmd.MarkFlagRequired("updater")

	bookmarkCmd.AddCommand(bookmarkTypeCmd)
	bookmarkTypeCmd.AddCommand(bookmarkTypeAddCmd)
	bookmarkTypeCmd.AddCommand(bookmarkTypeEditCmd)
	bookmarkTypeCmd.AddCommand(bookmarkTypeRemoveCmd)

	for _, subcommand := range bookmarkCmd.Commands() {
		if slices.Contains([]*cobra.Command{bookmarkAddCmd, bookmarkListCmd, bookmarkRemoveCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&Format, "format", "json", "The serialization format to use for i/o")
		}
	}

	for _, subcommand := range bookmarkCmd.Commands() {
		if slices.Contains([]*cobra.Command{bookmarkEditCmd, bookmarkListCmd, bookmarkRemoveCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&FilterRaw, "filter", "", "The filter to use for processing entities")
		}
	}

	for _, subcommand := range bookmarkCmd.Commands() {
		if slices.Contains([]*cobra.Command{bookmarkEditCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&FilterRaw, "updater", "", "The updater to use for processing entities")
		}
	}
}
