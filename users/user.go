package users

import (
	"fmt"
	"time"

	"github.com/eplacenciatz/godesde0/models"
)

func AltaUsuario() {
	u := new(models.User)
	u.AddUser(1, "José", time.Now(), true)
	fmt.Println(u)
}
