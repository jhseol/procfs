package procfs

import (
  "reflect"
  "testing"
)

func TestProcCgroup(t *testing.T) {
  fs, err := NewFS("fixtures")
  if err != nil {
    t.Fatal(err)
  }

  p, err := fs.NewProc(26231)
  if err != nil {
    t.Fatal(err)
  }

  c, err := p.NewCgroup()
  if err != nil {
    t.Fatal(err)
  }

  want := ProcCgroup{
    Cgroup: []cgroup{
      {ID: 10, Subsystems: []string{"perf_event"}, ControlGroup: "/docker/f81eb5834a547078be5a373f0503327dc99a957709153441f0c631e6917c341b"},
      {ID: 9, Subsystems: []string{"memory"}, ControlGroup: "/docker/f81eb5834a547078be5a373f0503327dc99a957709153441f0c631e6917c341b"},
      {ID: 8, Subsystems: []string{"hugetlb"}, ControlGroup: "/"},
      {ID: 7, Subsystems: []string{"freezer"}, ControlGroup: "/docker/f81eb5834a547078be5a373f0503327dc99a957709153441f0c631e6917c341b"},
      {ID: 6, Subsystems: []string{"devices"}, ControlGroup: "/docker/f81eb5834a547078be5a373f0503327dc99a957709153441f0c631e6917c341b"},
      {ID: 5, Subsystems: []string{"cpuacct", "cpu", "cpuset"}, ControlGroup: "/docker/f81eb5834a547078be5a373f0503327dc99a957709153441f0c631e6917c341b"},
      {ID: 4, Subsystems: []string{"blkio"}, ControlGroup: "/docker/f81eb5834a547078be5a373f0503327dc99a957709153441f0c631e6917c341b"},
    },
  }

  if !reflect.DeepEqual(want, c) {
    t.Errorf("want cgroup %v to equal %v", want, c)
  }
}
