package main

import (
	"3.Zamestitel/pkg"
	"fmt"
)

var (
	admin = "administrator"
	user  = "user"
	users = map[string]bool{
		admin: true,
		user:  false,
	}
)

func main() {
	proxy := pkg.ProxyDatabase{
		Users: users,
		Db:    &pkg.DataBase{},
	}
	adminData, err := proxy.GetData(admin)
	fmt.Printf("From [%s] Data: [%v] Err:[%v]\n", admin, adminData, err)
	userData, err := proxy.GetData(user)
	fmt.Printf("From [%s] Data: [%v] Err:[%v]\n", user, userData, err)

}

//Получается у нас прокси-объект контролирует процесс обращения с реальным объектом (Database) при помощи прав доступа по пользователю.
//Паттерн заместитель позволяет нам создать контроль над реальным объектом, то есть все запросы, которые будут приходить к реальному объету будут проходить через
//заместителя и мы можем заместителя создать какой-то функционал по контролю трафика к реальному объекту.
