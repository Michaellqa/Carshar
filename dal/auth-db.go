package dal

import (
	"database/sql"
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

type AuthDb struct {
	db *sql.DB
}

func NewAuthDb(db *sql.DB) *AuthDb {
	return &AuthDb{db: db}
}

func (a *AuthDb) CreateUser(u User) error {
	_, err := a.db.Exec(SqlCreateUser, u.Phone, u.Password, u.Name, u.BirthDate)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//return User, flag that user exists, error
func (a *AuthDb) FindUser(phone string) (User, bool, error) {
	var (
		u User
	)
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
