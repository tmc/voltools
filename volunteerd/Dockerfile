FROM golang:1.15-buster as go_layer

WORKDIR /go/src/github.com/HF-RapidResponse/geotools/volunteerd

COPY go.sum go.mod ./
RUN go mod download
COPY ./cmd/ ./cmd/
COPY ./models/ ./models/
RUN go install -v ./...

FROM node:15-buster as js_layer

WORKDIR /app
COPY volunteer-ui/package.json ./
COPY volunteer-ui/yarn.lock ./
RUN yarn
COPY volunteer-ui/ ./
RUN make

FROM golang:1.15-buster as final_layer

COPY --from=go_layer /go/bin/ /go/bin/
COPY --from=js_layer /app/dist ./volunteer-ui/dist/

COPY scripts/entrypoint.sh /usr/local/bin/

CMD ["entrypoint.sh"]

