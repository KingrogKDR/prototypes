#!/bin/bash
# watch_cgroup.sh
#
# Live dashboard for a single cgroup v2 directory.
# Run this in one terminal while you run test workloads in another,
# so you can watch the real kernel state change in response to what you do.
#
# Usage:
#   sudo ./watch_cgroup.sh /sys/fs/cgroup/mytest
#   sudo ./watch_cgroup.sh              # defaults to /sys/fs/cgroup/mytest

CG="${1:-/sys/fs/cgroup/mytest}"

if [ ! -d "$CG" ]; then
  echo "Cgroup directory not found: $CG"
  echo "Create it first, e.g.:"
  echo "  sudo mkdir $CG"
  exit 1
fi

read_file() {
  # Safely read a file, print "-" if it doesn't exist (some controllers/files
  # only appear once a controller is enabled, e.g. cpu.stat needs cpu enabled)
  if [ -r "$1" ]; then
    cat "$1" 2>/dev/null | tr '\n' ' '
  else
    echo "-"
  fi
}

while true; do
  clear
  echo "======================================================"
  echo " cgroup: $CG"
  echo " (ctrl-c to stop)"
  echo "======================================================"
  echo
  echo "--- membership ---"
  echo "processes (cgroup.procs):  $(read_file "$CG/cgroup.procs")"
  echo "controllers enabled here:  $(read_file "$CG/cgroup.controllers")"
  echo "controllers active:        $(read_file "$CG/cgroup.subtree_control")"
  echo
  echo "--- memory ---"
  echo "memory.max:      $(read_file "$CG/memory.max")"
  echo "memory.current:  $(read_file "$CG/memory.current")"
  echo "memory.events:   $(read_file "$CG/memory.events")"
  echo
  echo "--- cpu ---"
  echo "cpu.max:         $(read_file "$CG/cpu.max")"
  echo "cpu.weight:      $(read_file "$CG/cpu.weight")"
  echo "cpu.stat:        $(read_file "$CG/cpu.stat")"
  echo
  echo "--- pids ---"
  echo "pids.max:        $(read_file "$CG/pids.max")"
  echo "pids.current:    $(read_file "$CG/pids.current")"
  echo
  sleep 1
done
