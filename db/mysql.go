package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	model "github.com/robinsturmeit/imperial-fleet/models"
)

var DBConnect *DBConnector

type DBConnector struct {
	Connection *sqlx.DB
}

type FleetDBConnector interface {
	GetSpacecraftByProp(string, string) (model.Spacecraft, error)
	GetAllSpacecrafts() ([]model.Spacecraft, error)
}

func InitDBConnector() error {
	dbConnector := &DBConnector{}
	var err error
	dbDSN := os.Getenv("MYSQL_DSN")
	dbConnector.Connection, err = sqlx.Connect("mysql", dbDSN)
	if err != nil {
		log.Printf("[ERROR]: Could not connect to Imperial Database: %s", err.Error())
		return err
	}
	DBConnect = dbConnector
	return nil
}

func (db *DBConnector) GetSpacecraftByProp(prop string, value string) (model.Spacecraft, error) {
	spacecraft := model.Spacecraft{}
	err := db.Connection.Get(&spacecraft, fmt.Sprintf("SELECT * FROM spacecraft WHERE %s=?", prop), value)
	if err != nil {
		log.Printf("[ERROR]: Could not retrieve spacecraft by %s [%s]: %s", prop, value, err.Error())
		return spacecraft, err
	}

	var armament []model.Armament
	err = db.Connection.Select(&armament, "SELECT quantity, name FROM armament RIGHT JOIN weapons ON weapons.id = weapon_id WHERE ship_id=?", spacecraft.Id)
	if err != nil {
		log.Printf("[ERROR]: Could not retrieve armament for spacecraft [id=%d, name=%s]: %s", spacecraft.Id, spacecraft.Name, err.Error())
		return spacecraft, nil
	}
	spacecraft.Armament = armament
	return spacecraft, nil
}

func (db *DBConnector) GetAllSpacecrafts() ([]model.Spacecraft, error) {
	var spacecrafts []model.Spacecraft
	//FIXME: This needs pagination
	err := db.Connection.Select(&spacecrafts, "SELECT * FROM spacecraft")
	if err != nil {
		log.Printf("[ERROR]: Could not retrieve spacecrafts: %s", err.Error())
		return spacecrafts, err
	}

	var result []model.Spacecraft
	//FIXME: This could use parallelization/fan-out pattern
	for _, s := range spacecrafts {
		var armament []model.Armament
		err = db.Connection.Select(&armament, "SELECT quantity, name FROM armament RIGHT JOIN weapons ON weapons.id = weapon_id WHERE ship_id=?", s.Id)
		if err != nil {
			log.Printf("[ERROR]: Could not retrieve armament for spacecraft [id=%d, name=%s]: %s", s.Id, s.Name, err.Error())
			//FIXME: Use error codes to allow incomplete results with armament information missing with error field in response
		}
		s.Armament = armament
		result = append(result, s)
	}
	return result, nil
}
