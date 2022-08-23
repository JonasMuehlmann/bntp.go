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
				tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.Format)
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

				tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.Format)
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

				tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.Format)
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
				var filter *domain.TagFilter
				updater := &domain.TagUpdater{}
				var numAffectedRecords int64

				tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.Format)
				if err != nil {
					return err
				}

				err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(updater, cli.UpdaterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				//*********************    Use provided tags    ********************//
				if cli.FilterRaw == "" {
					err := cli.BNTPBackend.TagManager.Update(context.Background(), tags, updater)
					if err != nil {
						return err
					}

					numAffectedRecords = int64(len(args))

					//************************    Use filter    ************************//
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					numAffectedRecords, err = cli.BNTPBackend.TagManager.UpdateWhere(context.Background(), filter, updater)
					if err != nil {
						return err
					}
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

				serializedTags, err := cli.BNTPBackend.Marshallers[cli.Format].Marshall(tags)
				if err != nil {
					return err
				}

				afero.WriteFile(cli.Fs, args[0], []byte(serializedTags), 0o644)

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

				serializedTags, err := afero.ReadFile(cli.Fs, args[0])
				if err != nil {
					return err
				}

				if len(serializedTags) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				var tags []*domain.Tag

				err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(&tags, string(serializedTags))
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
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				var tags []*domain.Tag
				var filter *domain.TagFilter
				var output string
				var err error

				if cli.FilterRaw == "" {
					tags, err = cli.BNTPBackend.TagManager.GetAll(context.Background())
					if err != nil {
						return err
					}
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					tags, err = cli.BNTPBackend.TagManager.GetWhere(context.Background(), filter)
					if err != nil {
						return err
					}
				}

				output, err = cli.BNTPBackend.Marshallers[cli.Format].Marshall(tags)
				if err != nil {
					return EntityMarshallingError{Inner: err}
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
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				var filter *domain.TagFilter
				var err error
				var numAffectedRecords int64

				if cli.FilterRaw == "" {
					tags, err := UnmarshalEntities[domain.Tag](cli, args, cli.Format)
					if err != nil {
						return err
					}

					err = cli.BNTPBackend.TagManager.Delete(context.Background(), tags)
					if err != nil {
						return err
					}

					numAffectedRecords = int64(len(args))
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					numAffectedRecords, err = cli.BNTPBackend.TagManager.DeleteWhere(context.Background(), filter)
					if err != nil {
						return err
					}
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
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				var filter *domain.TagFilter
				var err error
				var result *domain.Tag
				var output string

				err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				result, err = cli.BNTPBackend.TagManager.GetFirstWhere(context.Background(), filter)
				if err != nil {
					return err
				}

				output, err = cli.BNTPBackend.Marshallers[cli.Format].Marshall(result)
				if err != nil {
					return EntityMarshallingError{Inner: err}
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
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				var filter *domain.TagFilter
				var count int64
				var err error

				if cli.FilterRaw == "" {
					count, err = cli.BNTPBackend.TagManager.CountAll(context.Background())
					if err != nil {
						return err
					}
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					count, err = cli.BNTPBackend.TagManager.CountWhere(context.Background(), filter)
					if err != nil {
						return err
					}
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
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				var filter *domain.TagFilter
				var err error
				var tag *domain.Tag
				var doesExist bool

				if cli.FilterRaw == "" {
					err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(tag, args[0])
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					doesExist, err = cli.BNTPBackend.TagManager.DoesExist(context.Background(), tag)
					if err != nil {
						return err
					}
				} else {
					err = cli.BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
					if err != nil {
						return EntityMarshallingError{Inner: err}
					}

					doesExist, err = cli.BNTPBackend.TagManager.DoesExistWhere(context.Background(), filter)
					if err != nil {
						return err
					}
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
			if slices.Contains([]*cobra.Command{cli.TagAddCmd, cli.TagListCmd, cli.TagRemoveCmd}, subcommand) {
				subcommand.PersistentFlags().StringVar(&cli.Format, "format", "json", "The serialization format to use for i/o")
			}
		}

		for _, subcommand := range cli.TagCmd.Commands() {
			if slices.Contains([]*cobra.Command{cli.TagEditCmd, cli.TagListCmd, cli.TagRemoveCmd}, subcommand) {
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

		cli.TagListCmd.Flags().BoolP("short", "s", false, "Whetever to list shortened tags instead of fully qualified ones")
	}
}
