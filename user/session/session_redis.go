package session

func (u UserSession) Store(s SessionData) error {
	err := u.conn.Set(s.Key, s.Token, u.exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u UserSession) Remove(key string) error {
	err := u.conn.Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}
