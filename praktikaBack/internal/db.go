package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type IDB interface {
	CloseDB()
	InsertVacancies() error
	GetVacancies() (*vacancies, error)
	RemoveVacancies() error
}

type db struct {
	db  *sql.DB
	ctx context.Context
}

func NewDB() (IDB, error) {
	cont := context.Background()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	sqlConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return &db{db: nil}, fmt.Errorf("error in connection to database: %s", err)
	}

	dbStruct := &db{db: sqlConn, ctx: cont};
	err = dbStruct.InstallSchema();
	if err != nil {
		return &db{db: nil}, err;
	}

	return dbStruct, nil
}

func (db *db) InstallSchema() error {
	schemaDownFilePath := "./schema/000001_vacancies.down.sql"
	schemaSQL, err := os.ReadFile(schemaDownFilePath);
	if err != nil {
		return fmt.Errorf("failed to read down schema file: %v", err)
	}

	_, err = db.db.ExecContext(context.Background(), string(schemaSQL))
	if err != nil {
		return fmt.Errorf("failed to execute down schema: %v", err)
	}

	schemaUpFilePath := "./schema/000001_vacancies.up.sql"
	schemaSQL, err = os.ReadFile(schemaUpFilePath)
	if err != nil {
		return fmt.Errorf("failed to read up schema file: %v", err)
	}

	_, err = db.db.ExecContext(context.Background(), string(schemaSQL))
	if err != nil {
		return fmt.Errorf("failed to execute up schema: %v", err)
	}

	return nil;
}

func (db *db) CloseDB() {
	db.db.Close();
}

func (db *db) InsertVacancies() error {
	contextWithTimeout, cancel := context.WithTimeout(db.ctx, time.Second*5)
	defer cancel()

	vacanciesData, err := getVacancies(contextWithTimeout)
	if err != nil {
		return err
	}

	insertStr := vacanciesData.getInsertString()

	result, err := db.db.Exec(insertStr)
	if err != nil {
		return fmt.Errorf("error in insert vacancies to table: %s", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in get rows affected info: %s", err)
	}
	log.Printf("Rows affected: %v\n", rows)

	return nil
}

func (db *db) GetVacancies() (*vacancies, error) {
	contextWithTimeout, cancel := context.WithTimeout(db.ctx, time.Second*5)
	defer cancel()

	getStr := `select address_building, address_city, address_description, address_metro_line_name, address_metro_station_name, 
		address_street, area_name, contacts_email, 
		contacts_name, created_at, description, employment_name, experience_name, id, key_skills, name, salary_currency,
		salary_from, salary_to, schedule_name from vacancies`

	result, err := db.db.QueryContext(contextWithTimeout, getStr)
	if err != nil {
		return nil, fmt.Errorf("error in insert vacancies to table: %s", err)
	}
	defer result.Close()

	vacList := vacancies{}
	for result.Next() {
		vacancy := vacancy{}
		skillsStr := ""
		metroLineStr := ""
		metroStationStr := ""
		err := result.Scan(&vacancy.Address.Building, &vacancy.Address.City, &vacancy.Address.Description,
			&metroLineStr, &metroStationStr,
			&vacancy.Address.Street, &vacancy.Area.Name, &vacancy.Contacts.Email, &vacancy.Contacts.Name, &vacancy.CreatedAt,
			&vacancy.Description, &vacancy.Employment.Name, &vacancy.Experience.Name, &vacancy.ID, &skillsStr, &vacancy.Name,
			&vacancy.Salary.Currency, &vacancy.Salary.From, &vacancy.Salary.To, &vacancy.Schedule.Name)
		if err != nil {
			return nil, fmt.Errorf("error in scan vacancies to struct: %s", err)
		}

		for _, el := range strings.Split(skillsStr, " ") {
			vacancy.KeySkills = append(vacancy.KeySkills, skill{Name: el})
		}

		metroLineArr := strings.Split(metroLineStr, " ")
		metroStationArr := strings.Split(metroStationStr, " ")
		for i := 0; i < len(metroLineArr); i++ {
			vacancy.Address.MetroStations = append(vacancy.Address.MetroStations,
				metroStation{LineName: metroLineArr[i], StationName: metroStationArr[i]})
		}

		vacList.Vacancies = append(vacList.Vacancies, vacancy)
	}

	return &vacList, nil
}

func (db *db) RemoveVacancies() error {
	removeStr := `delete from "vacancies"`

	result, err := db.db.Exec(removeStr)
	if err != nil {
		return fmt.Errorf("error in delete vacancies from table: %s", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in get rows affected info: %s", err)
	}
	log.Printf("Rows affected: %v\n", rows)

	return nil
}
