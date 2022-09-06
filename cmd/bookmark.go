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
	"github.com/JonasMuehlmann/datastructures.go/maps/hashmap"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/goaoi/functional"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

func WithBookmarkCommand() CliOption {
	return func(cli *Cli) (err error) {
		cli.BookmarkCmd = &cobra.Command{
			Use:   "bookmark",
			Short: "Manage bntp bookmarks",
			Long:  `A longer description`,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				return nil
			},
		}

		cli.BookmarkAddCmd = &cobra.Command{
			Use:   "add MODEL...",
			Short: "Add a bntp bookmark",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				bookmarks, err := UnmarshalEntities[domain.Bookmark](cli, args, cli.InFormat)
				if err != nil {
					return err
				}

				err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), bookmarks)
				if err != nil {
					return err
				}

				return nil
			},
		}

		cli.BookmarkReplaceCmd = &cobra.Command{
			Use:   "replace MODEL...",
			Short: "Replace a bntp bookmark with an updated version",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				bookmarks, err := UnmarshalEntities[domain.Bookmark](cli, args, cli.InFormat)
				if err != nil {
					return err
				}

				err = cli.BNTPBackend.BookmarkManager.Replace(context.Background(), bookmarks)
				if err != nil {
					return err
				}

				return nil
			},
		}

		cli.BookmarkUpsertCmd = &cobra.Command{
			Use:   "upsert MODEL...",
			Short: "Add or replace a bntp bookmark",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				bookmarks, err := UnmarshalEntities[domain.Bookmark](cli, args, cli.InFormat)
				if err != nil {
					return err
				}

				err = cli.BNTPBackend.BookmarkManager.Upsert(context.Background(), bookmarks)
				if err != nil {
					return err
				}

				return nil
			},
		}

		cli.BookmarkEditCmd = &cobra.Command{
			Use:   "edit MODEL...",
			Short: "Edit a bntp bookmark",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 && cli.FilterRaw == "" {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}
				if len(args) > 0 && cli.FilterRaw != "" {
					return ConflictingPositionalArgsAndFlagError{Flag: "filter"}
				}

				var err error
				filter := &domain.BookmarkFilter{}
				updater := &domain.BookmarkUpdater{}
				var numAffectedRecordsRaw int64

				err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(updater, cli.UpdaterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}
				if cli.FilterRaw == "" {
					bookmarks, err := UnmarshalEntities[domain.Bookmark](cli, args, cli.InFormat)
					if err != nil {
						return err
					}
					err = cli.BNTPBackend.BookmarkManager.Update(context.Background(), bookmarks, updater)
					if err != nil {
						return err
					}

					numAffectedRecordsRaw = int64(len(args))
				} else {
					tmp := hashmap.NewFromMap(domain.PredefinedBookmarkFilters)

					if _, err := goaoi.FindIfSlice(tmp.GetKeys(), functional.AreEqualPartial(cli.FilterRaw)); err == nil {
						filter = domain.PredefinedBookmarkFilters[cli.FilterRaw]
					} else {
						err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
						if err != nil {
							return EntityMarshallingError{Inner: err}
						}
					}

					numAffectedRecordsRaw, err = cli.BNTPBackend.BookmarkManager.UpdateWhere(context.Background(), filter, updater)
					if err != nil {
						return err
					}
				}

				numAffectedRecords, err := cli.BNTPBackend.Marshallers[cli.InFormat].Marshall(NumAffectedRecords{numAffectedRecordsRaw})
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), numAffectedRecords)

				return nil
			},
		}

		cli.BookmarkListCmd = &cobra.Command{
			Use:   "list",
			Short: "List bntp bookmarks",
			Long:  `A longer description`,
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				var bookmarks []*domain.Bookmark
				filter := &domain.BookmarkFilter{}
				var output string
				var err error

				if cli.FilterRaw == "" {
					bookmarks, err = cli.BNTPBackend.BookmarkManager.GetAll(context.Background())
					if err != nil {
						return err
					}
				} else {
					tmp := hashmap.NewFromMap(domain.PredefinedBookmarkFilters)

					if _, err := goaoi.FindIfSlice(tmp.GetKeys(), functional.AreEqualPartial(cli.FilterRaw)); err == nil {
						filter = domain.PredefinedBookmarkFilters[cli.FilterRaw]
					} else {
						err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
						if err != nil {
							return EntityMarshallingError{Inner: err}
						}
					}

					bookmarks, err = cli.BNTPBackend.BookmarkManager.GetWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				output, err = cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(bookmarks)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), output)

				return nil
			},
		}

		cli.BookmarkRemoveCmd = &cobra.Command{
			Use:   "remove [MODEL...]",
			Short: "Remove a bntp bookmark",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 && cli.FilterRaw == "" {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}
				if len(args) > 0 && cli.FilterRaw != "" {
					return ConflictingPositionalArgsAndFlagError{Flag: "filter"}
				}

				filter := &domain.BookmarkFilter{}
				var err error
				var numAffectedRecordsRaw int64

				if cli.FilterRaw == "" {
					bookmarks, err := UnmarshalEntities[domain.Bookmark](cli, args, cli.InFormat)
					if err != nil {
						return err
					}

					err = cli.BNTPBackend.BookmarkManager.Delete(context.Background(), bookmarks)
					if err != nil {
						return err
					}

					numAffectedRecordsRaw = int64(len(args))
				} else {
					tmp := hashmap.NewFromMap(domain.PredefinedBookmarkFilters)

					if _, err := goaoi.FindIfSlice(tmp.GetKeys(), functional.AreEqualPartial(cli.FilterRaw)); err == nil {
						filter = domain.PredefinedBookmarkFilters[cli.FilterRaw]
					} else {
						err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
						if err != nil {
							return EntityMarshallingError{Inner: err}
						}
					}

					numAffectedRecordsRaw, err = cli.BNTPBackend.BookmarkManager.DeleteWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				numAffectedRecords, err := cli.BNTPBackend.Marshallers[cli.InFormat].Marshall(NumAffectedRecords{numAffectedRecordsRaw})
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), numAffectedRecords)

				return nil
			},
		}

		cli.BookmarkFindCmd = &cobra.Command{
			Use:   "find-first",
			Short: "Find the first bookmark matching a filter",
			Long:  `A longer description`,
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				filter := &domain.BookmarkFilter{}
				var err error
				var result *domain.Bookmark
				var output string

				tmp := hashmap.NewFromMap(domain.PredefinedBookmarkFilters)

				if _, err := goaoi.FindIfSlice(tmp.GetKeys(), functional.AreEqualPartial(cli.FilterRaw)); err == nil {
					filter = domain.PredefinedBookmarkFilters[cli.FilterRaw]
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}
				}

				result, err = cli.BNTPBackend.BookmarkManager.GetFirstWhere(context.Background(), filter)
				if err != nil {
					return err
				}

				output, err = cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(result)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), output)

				return nil
			},
		}
		cli.BookmarkCountCmd = &cobra.Command{
			Use:   "count",
			Short: "Manage bntp bookmarks",
			Long:  `A longer description`,
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				filter := &domain.BookmarkFilter{}
				var countRaw int64
				var err error

				if cli.FilterRaw == "" {
					countRaw, err = cli.BNTPBackend.BookmarkManager.CountAll(context.Background())
					if err != nil {
						return err
					}
				} else {

					tmp := hashmap.NewFromMap(domain.PredefinedBookmarkFilters)

					if _, err := goaoi.FindIfSlice(tmp.GetKeys(), functional.AreEqualPartial(cli.FilterRaw)); err == nil {
						filter = domain.PredefinedBookmarkFilters[cli.FilterRaw]
					} else {
						err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
						if err != nil {
							return EntityMarshallingError{Inner: err}
						}
					}

					countRaw, err = cli.BNTPBackend.BookmarkManager.CountWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				count, err := cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(Count{countRaw})
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), count)

				return nil
			},
		}
		cli.BookmarkDoesExistCmd = &cobra.Command{
			Use:   "does-exist [MODEL]",
			Short: "Manage bntp bookmarks",
			Long:  `A longer description`,
			Args:  cobra.RangeArgs(0, 1),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 && cli.FilterRaw == "" {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}
				if len(args) > 0 && cli.FilterRaw != "" {
					return ConflictingPositionalArgsAndFlagError{Flag: "filter"}
				}

				filter := &domain.BookmarkFilter{}
				var err error
				bookmark := &domain.Bookmark{}
				var doesExistRaw bool

				if cli.FilterRaw == "" {
					err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(bookmark, args[0])
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					doesExistRaw, err = cli.BNTPBackend.BookmarkManager.DoesExist(context.Background(), bookmark)
					if err != nil {
						return err
					}
				} else {

					tmp := hashmap.NewFromMap(domain.PredefinedBookmarkFilters)

					if _, err := goaoi.FindIfSlice(tmp.GetKeys(), functional.AreEqualPartial(cli.FilterRaw)); err == nil {
						filter = domain.PredefinedBookmarkFilters[cli.FilterRaw]
					} else {
						err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
						if err != nil {
							return EntityMarshallingError{Inner: err}
						}
					}

					doesExistRaw, err = cli.BNTPBackend.BookmarkManager.DoesExistWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				doesExist, err := cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(DoesExist{doesExistRaw})
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), doesExist)

				return nil
			},
		}

		cli.RootCmd.AddCommand(cli.BookmarkCmd)

		cli.BookmarkCmd.AddCommand(cli.BookmarkListCmd)
		cli.BookmarkCmd.AddCommand(cli.BookmarkReplaceCmd)
		cli.BookmarkCmd.AddCommand(cli.BookmarkEditCmd)
		cli.BookmarkCmd.AddCommand(cli.BookmarkAddCmd)
		cli.BookmarkCmd.AddCommand(cli.BookmarkRemoveCmd)
		cli.BookmarkCmd.AddCommand(cli.BookmarkCountCmd)
		cli.BookmarkCmd.AddCommand(cli.BookmarkDoesExistCmd)
		cli.BookmarkCmd.AddCommand(cli.BookmarkFindCmd)
		cli.BookmarkCmd.AddCommand(cli.BookmarkUpsertCmd)

		for _, subcommand := range cli.BookmarkCmd.Commands() {
			if slices.Contains([]*cobra.Command{cli.BookmarkAddCmd, cli.BookmarkListCmd, cli.BookmarkRemoveCmd, cli.BookmarkFindCmd, cli.BookmarkDoesExistCmd}, subcommand) {
				subcommand.PersistentFlags().StringVar(&cli.InFormat, "out-format", "json", "The serialization format to use for reading input")
				subcommand.PersistentFlags().StringVar(&cli.OutFormat, "in-format", "json", "The serialization format to use for writing output")
			}
		}

		for _, subcommand := range cli.BookmarkCmd.Commands() {
			if slices.Contains([]*cobra.Command{cli.BookmarkEditCmd, cli.BookmarkListCmd, cli.BookmarkRemoveCmd, cli.BookmarkFindCmd, cli.BookmarkCountCmd, cli.BookmarkDoesExistCmd}, subcommand) {
				subcommand.PersistentFlags().StringVar(&cli.FilterRaw, "filter", "", "The filter to use for processing entities")
			}
		}

		for _, subcommand := range cli.BookmarkCmd.Commands() {
			if slices.Contains([]*cobra.Command{cli.BookmarkEditCmd}, subcommand) {
				subcommand.PersistentFlags().StringVar(&cli.UpdaterRaw, "updater", "", "The updater to use for processing entities")
			}
		}

		cli.BookmarkFindCmd.MarkPersistentFlagRequired("filter")
		cli.BookmarkEditCmd.MarkPersistentFlagRequired("updater")

		return
	}
}
