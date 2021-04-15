FROM golang:1.16.3-alpine as build

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build go build -o /out/server .

FROM scratch

COPY --from=build /out/server /

CMD ["/server"]
