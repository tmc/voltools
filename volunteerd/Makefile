SQLINSTANCE?=tmcdev:us-central1:test-pg-instance-1
PGUSER?=geodb
PGPASSWORD?=changeme
PGDATABASE?=geodb

.PHONY: all
all: volunteerd ui

.PHONY: lint
lint:
	@command -v staticcheck 2>&1 > /dev/null || go get honnef.co/go/tools/cmd/staticcheck
	staticcheck ./...

.PHONY: volunteerd
volunteerd:
	go install -v ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: ui
ui:
	make -C volunteer-ui

.PHONY: deps
deps:
	sh -c 'command -v brew > /dev/null || /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"'
	sh -c 'command -v go > /dev/null || brew install go'
	sh -c 'command -v shp2pgsql > /dev/null || brew install postgis'
	sh -c 'command -v migrate > /dev/null || go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate'

shapefiles/tl_2019_us_cd116.shp: shapefiles/tl_2019_us_cd116.zip
	sh -c 'cd shapefiles; unzip tl_2019_us_cd116.zip'

shapefiles/tl_2019_us_cd116.zip: shapefiles/tl_2019_us_county.zip
	curl -o $@ https://www2.census.gov/geo/tiger/TIGER2019/CD/tl_2019_us_cd116.zip

shapefiles/tl_2019_us_county.zip:
	curl -o $@ https://www2.census.gov/geo/tiger/TIGER2019/COUNTY/tl_2019_us_county.zip

shapefiles/tl_2019_us_county.shp: shapefiles/tl_2019_us_county.zip
	sh -c 'cd shapefiles; unzip tl_2019_us_county.zip'

shapefiles/tl_2019_us_zcta510.zip:
	curl -o $@ https://www2.census.gov/geo/tiger/TIGER2019/ZCTA5/tl_2019_us_zcta510.zip

shapefiles/tl_2019_us_zcta510.shp: shapefiles/tl_2019_us_zcta510.zip
	sh -c 'cd shapefiles; unzip tl_2019_us_zcta510.zip'

db/migrations/000002_congressional_districts.up.sql: shapefiles/tl_2019_us_cd116.shp
	shp2pgsql shapefiles/tl_2019_us_cd116.shp congressional_districts > db/migrations/000002_congressional_districts.up.sql

db/migrations/000003_counties.up.sql: shapefiles/tl_2019_us_county.shp
	shp2pgsql shapefiles/tl_2019_us_county.shp us_counties > db/migrations/000003_counties.up.sql

db/migrations/000005_zip_codes.up.pgsql: shapefiles/tl_2019_us_zcta510.shp
	shp2pgsql -D $< zip_codes > $@

db/migrations/000006_precincts.up.sql:
	shp2pgsql shapefiles/Precincts041714.shp precincts > db/migrations/000006_precincts.up.sql

cloud_sql_proxy:
	curl -so cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.darwin.amd64
	chmod +x cloud_sql_proxy

cloud_sql_proxy-linux:
	curl -so cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64
	chmod +x cloud_sql_proxy

.PHONY: run-proxy
run-proxy:
	./cloud_sql_proxy -instances=${SQLINSTANCE}=tcp:5433

.PHONY: run-dbshell
run-dbshell:
	@psql postgres://${PGUSER}:${PGPASSWORD}@127.0.0.1:5433/${PGDATABASE}

.PHONY: migrate
migrate:
	migrate -path db/migrations -database postgres://${PGUSER}:${PGPASSWORD}@localhost:5433/${PGDATABASE}?sslmode=disable up

.PHONY: migrate-down
migrate-down:
	migrate -path db/migrations -database postgres://${PGUSER}:${PGPASSWORD}@localhost:5433/${PGDATABASE}?sslmode=disable down -all

.PHONY: migrate-reset
migrate-reset:
	migrate -path db/migrations -database postgres://${PGUSER}:${PGPASSWORD}@localhost:5433/${PGDATABASE}?sslmode=disable force 1

