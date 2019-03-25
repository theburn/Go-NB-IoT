package utils

import (
	"os"
	"strings"
	"time"
)

func LogNoticeToFile(str_content string) {
	fd, _ := os.OpenFile("logs/notice.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	fd_time := time.Now().Format("2006-01-02 15:04:05")
	fd_content := strings.Join([]string{fd_time, " | ", str_content, "\n"}, "")
	buf := []byte(fd_content)
	fd.Write(buf)
}
