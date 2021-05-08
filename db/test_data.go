package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var spacecraft = `
CREATE TABLE IF NOT EXISTS spacecraft (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	name VARCHAR(25) NOT NULL,
	class VARCHAR(25),
	crew INT UNSIGNED,
    image VARCHAR(100),
	value DEC(10,2) UNSIGNED,
	status VARCHAR(17) NOT NULL,
	PRIMARY KEY (id)
);`

var armament = `CREATE TABLE IF NOT EXISTS armament (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	ship_id INT UNSIGNED NOT NULL,
	weapon_id SMALLINT UNSIGNED NOT NULL,
	quantity SMALLINT UNSIGNED NOT NULL,
	PRIMARY KEY (id),
	CONSTRAINT fk_weapon FOREIGN KEY (weapon_id)
	REFERENCES weapons(id),
	CONSTRAINT fk_ship FOREIGN KEY (ship_id)
	REFERENCES spacecraft(id)
);`

var weapons = `CREATE TABLE IF NOT EXISTS weapons (
	id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
	name VARCHAR(25),
	PRIMARY KEY (id)
);
`

func LoadTestData(db *sqlx.DB) {
	fmt.Println("POPULATING DB WITH TEST DATA")
	db.MustExec(weapons)
	db.MustExec(armament)
	db.MustExec(spacecraft)

	// tx := db.MustBegin()
	// tx.MustExec("INSERT INTO weapons (name) VALUES (?)", "Turbo Laser")
	// tx.MustExec("INSERT INTO weapons (name) VALUES (?)", "Ion Cannons")
	// tx.MustExec("INSERT INTO weapons (name) VALUES (?)", "Tractor Beam")
	// tx.Commit()
	// tx := db.MustBegin()
	// tx.MustExec("INSERT INTO armament (ship_id, weapon_id, quantity) VALUES (?, ?, ?)", 1, 13, 60)
	// tx.MustExec("INSERT INTO armament (ship_id, weapon_id, quantity) VALUES (?, ?, ?)", 1, 14, 60)
	// tx.MustExec("INSERT INTO armament (ship_id, weapon_id, quantity) VALUES (?, ?, ?)", 1, 15, 10)
	// tx.Commit()
	// tx = db.MustBegin()
	// tx.MustExec("INSERT INTO spacecraft (name, class, crew, image, value, status) VALUES (?, ?, ?, ?, ?, ?)", "Devastator", "Star Destroyer", 35000, "https://url.to.image", 1999.99, "Operational")
	// tx.MustExec("INSERT INTO spacecraft (name, status) VALUES (? ,?)", "Red Five", "Damaged")
	// tx.Commit()
}
