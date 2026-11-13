package Advanced

import (
	"fmt"
	"net/url"
)

func UrlLocator() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.User)
	fmt.Println(u.User.Password())
}
