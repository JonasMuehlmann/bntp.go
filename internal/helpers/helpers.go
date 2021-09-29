package helpers

import "github.com/jmoiron/sqlx"

func GetIdFromType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            DocumentType
        WHERE
            Type = ?;
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	}

	if err != nil {
		return -1, err
	}

	typeId := -1

	_ = statement.Get(&typeId, type_)

	err = statement.Close()

	if err != nil {
		return -1, err
	}

	return typeId, nil
}

func GetIdFromDocument(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPath string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            Document
        WHERE
            Path = ?;
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	}

	if err != nil {
		return -1, err
	}

	documentId := -1

	_ = statement.Get(&documentId, documentPath)

	err = statement.Close()

	if err != nil {
		return -1, err
	}

	return documentId, nil
}
