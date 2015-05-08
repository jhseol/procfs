package procfs

import (
  "reflect"
  "testing"
)

func TestProcMounts(t *testing.T) {
  fs, err := NewFS("fixtures")
  if err != nil {
    t.Fatal(err)
  }

  p, err := fs.NewProc(26231)
  if err != nil {
    t.Fatal(err)
  }

  m, err := p.NewMounts()
  if err != nil {
    t.Fatal(err)
  }

  want := ProcMounts{
    Mounts: []mount{
      {Device: "rootfs", MountPoint: "/", FSType: "rootfs", Options: "rw", Dump: 0, FSCK: 0},
      {Device: "proc", MountPoint: "/proc", FSType: "proc", Options: "rw,relatime", Dump: 0, FSCK: 0},
      {Device: "sysfs", MountPoint: "/sys", FSType: "sysfs", Options: "rw,seclabel,relatime", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/cpuset", FSType: "cgroup", Options: "rw,relatime,cpuset", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/cpu", FSType: "cgroup", Options: "rw,relatime,cpu", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/cpuacct", FSType: "cgroup", Options: "rw,relatime,cpuacct", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/memory", FSType: "cgroup", Options: "rw,relatime,memory", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/devices", FSType: "cgroup", Options: "rw,relatime,devices", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/freezer", FSType: "cgroup", Options: "rw,relatime,freezer", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/net_cls", FSType: "cgroup", Options: "rw,relatime,net_cls", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/blkio", FSType: "cgroup", Options: "rw,relatime,blkio", Dump: 0, FSCK: 0},
    },
  }

  if !reflect.DeepEqual(want, m) {
    t.Errorf("want mounts %v to equal %v", want, m)
  }
}

func TestProcMountsCgroup(t *testing.T) {
  m, err := testProcMounts(26231)
  if err != nil {
    t.Fatal(err)
  }

  want := ProcMounts{
    Mounts: []mount{
      {Device: "cgroup", MountPoint: "/cgroup/cpuset", FSType: "cgroup", Options: "rw,relatime,cpuset", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/cpu", FSType: "cgroup", Options: "rw,relatime,cpu", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/cpuacct", FSType: "cgroup", Options: "rw,relatime,cpuacct", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/memory", FSType: "cgroup", Options: "rw,relatime,memory", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/devices", FSType: "cgroup", Options: "rw,relatime,devices", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/freezer", FSType: "cgroup", Options: "rw,relatime,freezer", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/net_cls", FSType: "cgroup", Options: "rw,relatime,net_cls", Dump: 0, FSCK: 0},
      {Device: "cgroup", MountPoint: "/cgroup/blkio", FSType: "cgroup", Options: "rw,relatime,blkio", Dump: 0, FSCK: 0},
    },
  }

  got := m.Cgroups()

  if !reflect.DeepEqual(want, got) {
    t.Errorf("want mounts %v to equal %v", want, m)
  }
}

func testProcMounts(pid int) (ProcMounts, error) {
  p, err := testProcess(pid)
  if err != nil {
    return ProcMounts{}, err
  }

  return p.NewMounts()
}
