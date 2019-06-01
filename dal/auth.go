package dal

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

const (
	SqlCreateUser = `
INSERT INTO "User"("Phone", "Password", "Name", "BirthDate") VALUES 
($1, $2, $3, $4);
`
	SqlFindUser = `
SELECT "Id", "Phone", "Password", "Name", "BirthDate" FROM "User"
WHERE "Phone" = $1;
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
	var u User

	err := a.db.QueryRow(SqlFindUser, phone).Scan(&u.Id, &u.Phone, &u.Password, &u.Name, &u.BirthDate)
	if err == sql.ErrNoRows {
		return u, false, nil
	}
	if err != nil {
		log.Println(err)
		return u, false, err
	}
	return u, true, nil
}
