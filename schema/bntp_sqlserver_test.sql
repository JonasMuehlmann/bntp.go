

CREATE TABLE tags
(
    id  INTEGER PRIMARY KEY,
    tag VARCHAR(255)                NOT NULL UNIQUE
);

CREATE TABLE bookmark_types
(
    id   INTEGER PRIMARY KEY ,
    Type VARCHAR(255)    NOT NULL UNIQUE
);

CREATE TABLE bookmarks
(
    id               INTEGER   PRIMARY KEY,
    is_read          INTEGER   DEFAULT 0,
    title            VARCHAR(255)      UNIQUE,
    url              VARCHAR(255)      NOT NULL UNIQUE,
    time_added       DATETIME NOT NULL,
    bookmark_type_id INTEGER   ,
    is_collection    INTEGER   DEFAULT 0
);

CREATE TABLE document_types
(
    id            INTEGER PRIMARY KEY,
    document_type VARCHAR(255)    NOT NULL UNIQUE
);

CREATE TABLE documents
(
    id               INTEGER PRIMARY KEY,
    path             VARCHAR(255)    NOT NULL UNIQUE,
    document_type_id INTEGER NOT NULL 
);

CREATE TABLE links
(
    source_id      INTEGER  NOT NULL ,
    destination_id INTEGER  NOT NULL ,

    PRIMARY KEY(source_id, destination_id),
    CHECK(source_id != destination_id)
);

CREATE TABLE bookmark_contexts
(
    bookmark_id INTEGER  NOT NULL ,
    tag_id      INTEGER  NOT NULL ,

    PRIMARY KEY(tag_id, bookmark_id)
);

CREATE TABLE document_contexts
(
    document_id INTEGER  NOT NULL ,
    tag_id      INTEGER  NOT NULL ,

    PRIMARY KEY(tag_id, document_id)
);
