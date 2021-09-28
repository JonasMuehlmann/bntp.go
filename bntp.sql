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
    Title        TEXT    not null,
    Url          TEXT    not null
        unique,
    TimeAdded    TEXT    not null,
    BookmarkTypeId
        references BookmarkType (Id),
    IsCollection INTEGER not null
);
CREATE TABLE DocumentType
(
    Id INTEGER PRIMARY KEY,
    DocumentType Text
);
CREATE TABLE Document
(
    Id   INTEGER PRIMARY KEY,
    Path Text,
    DocumentTypeId REFERENCES DocumentType(Id)
);
CREATE TABLE Link
(
    Id           INTEGER PRIMARY KEY,
    SourceId REFERENCES Document(Id),
    DestinationId REFERENCES Document(Id)
);
CREATE TABLE BookmarkContext
(
    Id PRIMARY KEY,
    BookmarkId REFERENCES Bookmark(Id),
    BookmarkTypeId REFERENCES BookmarkType(Id)
);
CREATE TABLE DocumentContext
(
    Id PRIMARY KEY,
    DocumentId REFERENCES Document (Id),
    DocumentTypeId REFERENCES DocumentType (Id)
);
