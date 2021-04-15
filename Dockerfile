FROM golang:1.16.3-alpine as build

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download \
 && apk --no-cache add curl tar xz \
 && curl -L https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz | tar -xJf - --strip-components 1 upx-3.96-amd64_linux/upx \
 && ls

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build go build -ldflags "-s -w" -o /out/server . \
    && ls -lh /out/server \
    && ./upx --brute /out/server \
    && ls -lh /out/server


FROM scratch

COPY --from=build /out/server /

CMD ["/server"]
