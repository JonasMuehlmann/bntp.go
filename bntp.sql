PRAGMA foreign_keys = ON;

CREATE TABLE Tag
(
    Id  INTEGER PRIMARY KEY NOT NULL,
    Tag TEXT                NOT NULL UNIQUE
);
CREATE TABLE BookmarkType
(
    Id   INTEGER not null
        primary key,
    Type TEXT    not null
        unique
);
CREATE TABLE Bookmark
(
    Id           INTEGER not null
        primary key,
    IsRead       INTEGER default 0,
    Title        TEXT, 
    Url          TEXT    not null
        unique,
    TimeAdded    TEXT    not null,
    BookmarkTypeId
        references BookmarkType (Id),
    IsCollection INTEGER
);
CREATE TABLE DocumentType
(
    Id INTEGER PRIMARY KEY,
    DocumentType Text NOT NULL
);
CREATE TABLE Document
(
    Id   INTEGER PRIMARY KEY,
    Path Text NOT NULL,
    DocumentTypeId REFERENCES DocumentType(Id) NOT NULL
);
CREATE TABLE Link
(
    Id           INTEGER PRIMARY KEY,
    SourceId REFERENCES Document(Id) NOT NULL,
    DestinationId REFERENCES Document(Id) NOT NULL,
    UNIQUE(SourceId, DestinationId),
    CHECK(SourceId != DestinationId)
);
CREATE TABLE BookmarkContext
(
    Id PRIMARY KEY,
    BookmarkId REFERENCES Bookmark(Id) NOT NULL,
    TagId REFERENCES Tag(Id) NOT NULL 
);
CREATE TABLE DocumentContext
(
    Id PRIMARY KEY,
    DocumentId REFERENCES Document (Id)NOT NULL,
    TagId REFERENCES  Tag(Id) NOT NULL
);
