package errs

import "database/sql"

func ConvertSqlError(err error) error {
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return NewRecordNotFoundError()
	}
	return err
}
