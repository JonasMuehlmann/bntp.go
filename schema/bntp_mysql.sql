

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
    bookmark_type_id INTEGER   REFERENCES bookmark_types(id),
    is_collection    INTEGER   DEFAULT 0,
    created_at       TIMESTAMP NOT NULL,
    updated_at       TIMESTAMP NOT NULL,
    deleted_at       TIMESTAMP
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
    document_type_id INTEGER NOT NULL REFERENCES document_types(id),
    created_at       TIMESTAMP NOT NULL,
    updated_at       TIMESTAMP NOT NULL,
    deleted_at       TIMESTAMP
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
