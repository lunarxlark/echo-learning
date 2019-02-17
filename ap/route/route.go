package route

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/lunarxlark/echo-learning/ap/db"
)

type user struct {
	User_id      int    `json:"user_id"`
	Name_first   string `json:"name_first"`
	Name_familly string `json:"name_familly"`
}

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Path Parameters
	e.GET("/user/:id", getUser)
	// Query Parameters
	e.GET("/show", show)

	return e
}

func getUser(c echo.Context) error {
	d := db.Conn()

	id := c.Param("id")
	u := user{}
	u.User_id, _ = strconv.Atoi(id)
	d.First(&u)

	return c.JSON(http.StatusOK, u)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
