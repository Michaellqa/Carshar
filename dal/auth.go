package dal

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

const (
	SqlCreateUser = `
INSERT INTO "User"("PhoneNumber", "Password", "Name", "BirthDate") VALUES 
($1, $2, $3, $4);
`
)

type UserDb struct {
	db *sql.DB
}

func NewAuthDb(db *sql.DB) *UserDb {
	return &UserDb{db: db}
}

func (a *UserDb) CreateUser(u User) (bool, error) {
	_, err := a.db.Exec(SqlCreateUser, u.Phone, u.Password, u.Name, u.BirthDate)

	if err != nil {
		// 23505 unique values conflict
		if per := err.(*pq.Error); per.Code == "23505" {
			return false, nil
		}
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (a *UserDb) FindUser(phone string) (User, bool, error) {
	SqlFindUser := `SELECT "Id", "PhoneNumber", "Password", "Name", "BirthDate", "CreditAmount" FROM "User" WHERE "PhoneNumber" = $1;`
	var u User

	err := a.db.QueryRow(SqlFindUser, phone).Scan(&u.Id, &u.Phone, &u.Password, &u.Name, &u.BirthDate, &u.Balance)
	if err == sql.ErrNoRows {
		return u, false, nil
	}
	if err != nil {
		log.Println(err)
		return u, false, err
	}
	return u, true, nil
}

func (a *UserDb) GetUser(id int) (User, error) {
	SqlFindUser := `SELECT "Id", "PhoneNumber", "Password", "Name", "BirthDate", "CreditAmount" FROM "User" WHERE "Id" = $1;`
	var u User

	err := a.db.QueryRow(SqlFindUser, id).Scan(&u.Id, &u.Phone, &u.Password, &u.Name, &u.BirthDate, &u.Balance)
	if err != nil {
		log.Println(err)
		return u, err
	}
	return u, nil
}

func (a *UserDb) TransferMoney(from, to int, amount float64) error {
	SqlChangeAmount := `UPDATE "User" SET "CreditAmount" = "CreditAmount" + $1  WHERE "Id" = $2;`
	//begin-end transaction
	_, err := a.db.Exec(SqlChangeAmount, -amount, from)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = a.db.Exec(SqlChangeAmount, amount, to)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
