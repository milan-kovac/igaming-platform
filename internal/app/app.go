package app

import "igaming-platform/internal/database"

func InitializeApp() {
	database.Initialize()
}
