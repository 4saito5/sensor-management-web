package sign

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	// "html"
	// "strconv"
	// "fmt"
)

type (
	m_users struct {
		Email    string `db:"email" json:"email"`
		Password string `db:"password" json:"password"`
		Name     string `db:"name" json:"name"`
	}
)

var (
	tablename = "m_users"
	seq       = 1
	conn, _   = dbr.Open("mysql", "iot:@tcp(127.0.0.1:3306)/smw", nil)
	sess      = conn.NewSession(nil)
)

func SignIn(c echo.Context) error {
	user := new(m_users)
	if err := c.Bind(user); err != nil {
		return err
	}
	// fmt.Println("email=", user.Email)

	res := new(m_users)
	sess.Select("*").
		From(tablename).
		Where("email = ? AND password = ?", user.Email, user.Password).
		Load(&res)
	return c.JSON(http.StatusOK, res)
}

func SignUp(c echo.Context) error {
	user := new(m_users)
	if err := c.Bind(user); err != nil {
		return err
	}

	_, err := sess.InsertInto("m_users").
		Columns("email", "password", "name").
		Record(user).
		Exec()

	if err != nil {
		return err
	}

	// count, _ := result.RowsAffected()
	// fmt.Println(count)

	return c.JSON(http.StatusOK, user)
}
