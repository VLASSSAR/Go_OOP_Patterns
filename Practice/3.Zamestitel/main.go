package main

import (
	"3.Zamestitel/pkg"
	"fmt"
)

var (
	user  = "user"
	admin = "administrator"

	users = map[string]bool{
		user:  false,
		admin: true,
	}

	proxy = pkg.ProxyDatabase{
		Users: users,
		Db:    &pkg.Database{},
	}
)

func main() {
	adminData, err := proxy.GetData(admin)
	fmt.Printf("From: [%s] Data:[%v], Err:[%v]\n", admin, adminData, err)

	userData, err := proxy.GetData(user)
	fmt.Printf("From: [%s] Data:[%v], Err:[%v]\n", user, userData, err)
}
