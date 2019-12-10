package utils

import (
	"log"
	"testing"
)

func TestJoinPath(t *testing.T) {
	log.Println(JoinPaths("/", "/demon/write/"))
	log.Println(JoinPaths("/hello/", "/demon/write"))
	log.Println(insertBackSlash("/////demon/write"))
}
