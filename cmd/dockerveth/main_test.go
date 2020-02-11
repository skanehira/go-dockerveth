package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMakeTable(t *testing.T) {
	stdout = &bytes.Buffer{}
	*noColor = true
	rows := [][]string{
		{"2c5c7a5c1804", "veth1ce36c6", "php", "php:latest    ", "docker-php-entrypoint bash"},
	}

	makeTable(rows)

	excepted := strings.TrimLeft(`
+--------------+-------------+-------+----------------+----------------------------+
|  CONTAINER   |    VETH     | NAMES |     IMAGE      |            CMD             |
+--------------+-------------+-------+----------------+----------------------------+
| 2c5c7a5c1804 | veth1ce36c6 | php   | php:latest     | docker-php-entrypoint bash |
+--------------+-------------+-------+----------------+----------------------------+
`, "\n")

	result := stdout.(*bytes.Buffer).String()
	if result != excepted {
		t.Fatalf("want :\n%s\n, got:\n%s", excepted, result)
	}
}

func TestMakePlainText(t *testing.T) {
	stdout = &bytes.Buffer{}
	*noColor = true
	rows := [][]string{
		{"2c5c7a5c1804", "veth1ce36c6", "php", "php:latest", "docker-php-entrypoint bash"},
	}

	makePlainText(rows)

	excepted := "2c5c7a5c1804	veth1ce36c6	php	php:latest	docker-php-entrypoint bash	\n"

	result := stdout.(*bytes.Buffer).String()
	if result != excepted {
		t.Fatalf("want :\n%s\n, got:\n%s", []byte(excepted), []byte(result))
	}

}
