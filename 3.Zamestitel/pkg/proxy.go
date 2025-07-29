package pkg

import "errors"

type ProxyDatabase struct {
	Users map[string]bool
	Db    *DataBase //к нему будем обращаться если пользователь будет удовлетворять всем необходимым требованиям
}

func (pDb ProxyDatabase) GetData(user string) ([]string, error) {
	if !pDb.Users[user] {
		return nil, errors.New("Недостаточно прав доступа к данным!")
	}
	return pDb.Db.GetData(user)
}
