# build static React app files
FROM node AS static

WORKDIR /reactapp
COPY ./dashboard .
RUN npm i
RUN npm run build

# build Go API binary
FROM golang AS build

WORKDIR /go/src/manager
COPY ./manager .
COPY --from=static /reactapp/build ./static
RUN go mod download
RUN CGO_ENABLED=0 go build -o /fanamanager

# final container with binary
FROM scratch

COPY --from=build /fanamanager /fanamanager
ENTRYPOINT [ "/fanamanager" ]