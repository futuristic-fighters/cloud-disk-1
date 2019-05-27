package user

import (
	"cloud-disk/cfg"
	"database/sql"
)

func updateValidator(u *User, db *sql.DB) *cfg.Err {

	if u.Name != "" {
		if err := u.CheckName(); err != nil {
			return err
		}
	}

	//use db check the name is unique
	//....

	if err := u.CheckType(); err != nil {
		return err
	}

	return nil
}
