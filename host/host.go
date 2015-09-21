package host

import (
	"io/ioutil"
)

func Hosts() string {
	h, _ := ioutil.ReadFile("/etc/hosts")
	return string(h)
}
