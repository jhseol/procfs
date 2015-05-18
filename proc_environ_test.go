package procfs

import (
	"reflect"
	"testing"
)

func TestProcEnviron(t *testing.T) {
	fs, err := NewFS("fixtures")
	if err != nil {
		t.Fatal(err)
	}

	p, err := fs.NewProc(26231)
	if err != nil {
		t.Fatal(err)
	}

	e, err := p.NewEnviron()
	if err != nil {
		t.Fatal(err)
	}

	want := ProcEnviron{
		"HOSTNAME":           "localhost",
		"SHELL":              "/bin/bash",
		"TERM":               "xterm-color",
		"HISTSIZE":           "1000",
		"QTDIR":              "/usr/lib64/qt-3.3",
		"DUALCASE":           "1",
		"LC_ALL":             "ko_KR.UTF-8",
		"USER":               "root",
		"LANG":               "ko_KR.UTF-8",
		"BIN_SH":             "xpg4",
		"SHLVL":              "1",
		"SUDO_COMMAND":       "/bin/su",
		"HOME":               "/root",
		"LOGNAME":            "root",
		"CVS_RSH":            "ssh",
		"LC_CTYPE":           "UTF-8",
		"LESSOPEN":           "|/usr/bin/lesspipe.sh %s",
		"SUDO_GID":           "500",
		"G_BROKEN_FILENAMES": "1",
	}

	if !reflect.DeepEqual(want, e) {
		t.Errorf("want environ \n%q\nto equal \n%q\n", want, e)
	}
}
