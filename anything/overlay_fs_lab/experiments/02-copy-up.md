## Objective

Understand how OverlayFS handles writes to files that exist only in the lower layer.

## Background

Initially:

```

Application
│
▼
merged/hello.txt
│
▼
lower/hello.txt

```

Although the application accesses `merged/hello.txt`, the actual file resides in the lower directory tree.

Since the lower layer is read-only, OverlayFS cannot modify it.

Instead, it performs **copy-up**, i.e it copies the file from the lower directory tree to the upper directory tree and modifies the changes there.

## Directory Layout

Before mounting:

```
lower/
hello.txt

upper/
(empty)

merged/
(empty)

```

After mounting:

```
lower/
hello.txt

upper/
(empty)

merged/
hello.txt

```

### Read the File

```bash
cat merged/hello.txt
```

Observation:

The file is read successfully.

Upper layer is still empty.

### Modify the File

```bash
echo "changes in merged" >> merged/hello.txt
```

### Modify the file again

```bash
echo "Second write" >> merged/hello.txt
```

## Observations

- After the first modification:
  - `merged/` now contains the docs folder and the modified hello.txt file.
  - `upper/` only contains the modified hello.txt file.
  - `lower/` is unchanged.

- After the second modification:

No additional copy occurs.

OverlayFS now writes directly to the upper copy.

![modification_result](/images/02_modification_result.png)

- Lower Layer never changes

## Explanation

The lower filesystem is treated as immutable or read-only. Whenever a write operation occurs, instead of modifying it OverlayFS:

1. creates parent directories in the upper layer if necessary
2. copies metadata
3. copies file contents
4. redirects future accesses to the upper copy

This process is called **copy-up**.
