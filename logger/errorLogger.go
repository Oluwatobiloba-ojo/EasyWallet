package logger

import "os"

func ErrorLogger(err error) {
	file, errs := os.OpenFile("../errors/error", os.O_RDWR|os.O_APPEND, 0666)
	if errs != nil {
		return
	}
	_, errs = file.WriteString(err.Error())
	if errs != nil {
		return
	}
}
