package utils

import (
	"strconv"

	uuid "github.com/satori/go.uuid"
)


func IsValidUUID(str string) bool {
	_, err := uuid.FromString(str)
	return err == nil
}

func IsValidNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
