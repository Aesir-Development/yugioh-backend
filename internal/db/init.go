package dbConnection

import (
	"fmt"
)

// Setup the DB structure and tables if they don't exist
func Setup() error {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS cards (" +
		"id INT NOT NULL AUTO_INCREMENT," +
		"name VARCHAR(255) NOT NULL," +
		"type VARCHAR(255) NOT NULL," +
		"frame_type VARCHAR(255) NOT NULL," +
		"description VARCHAR(2000) NOT NULL," +
		"attack INT NOT NULL," +
		"defense INT NOT NULL," +
		"level INT NOT NULL," +
		"race VARCHAR(255) NOT NULL," +
		"attribute VARCHAR(255) NOT NULL," +
		"card_sets JSON," +
		"card_images JSON," +
		"card_prices JSON," +
		"PRIMARY KEY (id)" +
		")")

	if err != nil {
		return fmt.Errorf("error creating cards table: %s", err)
	}


	// TODO- Flesh out the users table
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id INT NOT NULL AUTO_INCREMENT," +
		"puuid VARCHAR(255) NOT NULL," +
		"username VARCHAR(255) NOT NULL," +
		"password VARCHAR(255) NOT NULL," +
		"deck JSON," +
		"PRIMARY KEY (id)" +
		")")

	if err != nil {
		return fmt.Errorf("error creating users table: %s", err)
	}

	// TODO - Deck table setup
	
	// TODO - Player table setup

	return nil

}