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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/filter_converter.go.tpl

package repository

import (
	"context"
	"database/sql"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	repoCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/volatiletech/null/v8"
	"time"
)

func BookmarkDomainToSqlRepositoryFilter(ctx context.Context, db *sql.DB, domainFilter *domain.BookmarkFilter) (sqlRepositoryFilter *BookmarkFilter, err error) {
	sqlRepositoryFilter = new(BookmarkFilter)

	sqlRepositoryFilter.URL = domainFilter.URL
	sqlRepositoryFilter.ID = domainFilter.ID

	//**********************    Set Timestamps    **********************//

	sqlRepositoryFilter.CreatedAt = domainFilter.CreatedAt
	sqlRepositoryFilter.UpdatedAt = domainFilter.UpdatedAt

	if domainFilter.DeletedAt.HasValue {
		var convertedFilter model.FilterOperation[null.Time]

		convertedFilter, err = model.ConvertFilter[null.Time, optional.Optional[time.Time]](domainFilter.DeletedAt.Wrappee, repoCommon.OptionalTimeToNullTime)
		if err != nil {
			return
		}

		sqlRepositoryFilter.DeletedAt.Push(convertedFilter)
	}

	//*************************    Set Title    ************************//
	if domainFilter.Title.HasValue {
		var convertedFilter model.FilterOperation[null.String]

		convertedFilter, err = model.ConvertFilter[null.String, optional.Optional[string]](domainFilter.Title.Wrappee, repoCommon.OptionalStringToNullString)
		if err != nil {
			return
		}

		sqlRepositoryFilter.Title.Push(convertedFilter)
	}

	//******************    Set IsRead/IsCollection    *****************//
	if domainFilter.IsRead.HasValue {
		var convertedFilter model.FilterOperation[int64]

		convertedFilter, err = model.ConvertFilter[int64, bool](domainFilter.IsRead.Wrappee, repoCommon.BoolToInt)
		if err != nil {
			return
		}

		sqlRepositoryFilter.IsRead.Push(convertedFilter)
	}

	if domainFilter.IsCollection.HasValue {
		var convertedFilter model.FilterOperation[int64]

		convertedFilter, err = model.ConvertFilter[int64, bool](domainFilter.IsCollection.Wrappee, repoCommon.BoolToInt)
		if err != nil {
			return
		}

		sqlRepositoryFilter.IsCollection.Push(convertedFilter)
	}

	//*************************    Set Tags    *************************//

	if domainFilter.Tags.HasValue {
		var convertedFilter model.FilterOperation[*Tag]

		convertedFilter, err = model.ConvertFilter[*Tag, *domain.Tag](domainFilter.Tags.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[domain.Tag, Tag](ctx, db, TagDomainToSqlRepositoryModel))
		if err != nil {
			return
		}

		sqlRepositoryFilter.Tags.Push(convertedFilter)
	}

	//*************************    Set Type    *************************//

	if domainFilter.BookmarkType.HasValue {
		var convertedTypeIDFilter model.FilterOperation[null.Int64]
		var convertedTypeFilter model.FilterOperation[*BookmarkType]

		convertedTypeFilter, err = model.ConvertFilter[*BookmarkType, optional.Optional[string]](domainFilter.BookmarkType.Wrappee, func(type_ optional.Optional[string]) (*BookmarkType, error) {
			if !type_.HasValue {
				return nil, nil
			}

			bookmarkType, err := BookmarkTypes(BookmarkTypeWhere.Type.EQ(type_.Wrappee)).One(ctx, db)

			return bookmarkType, err
		})
		if err != nil {
			return
		}

		convertedTypeIDFilter, err = model.ConvertFilter[null.Int64, optional.Optional[string]](domainFilter.BookmarkType.Wrappee, func(type_ optional.Optional[string]) (null.Int64, error) {
			if !type_.HasValue {
				return null.NewInt64(-1, false), nil
			}

			bookmarkType, err := BookmarkTypes(BookmarkTypeWhere.Type.EQ(type_.Wrappee)).One(ctx, db)

			return null.NewInt64(bookmarkType.ID, true), err
		})
		if err != nil {
			return
		}

		sqlRepositoryFilter.BookmarkType.Push(convertedTypeFilter)
		sqlRepositoryFilter.BookmarkTypeID.Push(convertedTypeIDFilter)
	}

	return
}

func BookmarkSqlRepositoryToDomainFilter(ctx context.Context, db *sql.DB, sqlRepositoryFilter *BookmarkFilter) (domainFilter *domain.BookmarkFilter, err error) {
	domainFilter = new(domain.BookmarkFilter)

	domainFilter.URL = sqlRepositoryFilter.URL
	domainFilter.ID = sqlRepositoryFilter.ID

	//*************************    Set Type    *************************//
	if sqlRepositoryFilter.BookmarkTypeID.HasValue {
		var convertedFilter model.FilterOperation[optional.Optional[string]]

		convertedFilter, err = model.ConvertFilter[optional.Optional[string], null.Int64](sqlRepositoryFilter.BookmarkTypeID.Wrappee, func(typeID null.Int64) (optional.Optional[string], error) {
			if typeID.Valid {
				bookmarkType, err := BookmarkTypes(BookmarkTypeWhere.ID.EQ(typeID.Int64)).One(ctx, db)
				if err != nil {
					return optional.Optional[string]{}, err
				}

				return optional.Make(bookmarkType.Type), nil
			}

			return optional.Optional[string]{}, nil
		})
		if err != nil {
			return
		}

		domainFilter.BookmarkType.Push(convertedFilter)
	}

	if sqlRepositoryFilter.BookmarkTypeID.HasValue {
		var convertedFilter model.FilterOperation[optional.Optional[string]]

		convertedFilter, err = model.ConvertFilter[optional.Optional[string], *BookmarkType](sqlRepositoryFilter.BookmarkType.Wrappee, func(type_ *BookmarkType) (optional.Optional[string], error) {
			if type_ != nil {
				return optional.Make(type_.Type), nil
			}

			return optional.Optional[string]{}, nil
		})
		if err != nil {
			return
		}

		domainFilter.BookmarkType.Push(convertedFilter)
	}

	//**********************    Set Timestamps    **********************//

	domainFilter.CreatedAt = sqlRepositoryFilter.CreatedAt
	domainFilter.UpdatedAt = sqlRepositoryFilter.UpdatedAt

	if sqlRepositoryFilter.DeletedAt.HasValue {
		var convertedFilter model.FilterOperation[optional.Optional[time.Time]]

		convertedFilter, err = model.ConvertFilter[optional.Optional[time.Time], null.Time](sqlRepositoryFilter.DeletedAt.Wrappee, repoCommon.NullStringToOptionalTime)
		if err != nil {
			return
		}

		domainFilter.DeletedAt.Push(convertedFilter)
	}

	//*************************    Set Title    ************************//
	if sqlRepositoryFilter.Title.HasValue {
		var convertedFilter model.FilterOperation[optional.Optional[string]]

		convertedFilter, err = model.ConvertFilter[optional.Optional[string], null.String](sqlRepositoryFilter.Title.Wrappee, repoCommon.NullStringToOptionalString)
		if err != nil {
			return
		}

		domainFilter.Title.Push(convertedFilter)
	}

	//******************    Set IsRead/IsCollection    *****************//
	if sqlRepositoryFilter.IsRead.HasValue {
		var convertedFilter model.FilterOperation[bool]

		convertedFilter, err = model.ConvertFilter[bool, int64](sqlRepositoryFilter.IsRead.Wrappee, repoCommon.IntToBool)
		if err != nil {
			return
		}

		domainFilter.IsRead.Push(convertedFilter)
	}

	if sqlRepositoryFilter.IsCollection.HasValue {
		var convertedFilter model.FilterOperation[bool]

		convertedFilter, err = model.ConvertFilter[bool, int64](sqlRepositoryFilter.IsCollection.Wrappee, repoCommon.IntToBool)
		if err != nil {
			return
		}

		domainFilter.IsCollection.Push(convertedFilter)
	}

	//*************************    Set Tags    *************************//
	if sqlRepositoryFilter.Tags.HasValue {
		var convertedFilter model.FilterOperation[*domain.Tag]

		convertedFilter, err = model.ConvertFilter[*domain.Tag, *Tag](sqlRepositoryFilter.Tags.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[Tag, domain.Tag](ctx, db, TagSqlRepositoryToDomainModel))
		if err != nil {
			return
		}

		domainFilter.Tags.Push(convertedFilter)
	}

	return
}

func DocumentDomainToSqlRepositoryFilter(ctx context.Context, db *sql.DB, domainFilter *domain.DocumentFilter) (sqlRepositoryFilter *DocumentFilter, err error) {
	sqlRepositoryFilter = new(DocumentFilter)

	sqlRepositoryFilter.Path = domainFilter.Path
	sqlRepositoryFilter.ID = domainFilter.ID

	//**********************    Set Timestamps    **********************//

	sqlRepositoryFilter.CreatedAt = domainFilter.CreatedAt
	sqlRepositoryFilter.UpdatedAt = domainFilter.UpdatedAt

	if domainFilter.DeletedAt.HasValue {
		var convertedFilter model.FilterOperation[null.Time]

		convertedFilter, err = model.ConvertFilter[null.Time, optional.Optional[time.Time]](domainFilter.DeletedAt.Wrappee, repoCommon.OptionalTimeToNullTime)
		if err != nil {
			return
		}

		sqlRepositoryFilter.DeletedAt.Push(convertedFilter)
	}

	//*************************    Set Tags    *************************//
	if domainFilter.Tags.HasValue {
		var convertedFilter model.FilterOperation[*Tag]

		convertedFilter, err = model.ConvertFilter[*Tag, *domain.Tag](domainFilter.Tags.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[domain.Tag, Tag](ctx, db, TagDomainToSqlRepositoryModel))
		if err != nil {
			return
		}

		sqlRepositoryFilter.Tags.Push(convertedFilter)
	}

	//*************************    Set Type    *************************//
	if domainFilter.DocumentType.HasValue {
		var convertedTypeIDFilter model.FilterOperation[null.Int64]
		var convertedTypeFilter model.FilterOperation[*DocumentType]

		convertedTypeFilter, err = model.ConvertFilter[*DocumentType, optional.Optional[string]](domainFilter.DocumentType.Wrappee, func(type_ optional.Optional[string]) (*DocumentType, error) {
			if !type_.HasValue {
				return nil, nil
			}

			bookmarkType, err := DocumentTypes(DocumentTypeWhere.DocumentType.EQ(type_.Wrappee)).One(ctx, db)

			return bookmarkType, err
		})
		if err != nil {
			return
		}

		convertedTypeIDFilter, err = model.ConvertFilter[null.Int64, optional.Optional[string]](domainFilter.DocumentType.Wrappee, func(type_ optional.Optional[string]) (null.Int64, error) {
			if !type_.HasValue {
				return null.NewInt64(-1, false), nil
			}

			bookmarkType, err := DocumentTypes(DocumentTypeWhere.DocumentType.EQ(type_.Wrappee)).One(ctx, db)

			return null.NewInt64(bookmarkType.ID, true), err
		})
		if err != nil {
			return
		}

		sqlRepositoryFilter.DocumentType.Push(convertedTypeFilter)
		sqlRepositoryFilter.DocumentTypeID.Push(convertedTypeIDFilter)
	}

	//**************    Set linked/backlinked documents    *************//
	if domainFilter.LinkedDocuments.HasValue {
		var convertedFilter model.FilterOperation[*Document]

		convertedFilter, err = model.ConvertFilter[*Document, *domain.Document](domainFilter.LinkedDocuments.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[domain.Document, Document](ctx, db, DocumentDomainToSqlRepositoryModel))
		if err != nil {
			return
		}

		sqlRepositoryFilter.SourceDocuments.Push(convertedFilter)
	}
	if domainFilter.BacklinkedDocuments.HasValue {
		var convertedFilter model.FilterOperation[*Document]

		convertedFilter, err = model.ConvertFilter[*Document, *domain.Document](domainFilter.BacklinkedDocuments.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[domain.Document, Document](ctx, db, DocumentDomainToSqlRepositoryModel))
		if err != nil {
			return
		}

		sqlRepositoryFilter.DestinationDocuments.Push(convertedFilter)
	}

	return
}

func DocumentSqlRepositoryToDomainFilter(ctx context.Context, db *sql.DB, sqlRepositoryFilter *DocumentFilter) (domainFilter *domain.DocumentFilter, err error) {
	domainFilter = new(domain.DocumentFilter)

	domainFilter.Path = sqlRepositoryFilter.Path
	domainFilter.ID = sqlRepositoryFilter.ID

	//*************************    Set Type    *************************//
	if sqlRepositoryFilter.DocumentTypeID.HasValue {
		var convertedFilter model.FilterOperation[optional.Optional[string]]

		convertedFilter, err = model.ConvertFilter[optional.Optional[string], null.Int64](sqlRepositoryFilter.DocumentTypeID.Wrappee, func(typeID null.Int64) (optional.Optional[string], error) {
			if typeID.Valid {
				documentType, err := DocumentTypes(DocumentTypeWhere.ID.EQ(typeID.Int64)).One(ctx, db)
				if err != nil {
					return optional.Optional[string]{}, err
				}

				return optional.Make(documentType.DocumentType), nil
			}

			return optional.Optional[string]{}, nil
		})
		if err != nil {
			return
		}

		domainFilter.DocumentType.Push(convertedFilter)
	}

	if sqlRepositoryFilter.DocumentTypeID.HasValue {
		var convertedFilter model.FilterOperation[optional.Optional[string]]

		convertedFilter, err = model.ConvertFilter[optional.Optional[string], *DocumentType](sqlRepositoryFilter.DocumentType.Wrappee, func(type_ *DocumentType) (optional.Optional[string], error) {
			if type_ != nil {
				return optional.Make(type_.DocumentType), nil
			}

			return optional.Optional[string]{}, nil
		})
		if err != nil {
			return
		}

		domainFilter.DocumentType.Push(convertedFilter)
	}

	//**********************    Set Timestamps    **********************//

	domainFilter.CreatedAt = sqlRepositoryFilter.CreatedAt
	domainFilter.UpdatedAt = sqlRepositoryFilter.UpdatedAt

	if sqlRepositoryFilter.DeletedAt.HasValue {
		var convertedFilter model.FilterOperation[optional.Optional[time.Time]]

		convertedFilter, err = model.ConvertFilter[optional.Optional[time.Time], null.Time](sqlRepositoryFilter.DeletedAt.Wrappee, repoCommon.NullStringToOptionalTime)
		if err != nil {
			return
		}

		domainFilter.DeletedAt.Push(convertedFilter)
	}

	//*************************    Set Tags    *************************//
	if sqlRepositoryFilter.Tags.HasValue {
		var convertedFilter model.FilterOperation[*domain.Tag]

		convertedFilter, err = model.ConvertFilter[*domain.Tag, *Tag](sqlRepositoryFilter.Tags.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[Tag, domain.Tag](ctx, db, TagSqlRepositoryToDomainModel))
		if err != nil {
			return
		}

		domainFilter.Tags.Push(convertedFilter)
	}

	//**************    Set linked/backlinked documents    *************//
	if sqlRepositoryFilter.SourceDocuments.HasValue {
		var convertedFilter model.FilterOperation[*domain.Document]

		convertedFilter, err = model.ConvertFilter[*domain.Document, *Document](sqlRepositoryFilter.SourceDocuments.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[Document, domain.Document](ctx, db, DocumentSqlRepositoryToDomainModel))
		if err != nil {
			return
		}

		domainFilter.LinkedDocuments.Push(convertedFilter)
	}

	if sqlRepositoryFilter.DestinationDocuments.HasValue {
		var convertedFilter model.FilterOperation[*domain.Document]

		convertedFilter, err = model.ConvertFilter[*domain.Document, *Document](sqlRepositoryFilter.DestinationDocuments.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[Document, domain.Document](ctx, db, DocumentSqlRepositoryToDomainModel))
		if err != nil {
			return
		}

		domainFilter.BacklinkedDocuments.Push(convertedFilter)
	}

	return
}

func TagDomainToSqlRepositoryFilter(ctx context.Context, db *sql.DB, domainFilter *domain.TagFilter) (sqlRepositoryFilter *TagFilter, err error) {
	sqlRepositoryFilter = new(TagFilter)

	sqlRepositoryFilter.ID = domainFilter.ID
	sqlRepositoryFilter.Tag = domainFilter.Tag

	//**********************    Set parent path    *********************//
	if domainFilter.ParentPath.HasValue {
		var convertedFilter model.FilterOperation[*Tag]

		convertedFilter, err = model.ConvertFilter[*Tag, *domain.Tag](domainFilter.ParentPath.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[domain.Tag, Tag](ctx, db, TagDomainToSqlRepositoryModel))
		if err != nil {
			return
		}

		sqlRepositoryFilter.ParentTagTags.Push(convertedFilter)
	}

	//**********************    Set child tags *********************//
	if domainFilter.Subtags.HasValue {
		var convertedFilter model.FilterOperation[*Tag]

		convertedFilter, err = model.ConvertFilter[*Tag, *domain.Tag](domainFilter.Subtags.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[domain.Tag, Tag](ctx, db, TagDomainToSqlRepositoryModel))
		if err != nil {
			return
		}

		sqlRepositoryFilter.ChildTagTags.Push(convertedFilter)
	}

	return
}

func TagSqlRepositoryToDomainFilter(ctx context.Context, db *sql.DB, sqlRepositoryFilter *TagFilter) (domainFilter *domain.TagFilter, err error) {
	domainFilter = new(domain.TagFilter)

	domainFilter.ID = sqlRepositoryFilter.ID
	domainFilter.Tag = sqlRepositoryFilter.Tag

	//**********************    Set parent path    *********************//
	if sqlRepositoryFilter.ParentTagTags.HasValue {
		var convertedFilter model.FilterOperation[*domain.Tag]

		convertedFilter, err = model.ConvertFilter[*domain.Tag, *Tag](sqlRepositoryFilter.ParentTagTags.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[Tag, domain.Tag](ctx, db, TagSqlRepositoryToDomainModel))
		if err != nil {
			return
		}

		domainFilter.ParentPath.Push(convertedFilter)
	}

	//**********************    Set child tags *********************//
	if sqlRepositoryFilter.ChildTagTags.HasValue {
		var convertedFilter model.FilterOperation[*domain.Tag]

		convertedFilter, err = model.ConvertFilter[*domain.Tag, *Tag](sqlRepositoryFilter.ChildTagTags.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverter[Tag, domain.Tag](ctx, db, TagSqlRepositoryToDomainModel))
		if err != nil {
			return
		}

		domainFilter.Subtags.Push(convertedFilter)
	}

	return
}
