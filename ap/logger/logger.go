package logger

import (
	"os"
)

func LogFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func LtsvLogFormat() string {
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

func TsvLogFormat() string {
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
