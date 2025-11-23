package database

func Initialize() {
	connect()

	migrationsUp()
}

func CloseConnection() {
	close()
}
