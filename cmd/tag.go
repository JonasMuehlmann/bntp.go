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

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags available for bntp entities",
	Long:  `A longer description`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		return nil
	},
}

var tagAddCmd = &cobra.Command{
	Use:   "add TAG...",
	Short: "Add new bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}
		tags := make([]*domain.Tag, len(args))

		for i, arg := range args {
			tags[i] = new(domain.Tag)

			err := BNTPBackend.Unmarshallers[Format].Unmarshall(tags[i], arg)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}
		}

		err := BNTPBackend.TagManager.Add(context.Background(), tags)
		if err != nil {
			return err
		}

		return nil
	},
}

// TODO: If ambiguous, return ambiguous component
var tagAmbiguousCmd = &cobra.Command{
	Use:   "ambiguous TAG...",
	Short: "Check if bntp tag's leafs are ambiguous",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		return nil
	},
}

var tagReplaceCmd = &cobra.Command{
	Use:   "replace MODEL...",
	Short: "Replace a bntp tag with an updated version",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		tags := make([]*domain.Tag, 0, len(args))

		for i, tagOut := range tags {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(tagOut, args[i])
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}
		}

		err := BNTPBackend.TagManager.Replace(context.Background(), tags)
		if err != nil {
			return err
		}

		return nil
	},
}

var tagUpsertCmd = &cobra.Command{
	Use:   "upsert MODEL...",
	Short: "Add or replace a bntp tag",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		tags := make([]*domain.Tag, 0, len(args))

		for i, tagOut := range tags {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(tagOut, args[i])
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}
		}

		err := BNTPBackend.TagManager.Upsert(context.Background(), tags)
		if err != nil {
			return err
		}

		return nil
	},
}

var tagEditCmd = &cobra.Command{
	Use:   "edit MODEL...",
	Short: "Edit a bntp tag",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var err error
		var filter *domain.TagFilter
		var updater *domain.TagUpdater
		var numAffectedRecords int64

		tags := make([]*domain.Tag, 0, len(args))

		for i, tagOut := range tags {
			err := BNTPBackend.Unmarshallers[Format].Unmarshall(tagOut, args[i])
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}
		}

		err = BNTPBackend.Unmarshallers[Format].Unmarshall(updater, UpdaterRaw)
		if err != nil {
			return EntityMarshallingError{Inner: err}
		}
		if FilterRaw == "" {
			err := BNTPBackend.TagManager.Update(context.Background(), tags, updater)
			if err != nil {
				return err
			}

			numAffectedRecords = int64(len(args))
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			numAffectedRecords, err = BNTPBackend.TagManager.UpdateWhere(context.Background(), filter, updater)
			if err != nil {
				return err
			}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), numAffectedRecords)

		return nil
	},
}

var tagExportCmd = &cobra.Command{
	Use:   "export FILE",
	Short: "Export bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		return nil
	},
}

var tagImportCmd = &cobra.Command{
	Use:   "import FILE",
	Short: "Import bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		return nil
	},
}

var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bntp tags",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var tags []*domain.Tag
		var filter *domain.TagFilter
		var output string
		var err error

		if FilterRaw == "" {
			tags, err = BNTPBackend.TagManager.GetAll(context.Background())
			if err != nil {
				return err
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			tags, err = BNTPBackend.TagManager.GetWhere(context.Background(), filter)
			if err != nil {
				return err
			}
		}

		output, err = BNTPBackend.Marshallers[Format].Marshall(tags)
		if err != nil {
			return EntityMarshallingError{Inner: err}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), output)

		return nil
	},
}

var tagRemoveCmd = &cobra.Command{
	Use:   "remove TAG...",
	Short: "Remove bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var filter *domain.TagFilter
		var err error
		var numAffectedRecords int64

		if FilterRaw == "" {
			tags := make([]*domain.Tag, 0, len(args))

			for i, tagOut := range tags {
				err := BNTPBackend.Unmarshallers[Format].Unmarshall(tagOut, args[i])
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}
			}

			err = BNTPBackend.TagManager.Delete(context.Background(), tags)
			if err != nil {
				return err
			}

			numAffectedRecords = int64(len(args))
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			numAffectedRecords, err = BNTPBackend.TagManager.DeleteWhere(context.Background(), filter)
			if err != nil {
				return err
			}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), numAffectedRecords)

		return nil
	},
}

var tagFindCmd = &cobra.Command{
	Use:   "find-first",
	Short: "Find the first tag matching a filter",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var filter *domain.TagFilter
		var err error
		var result *domain.Tag
		var output string

		err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
		if err != nil {
			return EntityMarshallingError{Inner: err}
		}

		result, err = BNTPBackend.TagManager.GetFirstWhere(context.Background(), filter)
		if err != nil {
			return err
		}

		output, err = BNTPBackend.Marshallers[Format].Marshall(result)
		if err != nil {
			return EntityMarshallingError{Inner: err}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), output)

		return nil
	},
}

var tagCountCmd = &cobra.Command{
	Use:   "count",
	Short: "Manage bntp tags",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var filter *domain.TagFilter
		var count int64
		var err error

		if FilterRaw == "" {
			count, err = BNTPBackend.TagManager.CountAll(context.Background())
			if err != nil {
				return err
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			count, err = BNTPBackend.TagManager.CountWhere(context.Background(), filter)
			if err != nil {
				return err
			}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), count)

		return nil
	},
}

var tagDoesExistCmd = &cobra.Command{
	Use:   "does-exist [MODEL]",
	Short: "Manage bntp tags",
	Long:  `A longer description`,
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		var filter *domain.TagFilter
		var err error
		var tag *domain.Tag
		var doesExist bool

		if FilterRaw == "" {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(tag, args[0])
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			doesExist, err = BNTPBackend.TagManager.DoesExist(context.Background(), tag)
			if err != nil {
				return err
			}
		} else {
			err = BNTPBackend.Unmarshallers[Format].Unmarshall(filter, FilterRaw)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			doesExist, err = BNTPBackend.TagManager.DoesExistWhere(context.Background(), filter)
			if err != nil {
				return err
			}
		}

		fmt.Fprintln(RootCmd.OutOrStdout(), doesExist)

		return nil
	},
}

var tagShortCmd = &cobra.Command{
	Use:   "short TAG...",
	Short: "Return shortened bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(tagCmd)

	tagCmd.AddCommand(tagShortCmd)
	tagCmd.AddCommand(tagRemoveCmd)
	tagCmd.AddCommand(tagListCmd)
	tagCmd.AddCommand(tagImportCmd)
	tagCmd.AddCommand(tagExportCmd)
	tagCmd.AddCommand(tagEditCmd)
	tagCmd.AddCommand(tagUpsertCmd)
	tagCmd.AddCommand(tagReplaceCmd)
	tagCmd.AddCommand(tagAmbiguousCmd)
	tagCmd.AddCommand(tagCountCmd)
	tagCmd.AddCommand(tagFindCmd)
	tagCmd.AddCommand(tagDoesExistCmd)
	tagCmd.AddCommand(tagAddCmd)

	tagFindCmd.MarkFlagRequired("filter")
	tagEditCmd.MarkFlagRequired("updater")

	for _, subcommand := range tagCmd.Commands() {
		if slices.Contains([]*cobra.Command{tagAddCmd, tagListCmd, tagRemoveCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&Format, "format", "json", "The serialization format to use for i/o")
		}
	}

	for _, subcommand := range tagCmd.Commands() {
		if slices.Contains([]*cobra.Command{tagEditCmd, tagListCmd, tagRemoveCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&FilterRaw, "filter", "", "The filter to use for processing entities")
		}
	}

	for _, subcommand := range tagCmd.Commands() {
		if slices.Contains([]*cobra.Command{tagEditCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&FilterRaw, "updater", "", "The updater to use for processing entities")
		}
	}

	tagListCmd.Flags().BoolP("short", "s", false, "Whetever to list shortened tags instead of fully qualified ones")
}
