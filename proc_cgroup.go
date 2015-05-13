package procfs

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "strings"
)

// ProcCgroup provide status information about the process,
// read from /proc/[pid]/cgroup
type ProcCgroup struct {
  Cgroup []cgroup
}

type cgroup struct {
  // hierarchy ID number
  ID int64
  // set of subsystems bound to hierarchy
  Subsystems []string
  // control group in the hierarchy to which the process belongs
  ControlGroup string
}

// NewCgroup returns the current cgroup information of the process
func (p Proc) NewCgroup() (ProcCgroup, error) {
  f, err := p.open("cgroup")
  if err != nil {
    return ProcCgroup{}, err
  }
  defer f.Close()

  data, err := ioutil.ReadAll(f)
  if err != nil {
    return ProcCgroup{}, err
  }

  pc := ProcCgroup{}
  cg := bytes.Split(data, []byte("\n"))

  var subsystems string
  for _, b := range cg {
		b = bytes.Replace(b, []byte(":"), []byte(" "), -1)
    c := cgroup{}
    _, err := fmt.Fscan(
      bytes.NewBuffer(b),
      &c.ID,
      &subsystems,
      &c.ControlGroup,
    )
    c.Subsystems = strings.Split(subsystems, ",")
    if err != nil {
      continue
    }
    pc.Cgroup = append(pc.Cgroup, c)
  }

  return pc, err
}
