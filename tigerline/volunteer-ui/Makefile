.PHONY: all
all: dist/index.html

.PHONY: serve
serve:
	$(shell npm bin)/parcel serve index.html

.PHONY: watch
watch:
	$(shell npm bin)/parcel watch index.html

fmt: node_modules/.bin/prettier
	node_modules/.bin/prettier -w .

node_modules/.bin/prettier:
	yarn

node_modules/.bin/parcel:
	yarn

.PHONY: dist/index.html
dist/index.html: node_modules/.bin/parcel
	$(shell npm bin)/parcel build index.html

.PHONY: deps
deps:
	yarn

.PHONY: clean
clean:
	rm -rf dist/*
	rm -rf .cache

.PHONY: realclean
realclean: clean
	rm -rf node_modules
