## Objective

Understand how OverlayFS creates a merged filesystem.

## Hypothesis

Files should appear in `merged/` without being copied into `upper/`.

## Commands and steps

1. Verify every directory before moving forward. Initially, lower/ contains the content, upper/ is empty, work/ is empty, merged/ is empty

2. Run:

```bash
make mount
```

3. Now inspect the merged/ directory

```bash
make inspect
```

4. Read a file:

```bash
cat merged/hello.txt
```

It shows: `Hello OverlayFS`

5. Compare inode numbers using:

```bash
stat lower/hello.txt
stat merged/hello.txt
```

6. Now unmount and inspect

```bash
make unmount
make inspect
```

## Observations

After step 2:

- `merged/` now contains the files.
- `upper/` is still empty.
- `lower/` is unchanged.

After step 6:

- `merged/` is empty
- `upper/` is empty
- `lower/` is unchanged

## Explanation

OverlayFS exposes a virtual merged view. Reads are served directly from the lower layer. No changes occur in the lower layer
and the merged view is only a combined abstraction of both the upper and lower layer directory trees.

The upper directory and work directory remain empty because we only performed read operations so far. No write operations occurred.
