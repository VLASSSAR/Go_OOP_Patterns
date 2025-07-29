package pkg

type Database struct {
}

func (db Database) GetData(user string) ([]string, error) {
	return []string{
		"Админка",
		"Рулит епта",
	}, nil
}
