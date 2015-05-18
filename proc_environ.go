package procfs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
)

// ProcEnviron provide status information about the process, read from /proc/[pid]/environ
type ProcEnviron map[string]string

var (
	wordFilter = regexp.MustCompile("^[\\w]")
)

func (p Proc) NewEnviron() (ProcEnviron, error) {
	f, err := p.open("environ")
	if err != nil {
		return ProcEnviron{}, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	pe := ProcEnviron{}

	envs := bytes.Split(data, []byte("\000"))
	for _, env := range envs {
		// Ignore an empty character and a newline special character
		if wordFilter.Match(env) {
			kv := bytes.SplitN(env, []byte("="), 2)
			switch len(kv) {
			case 1:
				pe[string(kv[0])] = ""
			case 2:
				pe[string(kv[0])] = string(kv[1])
			}
		}
	}

	return pe, err
}
