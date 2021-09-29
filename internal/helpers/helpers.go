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
			return 0, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	var typeId int

	err = statement.Get(&typeId, type_)

	if err != nil {
		return 0, err
	}

	err = statement.Close()

	if err != nil {
		return 0, err
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
			return 0, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	var documentId int

	err = statement.Get(&documentId, documentPath)

	if err != nil {
		return 0, err
	}

	err = statement.Close()

	if err != nil {
		return 0, err
	}

	return documentId, nil
}
