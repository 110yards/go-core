package postgres

import "110yards.ca/libs/go/core/logger"

type SeedFunction func() error

func SeedDatabase(seedFunctions []SeedFunction) error {

	for _, f := range seedFunctions {
		err := f()

		if err != nil {
			return err
		}
	}

	return nil
}

func doesTableExist(tableName string) (bool, error) {
	db := GetDb()

	// check if Seasons table exists:
	logger.Infof("Checking if %s table exists", tableName)
	result, err := db.Query(`SELECT EXISTS (
		SELECT FROM information_schema.tables
		WHERE  table_schema = 'public'
		AND    table_name   = $1
	);`, tableName)

	if err != nil {
		return false, err
	}

	defer result.Close()

	var exists bool

	for result.Next() {
		err = result.Scan(&exists)
		if err != nil {
			return false, err
		}
	}

	return exists, nil
}

func doesConstraintExist(tableName, constraintName string) (bool, error) {
	db := GetDb()

	// check if constraint exists:

	result, err := db.Query(`SELECT EXISTS (
		SELECT 1
		FROM   information_schema.table_constraints
		WHERE  constraint_name = $1
		AND    table_name = $2
	);`, constraintName, tableName)

	if err != nil {
		return false, err
	}

	defer result.Close()

	var exists bool

	for result.Next() {
		err = result.Scan(&exists)
		if err != nil {
			return false, err
		}
	}

	return exists, nil

}

func CreateTableIfMissing(tableName string, createQuery string) error {
	db := GetDb()

	exists, err := doesTableExist(tableName)

	if err != nil {
		return err
	}

	if exists {
		logger.Infof("%s table already exists", tableName)
		return nil
	}

	logger.Infof("%s table does not exist, creating it", tableName)

	_, err = db.Exec(createQuery)

	return err
}

func CreateConstraintIfMissing(tableName, constraintName, constraintQuery string) error {
	db := GetDb()

	exists, err := doesConstraintExist(tableName, constraintName)

	if err != nil {
		return err
	}

	if exists {
		logger.Infof("%s constraint already exists", constraintName)
		return nil
	}

	logger.Infof("%s constraint does not exist, creating it", constraintName)

	_, err = db.Exec(constraintQuery)

	return err
}
