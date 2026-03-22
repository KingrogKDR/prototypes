# Docker learnings

- Use --pull to get fresh base images

The following Dockerfile uses the 24.04 tag of the ubuntu image. Over time, that tag may resolve to a different underlying version of the ubuntu image, as the publisher rebuilds the image with new security patches and updated libraries.

```bash
FROM ubuntu:24.04
RUN apt-get -y update && apt-get install -y --no-install-recommends python3
```

To get the latest version of the base image, use the --pull flag:

```bash
docker build --pull -t my-image:my-tag .
```

The `--pull` flag forces Docker to check for and download a newer version of the base image, even if you have a version cached locally.

- Use --no-cache for clean builds

The `--no-cache` flag disables the build cache, forcing Docker to rebuild all layers from scratch:

```bash
docker build --no-cache -t my-image:my-tag .
```

This gets the latest available versions of dependencies from package managers like apt-get or npm. However, --no-cache doesn't pull a fresh base image - it only prevents reusing cached layers. For a completely fresh build with the latest base image, combine both flags:

```bash
docker build --pull --no-cache -t my-image:my-tag.
```

