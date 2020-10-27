all: dist/index.html

node_modules/.bin/webpack:
	yarn

.PHONY: build/index.html
build/index.html: node_modules/.bin/webpack
	$(shell npm bin)/parcel build index.html

.PHONY: deps
deps:
	yarn

.PHONY: clean
clean:
	rm dist/*

.PHONY: realclean
realclean: clean
	rm -rf node_modules
