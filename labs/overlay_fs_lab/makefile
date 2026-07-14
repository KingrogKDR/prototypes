LAB := $(CURDIR)/lab

LOWER := $(LAB)/lower
UPPER := $(LAB)/upper
WORK := $(LAB)/work
MERGED := $(LAB)/merged

.PHONY: help setup mount unmount status inspect clean reset

help:
	@echo "Available targets:"
	@echo "  make setup    - Create lab directories"
	@echo "  make mount    - Mount OverlayFS"
	@echo "  make umount   - Unmount OverlayFS"
	@echo "  make status   - Show mount status"
	@echo "  make inspect  - Inspect the directory layouts"
	@echo "  make clean    - Remove generated files"
	@echo "  make reset    - Clean + setup"

setup:
	@mkdir -p $(LOWER) $(UPPER) $(WORK) $(MERGED)

	@mkdir -p $(LOWER)/docs

	@echo "Hello OverlayFS" > $(LOWER)/hello.txt
	@echo "This is documentation." > $(LOWER)/docs/readme.md

	@echo "Lab initialized."

mount:
	@sudo mount -t overlay overlay \
	-o lowerdir=$(LOWER),upperdir=$(UPPER),workdir=$(WORK) \
	$(MERGED)

	@echo "OverlayFS mounted."

umount:
	@sudo umount $(MERGED)

	@echo "OverlayFS unmounted."

status:
	@mount | grep "$(MERGED)" || echo "OverlayFS is not mounted."

inspect:
	@echo
	@echo "Lower"
	@tree $(LOWER)
	@echo
	@echo "Upper"
	@tree $(UPPER)
	@echo
	@echo "Merged"
	@tree $(MERGED)

clean:
	@sudo umount $(MERGED) 2>/dev/null || true

	@rm -rf $(LOWER)/*
	@rm -rf $(UPPER)/*
	@rm -rf $(WORK)/*
	@rm -rf $(MERGED)/*

	@echo "Clean complete."

reset: clean setup

