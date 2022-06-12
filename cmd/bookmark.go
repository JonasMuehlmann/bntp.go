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
	"os"

	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var Format, FilterRaw, UpdaterRaw string

var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Manage bntp bookmarks",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkAddCmd = &cobra.Command{
	Use:   "add DATA...",
	Short: "Add a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		bookmarks := make([]*domain.Bookmark, 0, len(args))

		for i, bookmarkOut := range bookmarks {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
			if err != nil {
				panic(err)
			}
		}

		err := BNTPBackend.BookmarkManager.Add(context.Background(), bookmarks)
		if err != nil {
			panic(err)
		}
	},
}

var bookmarkReplaceCmd = &cobra.Command{
	Use:   "replace NEW_DATA...",
	Short: "Replace a bntp bookmark with an updated version",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		bookmarks := make([]*domain.Bookmark, 0, len(args))

		for i, bookmarkOut := range bookmarks {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
			if err != nil {
				panic(err)
			}
		}

		err := BNTPBackend.BookmarkManager.Replace(context.Background(), bookmarks)
		if err != nil {
			panic(err)
		}
	},
}

var bookmarkUpsertCmd = &cobra.Command{
	Use:   "upsert NEW_DATA...",
	Short: "Add or replace a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		bookmarks := make([]*domain.Bookmark, 0, len(args))

		for i, bookmarkOut := range bookmarks {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
			if err != nil {
				panic(err)
			}
		}

		err := BNTPBackend.BookmarkManager.Upsert(context.Background(), bookmarks)
		if err != nil {
			panic(err)
		}
	},
}

var bookmarkEditCmd = &cobra.Command{
	Use:   "edit NEW_DATA...",
	Short: "Edit a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		var err error
		var filter *domain.BookmarkFilter
		var updater *domain.BookmarkUpdater
		var numAffectedRecords int64

		bookmarks := make([]*domain.Bookmark, 0, len(args))

		for i, bookmarkOut := range bookmarks {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
			if err != nil {
				panic(err)
			}
		}

		err = BNTPBackend.Unmarshallers[Format].Unmarshall(updater, UpdaterRaw)
		if err != nil {
			panic(err)
		}
		// TODO: We should also have Update and Upsert methods in the managers and repositories
		if FilterRaw == "" {
			err := BNTPBackend.BookmarkManager.Update(context.Background(), bookmarks, updater)
			if err != nil {
				panic(err)
			}

			numAffectedRecords = int64(len(args))
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				panic(err)
			}

			numAffectedRecords, err = BNTPBackend.BookmarkManager.UpdateWhere(context.Background(), filter, updater)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(numAffectedRecords)
	},
}

// TODO: Implement output filters
var bookmarkListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		var bookmarks []*domain.Bookmark
		var filter *domain.BookmarkFilter
		var output string
		var err error

		if FilterRaw == "" {
			bookmarks, err = BNTPBackend.BookmarkManager.GetAll(context.Background())
			if err != nil {
				panic(err)
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				panic(err)
			}

			bookmarks, err = BNTPBackend.BookmarkManager.GetWhere(context.Background(), filter)
			if err != nil {
				panic(err)
			}
		}

		output, err = BNTPBackend.Marshallers[Format].Marshall(bookmarks)
		if err != nil {
			panic(err)
		}

		fmt.Println(output)
	},
}

var bookmarkRemoveCmd = &cobra.Command{
	Use:   "remove [BOOKMARK...]",
	Short: "Remove a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		var filter *domain.BookmarkFilter
		var err error
		var numAffectedRecords int64

		if FilterRaw == "" {
			bookmarks := make([]*domain.Bookmark, 0, len(args))

			for i, bookmarkOut := range bookmarks {
				err := BNTPBackend.Unmarshallers[Format].Unmarshall(bookmarkOut, args[i])
				if err != nil {
					panic(err)
				}
			}

			err = BNTPBackend.BookmarkManager.Delete(context.Background(), bookmarks)
			if err != nil {
				panic(err)
			}

			numAffectedRecords = int64(len(args))
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				panic(err)
			}

			numAffectedRecords, err = BNTPBackend.BookmarkManager.DeleteWhere(context.Background(), filter)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(numAffectedRecords)
	},
}

var bookmarkFindCmd = &cobra.Command{
	Use:   "find-first",
	Short: "Find the first bookmark matching a filter",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		var filter *domain.BookmarkFilter
		var err error
		var result *domain.Bookmark
		var output string

		err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
		if err != nil {
			panic(err)
		}

		result, err = BNTPBackend.BookmarkManager.GetFirstWhere(context.Background(), filter)
		if err != nil {
			panic(err)
		}

		output, err = BNTPBackend.Marshallers[Format].Marshall(result)
		if err != nil {
			panic(err)
		}

		fmt.Println(output)
	},
}
var bookmarkCountCmd = &cobra.Command{
	Use:   "count",
	Short: "Manage bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		var filter *domain.BookmarkFilter
		var count int64
		var err error

		if FilterRaw == "" {
			count, err = BNTPBackend.BookmarkManager.CountAll(context.Background())
			if err != nil {
				panic(err)
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				panic(err)
			}

			count, err = BNTPBackend.BookmarkManager.CountWhere(context.Background(), filter)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(count)
	},
}
var bookmarkDoesExistCmd = &cobra.Command{
	Use:   "does-exist [DATA]",
	Short: "Manage bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		var filter *domain.BookmarkFilter
		var err error
		var bookmark *domain.Bookmark
		var doesExist bool

		if FilterRaw == "" {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(bookmark, args[0])
			if err != nil {
				panic(err)
			}

			doesExist, err = BNTPBackend.BookmarkManager.DoesExist(context.Background(), bookmark)
			if err != nil {
				panic(err)
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				panic(err)
			}

			doesExist, err = BNTPBackend.BookmarkManager.DoesExistWhere(context.Background(), filter)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(doesExist)
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
