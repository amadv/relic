.PHONY: build clean install run test

BINARY=relic
PREFIX?=$(HOME)/.local
INSTALL_PATH=$(PREFIX)/bin

build:
	go build -o $(BINARY) .

clean:
	rm -f $(BINARY)
	rm -rf public/

install: build
	install -d $(INSTALL_PATH)
	install -m 755 $(BINARY) $(INSTALL_PATH)/

run: build
	./$(BINARY) example public

test: build
	./$(BINARY) example public
	@echo "✓ Site built successfully in public/"
