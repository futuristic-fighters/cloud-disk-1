package user

func updateValidator(u *User) error {

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
