package simple

type Database struct {
	Name string
}

type DatabasePostgresSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgresSQL() *DatabasePostgresSQL {
	return (*DatabasePostgresSQL)(&Database{Name: "PostgresSQL"})
}

type DatabaseRepository struct {
	DatabasePostgresSQL *DatabasePostgresSQL
	DatabaseMongoDB     *DatabaseMongoDB
}

func NewDatabaseRepository(databaseMongoDB *DatabaseMongoDB, databasePostgresSQL *DatabasePostgresSQL) *DatabaseRepository {
	return &DatabaseRepository{DatabaseMongoDB: databaseMongoDB, DatabasePostgresSQL: databasePostgresSQL}
}
