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
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

func WithTagCommand() CliOption {
	return func(cli *Cli) {
		cli.TagCmd = &cobra.Command{
			Use:   "tag",
			Short: "Manage tags available for bntp entities",
			Long:  `A longer description`,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				return nil
			},
		}

		cli.TagAddCmd = &cobra.Command{
			Use:   "add TAG...",
			Short: "Add new bntp tags",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}
				tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.InFormat)
				if err != nil {
					return err
				}
				err = cli.BNTPBackend.TagManager.Add(context.Background(), tags)

				return err
			},
		}

		// TODO: If ambiguous, return ambiguous component
		cli.TagAmbiguousCmd = &cobra.Command{
			Use:   "ambiguous TAG...",
			Short: "Check if bntp tag's leafs are ambiguous",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				panic("not implemented")

				return nil
			},
		}

		cli.TagReplaceCmd = &cobra.Command{
			Use:   "replace MODEL...",
			Short: "Replace a bntp tag with an updated version",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.InFormat)
				if err != nil {
					return err
				}

				err = cli.BNTPBackend.TagManager.Replace(context.Background(), tags)

				return err
			},
		}

		cli.TagUpsertCmd = &cobra.Command{
			Use:   "upsert MODEL...",
			Short: "Add or replace a bntp tag",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.InFormat)
				if err != nil {
					return err
				}

				err = cli.BNTPBackend.TagManager.Upsert(context.Background(), tags)

				return err
			},
		}

		cli.TagEditCmd = &cobra.Command{
			Use:   "edit MODEL...",
			Short: "Edit a bntp tag",
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
				filter := &domain.TagFilter{}
				updater := &domain.TagUpdater{}
				var numAffectedRecordsRaw int64

				err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(updater, cli.UpdaterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				//*********************    Use provided tags    ********************//
				if cli.FilterRaw == "" {
					tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.InFormat)
					if err != nil {
						return err
					}

					err = cli.BNTPBackend.TagManager.Update(context.Background(), tags, updater)
					if err != nil {
						return err
					}

					numAffectedRecordsRaw = int64(len(args))

					//************************    Use filter    ************************//
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					numAffectedRecordsRaw, err = cli.BNTPBackend.TagManager.UpdateWhere(context.Background(), filter, updater)
					if err != nil {
						return err
					}
				}

				numAffectedRecords, err := cli.BNTPBackend.Marshallers[cli.InFormat].Marshall(NumAffectedRecords{numAffectedRecordsRaw})
				if err != nil {
					return err
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), numAffectedRecords)

				return nil
			},
		}

		cli.TagExportCmd = &cobra.Command{
			Use:   "export FILE",
			Short: "Export bntp tags",
			Long:  `A longer description`,
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				tags, err := cli.BNTPBackend.TagManager.GetAll(context.Background())
				if err != nil {
					return err
				}

				serializedTags, err := cli.BNTPBackend.Marshallers[cli.InFormat].Marshall(tags)
				if err != nil {
					return err
				}

				afero.WriteFile(cli.FsOverride, args[0], []byte(serializedTags), 0o644)

				return nil
			},
		}

		cli.TagImportCmd = &cobra.Command{
			Use:   "import FILE",
			Short: "Import bntp tags",
			Long:  `A longer description`,
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				serializedTags, err := afero.ReadFile(cli.FsOverride, args[0])
				if err != nil {
					return err
				}

				if len(serializedTags) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				var tags []*domain.Tag

				err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(&tags, string(serializedTags))
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				err = cli.BNTPBackend.TagManager.Add(context.Background(), tags)

				return err
			},
		}

		cli.TagListCmd = &cobra.Command{
			Use:   "list",
			Short: "List bntp tags",
			Long:  `A longer description`,
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				var tags []*domain.Tag
				filter := &domain.TagFilter{}
				var output string
				var err error

				//**************************    Get all    *************************//
				if cli.FilterRaw == "" {
					tags, err = cli.BNTPBackend.TagManager.GetAll(context.Background())
					if err != nil {
						return err
					}

					//********************    Use provided filter    *******************//
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					tags, err = cli.BNTPBackend.TagManager.GetWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				if cli.PathFormat || cli.ShortFormat {
					pathMarshaller := func(t *domain.Tag) (string, error) {
						return cli.BNTPBackend.TagManager.MarshalPath(context.Background(), t, cli.ShortFormat)
					}
					var paths []string

					paths, err = goaoi.TransformCopySlice(tags, pathMarshaller)
					if err != nil {
						return err
					}

					output = strings.Join(paths, "\n")
				} else {

					output, err = cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(tags)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), output)

				return nil
			},
		}

		cli.TagRemoveCmd = &cobra.Command{
			Use:   "remove TAG...",
			Short: "Remove bntp tags",
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
				filter := &domain.TagFilter{}
				var numAffectedRecordsRaw int64

				//********************    Use provided tags      *******************//
				if cli.FilterRaw == "" {
					tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.InFormat)
					if err != nil {
						return err
					}

					err = cli.BNTPBackend.TagManager.Delete(context.Background(), tags)
					if err != nil {
						return err
					}

					numAffectedRecordsRaw = int64(len(args))

					//********************       Use filter      *******************//
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					numAffectedRecordsRaw, err = cli.BNTPBackend.TagManager.DeleteWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				numAffectedRecords, err := cli.BNTPBackend.Marshallers[cli.InFormat].Marshall(NumAffectedRecords{numAffectedRecordsRaw})
				if err != nil {
					return err
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), numAffectedRecords)

				return nil
			},
		}

		cli.TagFindCmd = &cobra.Command{
			Use:   "find-first",
			Short: "Find the first tag matching a filter",
			Long:  `A longer description`,
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				filter := &domain.TagFilter{}
				var err error
				var result *domain.Tag
				var output string

				err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				result, err = cli.BNTPBackend.TagManager.GetFirstWhere(context.Background(), filter)
				if err != nil {
					return err
				}
				if cli.PathFormat || cli.ShortFormat {
					output, err = cli.BNTPBackend.TagManager.MarshalPath(context.Background(), result, cli.ShortFormat)

					if err != nil {
						return err
					}
				} else {

					output, err = cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(result)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), output)

				return nil
			},
		}

		cli.TagCountCmd = &cobra.Command{
			Use:   "count",
			Short: "Manage bntp tags",
			Long:  `A longer description`,
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				filter := &domain.TagFilter{}
				var countRaw int64
				var err error

				if cli.FilterRaw == "" {
					countRaw, err = cli.BNTPBackend.TagManager.CountAll(context.Background())
					if err != nil {
						return err
					}
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					countRaw, err = cli.BNTPBackend.TagManager.CountWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				count, err := cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(Count{countRaw})
				if err != nil {
					return err
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), count)

				return nil
			},
		}

		cli.TagDoesExistCmd = &cobra.Command{
			Use:   "does-exist [MODEL]",
			Short: "Manage bntp tags",
			Long:  `A longer description`,
			Args:  cobra.RangeArgs(0, 1),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 && cli.FilterRaw == "" {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}
				if len(args) > 0 && cli.FilterRaw != "" {
					return ConflictingPositionalArgsAndFlagError{Flag: "filter"}
				}

				filter := &domain.TagFilter{}
				var err error
				tag := &domain.Tag{}
				var doesExistRaw bool

				if cli.FilterRaw == "" {
					err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(tag, args[0])
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					doesExistRaw, err = cli.BNTPBackend.TagManager.DoesExist(context.Background(), tag)
					if err != nil {
						return err
					}
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.InFormat].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					doesExistRaw, err = cli.BNTPBackend.TagManager.DoesExistWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				doesExist, err := cli.BNTPBackend.Marshallers[cli.OutFormat].Marshall(DoesExist{doesExistRaw})
				if err != nil {
					return err
				}

				fmt.Fprintln(cli.RootCmd.OutOrStdout(), doesExist)

				return nil
			},
		}

		cli.TagShortCmd = &cobra.Command{
			Use:   "short TAG...",
			Short: "Return shortened bntp tags",
			Long:  `A longer description`,
			Args:  cobra.ArbitraryArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				panic("not implemented")

				return nil
			},
		}

		cli.RootCmd.AddCommand(cli.TagCmd)

		cli.TagCmd.AddCommand(cli.TagShortCmd)
		cli.TagCmd.AddCommand(cli.TagRemoveCmd)
		cli.TagCmd.AddCommand(cli.TagListCmd)
		cli.TagCmd.AddCommand(cli.TagImportCmd)
		cli.TagCmd.AddCommand(cli.TagExportCmd)
		cli.TagCmd.AddCommand(cli.TagEditCmd)
		cli.TagCmd.AddCommand(cli.TagUpsertCmd)
		cli.TagCmd.AddCommand(cli.TagReplaceCmd)
		cli.TagCmd.AddCommand(cli.TagAmbiguousCmd)
		cli.TagCmd.AddCommand(cli.TagCountCmd)
		cli.TagCmd.AddCommand(cli.TagFindCmd)
		cli.TagCmd.AddCommand(cli.TagDoesExistCmd)
		cli.TagCmd.AddCommand(cli.TagAddCmd)

		for _, subcommand := range cli.TagCmd.Commands() {
			// TODO: Should this flag be used for every command?
			if slices.Contains([]*cobra.Command{cli.TagAddCmd, cli.TagListCmd, cli.TagRemoveCmd, cli.TagFindCmd, cli.TagDoesExistCmd}, subcommand) {
				subcommand.PersistentFlags().StringVar(&cli.InFormat, "out-format", "json", "The serialization format to use for reading input")
				subcommand.PersistentFlags().StringVar(&cli.OutFormat, "in-format", "json", "The serialization format to use for writing output")
			}
		}

		for _, subcommand := range cli.TagCmd.Commands() {
			if slices.Contains([]*cobra.Command{cli.TagEditCmd, cli.TagListCmd, cli.TagRemoveCmd, cli.TagFindCmd, cli.TagCountCmd, cli.TagDoesExistCmd}, subcommand) {
				subcommand.PersistentFlags().StringVar(&cli.FilterRaw, "filter", "", "The filter to use for processing entities")
			}
		}

		for _, subcommand := range cli.TagCmd.Commands() {
			if slices.Contains([]*cobra.Command{cli.TagEditCmd}, subcommand) {
				subcommand.PersistentFlags().StringVar(&cli.UpdaterRaw, "updater", "", "The updater to use for processing entities")
			}
		}

		cli.TagFindCmd.MarkPersistentFlagRequired("filter")
		cli.TagEditCmd.MarkPersistentFlagRequired("updater")

		for _, subcommand := range cli.TagCmd.Commands() {
			if slices.Contains([]*cobra.Command{cli.TagListCmd, cli.TagFindCmd}, subcommand) {
				subcommand.PersistentFlags().BoolVar(&cli.PathFormat, "path-format", false, "Whetever to list tags in path format instead of --format format")
				subcommand.PersistentFlags().BoolVar(&cli.ShortFormat, "short-format", false, "Whetever to list tags in shortened path format instead of --format format")

				subcommand.MarkFlagsMutuallyExclusive("path-format", "short-format", "out-format")
			}
		}

	}
}
