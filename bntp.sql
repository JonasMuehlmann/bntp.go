PRAGMA foreign_keys = ON;

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
    document_type Text    NOT NULL UNIQUE
);
CREATE TABLE documents
(
    id               INTEGER PRIMARY KEY NOT NULL,
    path             Text    NOT NULL UNIQUE,
    document_type_id INTEGER REFERENCES document_types(id) NOT NULL
);
CREATE TABLE links
(
    source_id      INTEGER REFERENCES documents(id) NOT NULL,
    destination_id INTEGER REFERENCES documents(id) NOT NULL,

    PRIMARY KEY(source_id, destination_id),
    CHECK(source_id != destination_id)
);
CREATE TABLE bookmark_contexts
(
    bookmark_id INTEGER REFERENCES bookmarks(id) NOT NULL,
    tag_id      INTEGER REFERENCES tags(id)      NOT NULL,

    PRIMARY KEY(tag_id, bookmark_id)
);
CREATE TABLE document_contexts
(
    document_id INTEGER REFERENCES documents(id) NOT NULL,
    tag_id      INTEGER REFERENCES  tags(id)     NOT NULL,

    PRIMARY KEY(tag_id, document_id)
);
