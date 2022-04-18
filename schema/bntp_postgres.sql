

CREATE TABLE tags
(
    id  INTEGER PRIMARY KEY NOT NULL,
    tag TEXT                NOT NULL UNIQUE
);

CREATE TABLE bookmark_types
(
    id   INTEGER PRIMARY KEY NOT NULL ,
    Type TEXT    NOT NULL UNIQUE
);

CREATE TABLE bookmarks
(
    id               INTEGER   PRIMARY KEY NOT NULL,
    is_read          INTEGER   DEFAULT 0,
    title            TEXT      UNIQUE,
    url              TEXT      NOT NULL UNIQUE,
    time_added       TIMESTAMP NOT NULL,
    bookmark_type_id INTEGER   REFERENCES bookmark_types(id),
    is_collection    INTEGER   DEFAULT 0
);

CREATE TABLE document_types
(
    id            INTEGER PRIMARY KEY NOT NULL,
    document_type TEXT    NOT NULL UNIQUE
);

CREATE TABLE documents
(
    id               INTEGER PRIMARY KEY NOT NULL,
    path             TEXT    NOT NULL UNIQUE,
    document_type_id INTEGER NOT NULL REFERENCES document_types(id)
);

CREATE TABLE links
(
    source_id      INTEGER  NOT NULL REFERENCES documents(id),
    destination_id INTEGER  NOT NULL REFERENCES documents(id),

    PRIMARY KEY(source_id, destination_id),
    CHECK(source_id != destination_id)
);

CREATE TABLE bookmark_contexts
(
    bookmark_id INTEGER  NOT NULL REFERENCES bookmarks(id),
    tag_id      INTEGER  NOT NULL REFERENCES tags(id),

    PRIMARY KEY(tag_id, bookmark_id)
);

CREATE TABLE document_contexts
(
    document_id INTEGER  NOT NULL REFERENCES documents(id),
    tag_id      INTEGER  NOT NULL REFERENCES  tags(id),

    PRIMARY KEY(tag_id, document_id)
);