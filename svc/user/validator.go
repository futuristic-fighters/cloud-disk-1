package user

import "cloud-disk/cfg"

func updateValidator(u *User) *cfg.Err {

	if u.Name != "" {
		if err := u.CheckName(); err != nil {
			return err
		}
	}

	if err := u.CheckType(); err != nil {
		return err
	}

	return nil
}
