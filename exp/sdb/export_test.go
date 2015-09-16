package sdb

import (
	"github.com/upfluence/goamz/aws"
)

func Sign(auth aws.Auth, method, path string, params map[string][]string, headers map[string][]string) {
	sign(auth, method, path, params, headers)
}
