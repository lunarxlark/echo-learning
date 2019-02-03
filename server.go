package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// log
	//file, err := os.OpenFile("access.log", os.O_WRONLY|os.O_CREATE, 0644)
	file, err := os.OpenFile("log/access.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
  e.GET("/users/:id", getUser)
  e.GET("/users/:region/:id", getUser)
  //e.POST("/users", saveUser)
  //e.PUT("/users/:id", updateUser)
  //e.DELETE("/users/:id", deleteUesr)

	// default
	//e.Use(middleware.Logger())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// tsv logformat
		//Format: tsvLogFormat(),
		// ltsv logformat
		Format:           ltsvLogFormat(),
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		//Output:           os.Stdout,
		Output: file,
	}))

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
  // User id from path `/users/:id`
  region := c.Param("region")
  id := c.Param("id")
  return c.String(http.StatusOK, region + ":" + id)
}



func ltsvLogFormat() string {
	var format string
	//format += "time:${time_rfc3339}\t"
	//format += "time:${time_unix}\t"
	format += "time:${time_custom}\t"
	//TODO : LoggerConfigのidが何を表すか
	//format += "id:${id}\t"
	format += "host:${host}\t"
	format += "remote_ip:${remote_ip}\t"
	format += "uri:${uri}\t"
	format += "method:${method}\t"
	//reqのURL.path
	format += "path:${path}\t"
	format += "protocol:${protocol}\t"
	format += "referer:${referer}\t"
	format += "UA:${user-agent}\t"
	format += "status:${status}\t"
	format += "error:${error}\t"
	//format += "latency:${latency}\t"
	format += "latency:${latency}\t"
	format += "latency_human:${latency_human}\t"
	format += "byte_in:${byte_in}\t"
	format += "byte_out:${byte_out}\n"
	return format
}

func tsvLogFormat() string {
	var format string
	format += "${time_rfc3339}\t"
	format += "${host}\t"
	format += "${remote_ip}\t"
	format += "${uri}\t"
	format += "${method}\t"
	format += "${path}\t"
	format += "${protocol}\t"
	format += "${referer}\t"
	format += "${user-agent}\t"
	format += "${status}\t"
	//format += "${error}\t"
	format += "${latency}\t"
	format += "${latency_human}\t"
	format += "${byte_in}\t"
	format += "${byte_out}\n"
	return format
}
