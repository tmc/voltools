all: build/index.html

node_modules/.bin/webpack:
	yarn

.PHONY: build/index.html
build/index.html: node_modules/.bin/webpack
	$(shell npm bin)/react-scripts build

.PHONY: deps
deps:
	yarn

.PHONY: clean
clean:
	rm build/*

.PHONY: realclean
realclean: clean
	rm -rf node_modules
