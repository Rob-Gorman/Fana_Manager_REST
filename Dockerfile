# build static React app files
FROM node AS static

WORKDIR /reactapp
COPY ./_ui .
RUN npm i
RUN npm run build

# build Go API binary
FROM golang AS build
WORKDIR /go/src/manager
COPY ./manager .
COPY --from=static /reactapp/build ./cmd/build
RUN CGO_ENABLED=0 go build -o /fanamanager ./cmd/main.go

# final container with binary
FROM scratch
COPY --from=build /fanamanager /fanamanager
ENTRYPOINT [ "/fanamanager" ]