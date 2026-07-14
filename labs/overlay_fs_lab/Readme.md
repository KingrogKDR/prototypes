# OverlayFS Lab

A collection of reproducible experiments that explore the behavior of Linux OverlayFS through observation and hands-on experimentation.

This repository is not a tutorial on using OverlayFS.

Rather than explaining how to use OverlayFS, this repository investigates **why** it behaves the way it does by reproducing common scenarios and analyzing their results.

## Prerequisites:

We'll need:

- A Linux machine or VM with OverlayFS support
- Root privileges (sudo)
- `tree` command

Verify OverlayFS support:

```bash
grep overlay /proc/filesystems
```

Expected:

```
nodev overlay
```

## Setup

This lab lives inside the `prototypes` repository. If you only want this lab locally, use Git's sparse checkout feature.

```bash
git clone --filter=blob:none --sparse \
    https://github.com/KingrogKDR/prototypes.git

cd prototypes

git sparse-checkout set labs/overlay_fs_lab

cd labs/overlay_fs_lab
```

Create the lab environment.

```bash
make setup
```

Display the available commands:

```bash
make help
```

## Cleanup

Unmount OverlayFS before deleting the lab directories.

```bash
make clean
```

## Experiments

| #   | Experiment  | Concept                           |
| --- | ----------- | --------------------------------- |
| 1   | Basic Mount | Creating a merged filesystem      |
| 2   | Copy-up     | First write to a lower-layer file |

The experiments are designed to be completed in order.

## References

- [Linux Kernel OverlayFS Documentation](https://www.kernel.org/doc/html/latest/filesystems/overlayfs.html)
- `man 8 mount`
- `man 2 stat`
- `man 7 mount_namespaces`
