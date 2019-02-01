package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// default
	//e.Use(middleware.Logger())

	// tlsv logformat
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           tlsvLogFormat(),
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           os.Stdout,
	}))

	// tsv logformat
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: tsvLogFormat(),
	//}))

	e.Logger.Fatal(e.Start(":1323"))
}

func tlsvLogFormat() string {
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
