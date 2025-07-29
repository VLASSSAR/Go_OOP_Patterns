package pkg

import "errors"

type ProxyDatabase struct {
	Users map[string]bool
	Db    *Database
}

func (pDb ProxyDatabase) GetData(user string) ([]string, error) {
	if !pDb.Users[user] {
		return nil, errors.New("Нихуя не пройдешь, прав нет xD")
	}
	return pDb.Db.GetData(user)
}
