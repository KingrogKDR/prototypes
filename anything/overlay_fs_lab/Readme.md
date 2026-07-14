# OverlayFS Lab

A collection of reproducible experiments for understanding Linux OverlayFS from first principles.

This repository is not a tutorial on using OverlayFS.

Instead, it investigates how OverlayFS behaves by observing the filesystem before and after different operations.

## Pre-Requisites:

We'll need:

- A Linux machine or VM with OverlayFS support
- Root privileges (sudo)
- OverlayFS support (modern Linux kernels include it)
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

Clone the repository.

```bash
git clone https://github.com/KingrogKDR/overlayfs-lab.git
cd overlayfs-lab
```

Create the lab environment.

```bash
make setup
```

> If you require any help with the available commands, use `make help`

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

- Linux Kernel OverlayFS Documentation
- Man page: `mount(8)`
- Man page: `stat(2)`
