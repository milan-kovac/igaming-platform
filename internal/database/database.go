package database

func Initialize() {
	connect()

	migrationsUp()

	defer close()
}
