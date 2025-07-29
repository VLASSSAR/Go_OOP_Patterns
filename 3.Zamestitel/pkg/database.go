package pkg

// Создадим объект реальной базы данных database
type DataBase struct {
}

func (db DataBase) GetData(user string) ([]string, error) {
	return []string{
		"Строка 1",
		"Последняя строка",
	}, nil
}
