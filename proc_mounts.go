package procfs

import (
  "bytes"
  "fmt"
  "io/ioutil"
)

// ProcMounts provide status information about the process,
// read from /proc/[pid]/mounts
type ProcMounts struct {
  Mounts []mount
}

type mount struct {
  // The Device or remote filesystem name to be mounted.
  Device string
  // The mount point for the filesystem.
  MountPoint string
  // The type of the filesystem.
  FSType string
  // The mount options associated with filesystem.
  Options string
  // The flag to determine which filesystem need to be dumped.
  Dump int
  // The flag to determine the order in which filesystem checks are done at reboot time
  FSCK int
}

// NewMounts returns the current mounts information of the process
func (p Proc) NewMounts() (ProcMounts, error) {
  f, err := p.open("mounts")
  if err != nil {
    return ProcMounts{}, err
  }
  defer f.Close()

  data, err := ioutil.ReadAll(f)
  if err != nil {
    return ProcMounts{}, err
  }

  pm := ProcMounts{}

  mb := bytes.Split(data, []byte("\n"))

  for _, b := range mb {
    m := mount{}
    _, err := fmt.Fscan(
      bytes.NewBuffer(b),
      &m.Device,
      &m.MountPoint,
      &m.FSType,
      &m.Options,
      &m.Dump,
      &m.FSCK,
    )
    if err != nil {
      continue
    }
    pm.Mounts = append(pm.Mounts, m)
  }

  return pm, err
}

// Cgroups returns the cgroups of the process
func (p ProcMounts) Cgroups() ProcMounts {
  m := ProcMounts{}

  for _, v := range p.Mounts {
    if v.Device == "cgroup" {
      m.Mounts = append(m.Mounts, v)
    }
  }

  return m
}
