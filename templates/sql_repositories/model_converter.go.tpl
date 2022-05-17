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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/model_converter.go.tpl

package repository

import (
	"github.com/JonasMuehlmann/optional.go"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
    "context"
    "database/sql"
    "time"
)

func BookmarkDomainToSqlRepositoryModel(db *sql.DB, domainModel domain.{{.Entities.Bookmark}}) ( sqlRepositoryModel {{.Entities.Bookmark}}, err error)  {
    sqlRepositoryModel = {{.Entities.Bookmark}}{}

    sqlRepositoryModel.URL = domainModel.URL
    sqlRepositoryModel.ID = domainModel.ID


    //**********************    Set Timestamps    **********************//
    {{ if eq .DatabaseName "sqlite3"}}
    sqlRepositoryModel.CreatedAt = domainModel.CreatedAt.Format("2006-01-02 15:04:05")
    sqlRepositoryModel.UpdatedAt = domainModel.UpdatedAt.Format("2006-01-02 15:04:05")

    if domainModel.DeletedAt.HasValue {
        sqlRepositoryModel.DeletedAt.Valid = true
        sqlRepositoryModel.DeletedAt.String = domainModel.DeletedAt.Wrappee.Format("2006-01-02 15:04:05")
    }
    {{else}}
    sqlRepositoryModel.CreatedAt = domainModel.CreatedAt
    sqlRepositoryModel.UpdatedAt = domainModel.UpdatedAt
    sqlRepositoryModel.DeletedAt = domainModel.DeletedAt

    if domainModel.DeletedAt.HasValue {
        sqlRepositoryModel.DeletedAt.Valid = true
        sqlRepositoryModel.DeletedAt.String = domainModel.DeletedAt
    }
    {{end}}


    //*************************    Set Title    ************************//
    if domainModel.Title.HasValue {
        sqlRepositoryModel.Title.Valid = true
        sqlRepositoryModel.Title.String = domainModel.Title.Wrappee
    }



    //******************    Set IsRead/IsCollection    *****************//
    {{ if eq .DatabaseName "sqlite3"}}
    if domainModel.IsRead {
        sqlRepositoryModel.IsRead = 1
    }

    if domainModel.IsCollection {
        sqlRepositoryModel.IsCollection = 1
    }
    {{else}}
    sqlRepositoryModel.IsRead = domainModel.IsRead
    sqlRepositoryModel.IsCollection = domainModel.IsCollection
    {{end}}


    //*************************    Set Tags    *************************//
    var repositoryTag *Tag

    sqlRepositoryModel.R.Tags = make(TagSlice, 0, len(domainModel.Tags))
	for _,  domainTag := range domainModel.Tags {
		repositoryTag, err = Tags(TagWhere.Tag.EQ(domainTag.Tag)).One(context.Background(), db)
		if err != nil {
			return
		}

        sqlRepositoryModel.R.Tags = append(sqlRepositoryModel.R.Tags, &Tag{Tag: repositoryTag.Tag, ID: repositoryTag.ID})
	}


    //*************************    Set Type    *************************//
	if domainModel.BookmarkType.HasValue {
        var repositoryBookmarkType *BookmarkType

		sqlRepositoryModel.R.BookmarkType.Type = domainModel.BookmarkType.Wrappee
		repositoryBookmarkType, err = BookmarkTypes(BookmarkTypeWhere.Type.EQ(domainModel.BookmarkType.Wrappee)).One(context.Background(), db)
		if err != nil {
			return
		}

		sqlRepositoryModel.BookmarkTypeID.Int64 = repositoryBookmarkType.ID
		sqlRepositoryModel.BookmarkTypeID.Valid = true

		sqlRepositoryModel.R.BookmarkType.ID = repositoryBookmarkType.ID
	}

    return
}

func BookmarkSqlRepositoryToDomainModel(db *sql.DB, sqlRepositoryModel {{.Entities.Bookmark}}) (domainModel domain.{{.Entities.Bookmark}}, err error) {
    domainModel = domain.{{.Entities.Bookmark}}{}

    domainModel.URL = sqlRepositoryModel.URL
    domainModel.ID = sqlRepositoryModel.ID
	domainModel.BookmarkType = optional.Make(sqlRepositoryModel.R.BookmarkType.Type)

    //**********************    Set Timestamps    **********************//
    {{ if eq .DatabaseName "sqlite3"}}
    domainModel.CreatedAt, err = time.Parse("2006-01-02 15:04:05", sqlRepositoryModel.CreatedAt)
    if err != nil {
        return
    }

    domainModel.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", sqlRepositoryModel.UpdatedAt)
    if err != nil {
        return
    }

    if sqlRepositoryModel.DeletedAt.Valid {
        var t time.Time

        t, err = time.Parse("2006-01-02 15:04:05", sqlRepositoryModel.DeletedAt.String)
        if err != nil {
            return
        }

        domainModel.DeletedAt.Push(t)
    }
    {{else}}
    domainModel.CreatedAt = sqlRepositoryModel.CreatedAt
    domainModel.UpdatedAt = sqlRepositoryModel.UpdatedAt
    domainModel.DeletedAt = sqlRepositoryModel.DeletedAt

    if sqlRepositoryModel.DeletedAt.Valid {
        domainModel.DeletedAt.Push(sqlRepositoryModel.DeletedAt.Time)
    }
    {{end}}

    //*************************    Set Title    ************************//
    if sqlRepositoryModel.Title.Valid {
        domainModel.Title.Push(sqlRepositoryModel.Title.String)
    }

    //******************    Set IsRead/IsCollection    *****************//
    {{ if eq .DatabaseName "sqlite3"}}
    domainModel.IsRead = sqlRepositoryModel.IsRead > 0
    domainModel.IsCollection = sqlRepositoryModel.IsCollection > 0
    {{else}}
    domainModel.IsRead = sqlRepositoryModel.IsRead
    domainModel.IsCollection = sqlRepositoryModel.IsCollection
    {{end}}

    //*************************    Set Tags    *************************//
    var domainTag domain.Tag

	domainModel.Tags = make([]*domain.Tag, 0, len(sqlRepositoryModel.R.Tags))
    for _, repositoryTag := range sqlRepositoryModel.R.Tags {
        domainTag, err = TagSqlRepositoryToDomainModel(db, *repositoryTag)
        if err != nil {
            return
        }

        domainModel.Tags = append(domainModel.Tags, &domainTag)
    }

    return
}

func DocumentDomainToSqlRepositoryModel(db *sql.DB, domainModel domain.{{.Entities.Document}}) (sqlRepositoryModel {{.Entities.Document}}, err error)  {
    sqlRepositoryModel = {{.Entities.Document}}{}

    sqlRepositoryModel.Path = domainModel.Path
    sqlRepositoryModel.ID = domainModel.ID


    //**********************    Set Timestamps    **********************//
    {{ if eq .DatabaseName "sqlite3"}}
    sqlRepositoryModel.CreatedAt = domainModel.CreatedAt.Format("2006-01-02 15:04:05")
    sqlRepositoryModel.UpdatedAt = domainModel.UpdatedAt.Format("2006-01-02 15:04:05")

    if domainModel.DeletedAt.HasValue {
        sqlRepositoryModel.DeletedAt.Valid = true
        sqlRepositoryModel.DeletedAt.String = domainModel.DeletedAt.Wrappee.Format("2006-01-02 15:04:05")
    }
    {{else}}
    sqlRepositoryModel.CreatedAt = domainModel.CreatedAt
    sqlRepositoryModel.UpdatedAt = domainModel.UpdatedAt
    sqlRepositoryModel.DeletedAt = domainModel.DeletedAt

    if domainModel.DeletedAt.HasValue {
        sqlRepositoryModel.DeletedAt.Valid = true
        sqlRepositoryModel.DeletedAt.String = domainModel.DeletedAt
    }
    {{end}}

    //*************************    Set Tags    *************************//
    var repositoryTag *Tag

	sqlRepositoryModel.R.Tags = make(TagSlice, 0, len(domainModel.Tags))
	for _, modelTag := range domainModel.Tags {
		repositoryTag, err = Tags(TagWhere.Tag.EQ(modelTag.Tag)).One(context.Background(), db)
		if err != nil {
			return
		}

		sqlRepositoryModel.R.Tags  = append(sqlRepositoryModel.R.Tags, &Tag{Tag: modelTag.Tag, ID: repositoryTag.ID})
	}

    //*************************    Set Type    *************************//
    var repositoryDocumentType *DocumentType

	if domainModel.DocumentType.HasValue {
		sqlRepositoryModel.R.DocumentType.DocumentType = domainModel.DocumentType.Wrappee
		repositoryDocumentType, err = DocumentTypes(DocumentTypeWhere.DocumentType.EQ(domainModel.DocumentType.Wrappee)).One(context.Background(), db)
		if err != nil {
			return
		}

		sqlRepositoryModel.DocumentTypeID = repositoryDocumentType.ID
		sqlRepositoryModel.R.DocumentType.ID = repositoryDocumentType.ID
	}


    //**************    Set linked/backlinked documents    *************//
    var repositoryDocument Document

    sqlRepositoryModel.R.SourceDocuments  = make(DocumentSlice, 0, len(domainModel.LinkedDocuments))
    sqlRepositoryModel.R.DestinationDocuments  = make(DocumentSlice, 0, len(domainModel.BacklinkedDocuments))

    for _ , link := range domainModel.LinkedDocuments {
        repositoryDocument, err = DocumentDomainToSqlRepositoryModel(db, *link)
        if err != nil {
            return
        }

        sqlRepositoryModel.R.SourceDocuments = append(sqlRepositoryModel.R.SourceDocuments, &repositoryDocument)
    }

    for _ , backlink := range domainModel.BacklinkedDocuments {
        repositoryDocument, err = DocumentDomainToSqlRepositoryModel(db, *backlink)
        if err != nil {
            return
        }

        sqlRepositoryModel.R.DestinationDocuments = append(sqlRepositoryModel.R.DestinationDocuments, &repositoryDocument)
    }

    return
}

func DocumentSqlRepositoryToDomainModel(db *sql.DB, sqlRepositoryModel {{.Entities.Document}}) (domainModel domain.{{.Entities.Document}}, err error) {
    domainModel = domain.{{.Entities.Document}}{}

    domainModel.Path = sqlRepositoryModel.Path
    domainModel.ID = sqlRepositoryModel.ID
	domainModel.DocumentType = optional.Make(sqlRepositoryModel.R.DocumentType.DocumentType)

    //**********************    Set Timestamps    **********************//
    {{ if eq .DatabaseName "sqlite3"}}
    domainModel.CreatedAt, err = time.Parse("2006-01-02 15:04:05", sqlRepositoryModel.CreatedAt)
    if err != nil {
        return
    }

    domainModel.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", sqlRepositoryModel.UpdatedAt)
    if err != nil {
        return
    }

    var t time.Time

    if sqlRepositoryModel.DeletedAt.Valid {
        t, err = time.Parse("2006-01-02 15:04:05", sqlRepositoryModel.DeletedAt.String)
        if err != nil {
            return
        }

        domainModel.DeletedAt.Push(t)
    }
    {{else}}
    domainModel.CreatedAt = sqlRepositoryModel.CreatedAt
    domainModel.UpdatedAt = sqlRepositoryModel.UpdatedAt
    domainModel.DeletedAt = sqlRepositoryModel.DeletedAt

    if sqlRepositoryModel.DeletedAt.Valid {
        domainModel.DeletedAt.Push(sqlRepositoryModel.DeletedAt.Time)
    }
    {{end}}

    //*************************    Set Tags    *************************//
    var domainTag domain.Tag

	domainModel.Tags = make([]*domain.Tag, 0, len(sqlRepositoryModel.R.Tags))
    for _, repositoryTag := range sqlRepositoryModel.R.Tags {
    domainTag, err = TagSqlRepositoryToDomainModel(db, *repositoryTag)
        if err != nil {
            return
        }

        domainModel.Tags = append(domainModel.Tags, &domainTag)
    }

    //**************    Set linked/backlinked documents    *************//
    var domainDocument domain.Document

    domainModel.LinkedDocuments = make([]*domain.Document, 0, len(sqlRepositoryModel.R.SourceDocuments))
    domainModel.BacklinkedDocuments = make([]*domain.Document, 0, len(sqlRepositoryModel.R.DestinationDocuments))

    for _ , link := range sqlRepositoryModel.R.SourceDocuments {
        domainDocument, err = DocumentSqlRepositoryToDomainModel(db, *link)
        if err != nil {
            return
        }

        domainModel.LinkedDocuments = append(domainModel.LinkedDocuments, &domainDocument)
    }

    for _ , backlink := range sqlRepositoryModel.R.DestinationDocuments {
        domainDocument, err = DocumentSqlRepositoryToDomainModel(db, *backlink)
        if err != nil {
            return
        }

        domainModel.BacklinkedDocuments = append(domainModel.BacklinkedDocuments, &domainDocument)
    }

    return
}

func TagDomainToSqlRepositoryModel(db *sql.DB, domainModel domain.{{.Entities.Tag}}) (sqlRepositoryModel {{.Entities.Tag}}, err error)  {
    sqlRepositoryModel = {{.Entities.Tag}}{}

    sqlRepositoryModel.ID = domainModel.ID
    sqlRepositoryModel.Tag = domainModel.Tag


    //**********************    Set parent path    *********************//
    var repositoryTag Tag
    var repositoryParentPathTag TagParentPath

    sqlRepositoryModel.R.ParentTagTagParentPaths = make(TagParentPathSlice, 0, len(domainModel.ParentPath))
    sqlRepositoryModel.R.ChildTagTags = make(TagSlice, 0, len(domainModel.Subtags))

    for distance, domainTag := range domainModel.ParentPath {
        repositoryParentPathTag, err = tagDomainToRepositoryParentPathModel(db, *domainTag, distance)
        if err != nil {
            return
        }

        sqlRepositoryModel.R.ParentTagTagParentPaths = append(sqlRepositoryModel.R.ParentTagTagParentPaths, &repositoryParentPathTag)
    }

    //**********************    Set child tags *********************//
    for _, domainTag := range domainModel.Subtags {
        repositoryTag, err = TagDomainToSqlRepositoryModel(db, *domainTag)
        if err != nil {
            return
        }

        sqlRepositoryModel.R.ChildTagTags = append(sqlRepositoryModel.R.ChildTagTags, &repositoryTag)
    }


    return
}

func  tagDomainToRepositoryParentPathModel(db *sql.DB, domainModel domain.{{.Entities.Tag}},distance int) (sqlRepositoryModel TagParentPath, err error)  {
    sqlRepositoryModel = TagParentPath{}

    sqlRepositoryModel.TagID = domainModel.ID
    sqlRepositoryModel.ParentTagID = domainModel.ID
    sqlRepositoryModel.Distance = int64(distance)

    return
}

func TagSqlRepositoryToDomainModel(db *sql.DB, sqlRepositoryModel {{.Entities.Tag}}) (domainModel domain.{{.Entities.Tag}}, err error) {
    domainModel = domain.{{.Entities.Tag}}{}

    domainModel.ID = sqlRepositoryModel.ID
    domainModel.Tag = sqlRepositoryModel.Tag

    //**********************    Set parent path    *********************//
    var domainTag domain.{{.Entities.Tag}}

    domainModel.ParentPath   = make([]*domain.{{.Entities.Tag}}, 0, len(sqlRepositoryModel.R.ParentTagTagParentPaths))

    for _, repositoryParentPathTag := range sqlRepositoryModel.R.ParentTagTagParentPaths  {
        domainTag, err = TagSqlRepositoryToDomainModel(db,  *repositoryParentPathTag.R.Tag)
        if err != nil {
            return
        }

        domainModel.ParentPath[repositoryParentPathTag.Distance] = &domainTag
    }

    //**********************    Set child tags *********************//
    domainModel.Subtags = make([]*domain.{{.Entities.Tag}}, 0, len(sqlRepositoryModel.R.ChildTagTags))

    for _, repositoryTag := range sqlRepositoryModel.R.ChildTagTags {
        domainTag, err = TagSqlRepositoryToDomainModel(db, *repositoryTag)
        if err != nil {
            return
        }

        domainModel.Subtags = append(domainModel.Subtags, &domainTag)
    }

    return
}
