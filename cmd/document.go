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

func WithDocument(cli *Cli) {
	cli.DocumentCmd = &cobra.Command{
		Use:   "document",
		Short: "Manage bntp documents",
		Long:  `A longer description`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			return nil
		},
	}

	cli.DocumentAddCmd = &cobra.Command{
		Use:   "add MODEL...",
		Short: "Add bntp documents",
		Long:  `A longer description`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			documents, err := UnmarshalEntities[domain.Document](args, cli.Format)
			if err != nil {
				return err
			}

			// NOTE: Should we also try to add an empty document?
			err = BNTPBackend.DocumentManager.Add(context.Background(), documents)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cli.DocumentReplaceCmd = &cobra.Command{
		Use:   "replace MODEL...",
		Short: "Replace a bntp document with an updated version",
		Long:  `A longer description`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			documents, err := UnmarshalEntities[domain.Document](args, cli.Format)
			if err != nil {
				return err
			}

			err = BNTPBackend.DocumentManager.Replace(context.Background(), documents)
			if err != nil {
				return err
			}

			err = BNTPBackend.DocumentContentManager.UpdateDocumentContentsFromNewModels(context.Background(), documents, &BNTPBackend.DocumentManager)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cli.DocumentUpsertCmd = &cobra.Command{
		Use:   "upsert MODEL...",
		Short: "Add or replace a bntp document",
		Long:  `A longer description`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			documents, err := UnmarshalEntities[domain.Document](args, cli.Format)
			if err != nil {
				return err
			}

			err = BNTPBackend.DocumentManager.Upsert(context.Background(), documents)
			if err != nil {
				return err
			}

			err = BNTPBackend.DocumentContentManager.UpdateDocumentContentsFromNewModels(context.Background(), documents, &BNTPBackend.DocumentManager)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cli.DocumentEditCmd = &cobra.Command{
		Use:   "edit MODEL...",
		Short: "Edit a bntp document",
		Long:  `A longer description`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			var err error
			var filter *domain.DocumentFilter
			var updater *domain.DocumentUpdater
			var numAffectedRecords int64

			documents, err := UnmarshalEntities[domain.Document](args, cli.Format)
			if err != nil {
				return err
			}

			err = BNTPBackend.Unmarshallers[cli.Format].Unmarshall(updater, cli.UpdaterRaw)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}
			if cli.FilterRaw == "" {
				err := BNTPBackend.DocumentManager.Update(context.Background(), documents, updater)
				if err != nil {
					return err
				}

				numAffectedRecords = int64(len(args))

				err = BNTPBackend.DocumentContentManager.UpdateDocumentContentsFromFilterAndUpdater(context.Background(), filter, updater, &BNTPBackend.DocumentManager)
				if err != nil {
					return err
				}
			} else {
				err = BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				numAffectedRecords, err = BNTPBackend.DocumentManager.UpdateWhere(context.Background(), filter, updater)
				if err != nil {
					return err
				}

				err = BNTPBackend.DocumentContentManager.UpdateDocumentContentsFromNewModels(context.Background(), documents, &BNTPBackend.DocumentManager)
				if err != nil {
					return err
				}
			}

			fmt.Fprintln(cli.RootCmd.OutOrStdout(), numAffectedRecords)

			return nil
		},
	}
	cli.DocumentListCmd = &cobra.Command{
		Use:   "list",
		Short: "List bntp documents",
		Long:  `A longer description`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			var documents []*domain.Document
			var filter *domain.DocumentFilter
			var output string
			var err error

			if cli.FilterRaw == "" {
				documents, err = BNTPBackend.DocumentManager.GetAll(context.Background())
				if err != nil {
					return err
				}
			} else {
				err = BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				documents, err = BNTPBackend.DocumentManager.GetWhere(context.Background(), filter)
				if err != nil {
					return err
				}
			}

			output, err = BNTPBackend.Marshallers[cli.Format].Marshall(documents)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			fmt.Fprintln(cli.RootCmd.OutOrStdout(), output)

			return nil
		},
	}

	cli.DocumentRemoveCmd = &cobra.Command{
		Use:   "remove [MODEL...]",
		Short: "Remove a bntp document",
		Long:  `A longer description`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			var filter *domain.DocumentFilter
			var err error
			var numAffectedRecords int64

			if cli.FilterRaw == "" {
				documents, err := UnmarshalEntities[domain.Document](args, cli.Format)
				if err != nil {
					return err
				}

				err = BNTPBackend.DocumentManager.Delete(context.Background(), documents)
				if err != nil {
					return err
				}

				numAffectedRecords = int64(len(args))
			} else {
				err = BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				numAffectedRecords, err = BNTPBackend.DocumentManager.DeleteWhere(context.Background(), filter)
				if err != nil {
					return err
				}
			}

			fmt.Fprintln(cli.RootCmd.OutOrStdout(), numAffectedRecords)

			return nil
		},
	}

	cli.DocumentFindCmd = &cobra.Command{
		Use:   "find-first",
		Short: "Find the first document matching a filter",
		Long:  `A longer description`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			var filter *domain.DocumentFilter
			var err error
			var result *domain.Document
			var output string

			err = BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			result, err = BNTPBackend.DocumentManager.GetFirstWhere(context.Background(), filter)
			if err != nil {
				return err
			}

			output, err = BNTPBackend.Marshallers[cli.Format].Marshall(result)
			if err != nil {
				return EntityMarshallingError{Inner: err}
			}

			fmt.Fprintln(cli.RootCmd.OutOrStdout(), output)

			return nil
		},
	}
	cli.DocumentCountCmd = &cobra.Command{
		Use:   "count",
		Short: "Manage bntp documents",
		Long:  `A longer description`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			var filter *domain.DocumentFilter
			var count int64
			var err error

			if cli.FilterRaw == "" {
				count, err = BNTPBackend.DocumentManager.CountAll(context.Background())
				if err != nil {
					return err
				}
			} else {
				err = BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				count, err = BNTPBackend.DocumentManager.CountWhere(context.Background(), filter)
				if err != nil {
					return err
				}
			}

			fmt.Fprintln(cli.RootCmd.OutOrStdout(), count)

			return nil
		},
	}
	cli.DocumentDoesExistCmd = &cobra.Command{
		Use:   "does-exist [MODEL]",
		Short: "Manage bntp documents",
		Long:  `A longer description`,
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			var filter *domain.DocumentFilter
			var err error
			var document *domain.Document
			var doesExist bool

			if cli.FilterRaw == "" {
				err = BNTPBackend.Unmarshallers[cli.Format].Unmarshall(document, args[0])
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				doesExist, err = BNTPBackend.DocumentManager.DoesExist(context.Background(), document)
				if err != nil {
					return err
				}
			} else {
				err = BNTPBackend.Unmarshallers[cli.Format].Unmarshall(filter, cli.FilterRaw)
				if err != nil {
					return EntityMarshallingError{Inner: err}
				}

				doesExist, err = BNTPBackend.DocumentManager.DoesExistWhere(context.Background(), filter)
				if err != nil {
					return err
				}
			}

			fmt.Fprintln(cli.RootCmd.OutOrStdout(), doesExist)

			return nil
		},
	}

	cli.RootCmd.AddCommand(cli.DocumentCmd)

	cli.DocumentCmd.AddCommand(cli.DocumentAddCmd)
	cli.DocumentCmd.AddCommand(cli.DocumentEditCmd)
	cli.DocumentCmd.AddCommand(cli.DocumentListCmd)
	cli.DocumentCmd.AddCommand(cli.DocumentRemoveCmd)
	cli.DocumentCmd.AddCommand(cli.DocumentCountCmd)
	cli.DocumentCmd.AddCommand(cli.DocumentDoesExistCmd)
	cli.DocumentCmd.AddCommand(cli.DocumentFindCmd)
	cli.DocumentCmd.AddCommand(cli.DocumentUpsertCmd)

	cli.DocumentFindCmd.MarkFlagRequired("filter")
	cli.DocumentEditCmd.MarkFlagRequired("updater")

	for _, subcommand := range cli.DocumentCmd.Commands() {
		if slices.Contains([]*cobra.Command{cli.DocumentAddCmd, cli.DocumentListCmd, cli.DocumentRemoveCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&cli.Format, "format", "json", "The serialization format to use for i/o")
		}
	}

	for _, subcommand := range cli.DocumentCmd.Commands() {
		if slices.Contains([]*cobra.Command{cli.DocumentEditCmd, cli.DocumentListCmd, cli.DocumentRemoveCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&cli.FilterRaw, "filter", "", "The filter to use for processing entities")
		}
	}

	for _, subcommand := range cli.DocumentCmd.Commands() {
		if slices.Contains([]*cobra.Command{cli.DocumentEditCmd}, subcommand) {
			subcommand.PersistentFlags().StringVar(&cli.FilterRaw, "updater", "", "The updater to use for processing entities")
		}
	}

	// TODO: Add flag to not update document content, because they already were added through an editor
}
