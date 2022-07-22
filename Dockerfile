# dev
FROM golang:1.18-bullseye AS dev
WORKDIR /work/src

COPY ./ ./
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin && \
    go install github.com/swaggo/swag/cmd/swag@latest
RUN make export mod build

# release
FROM alpine:latest AS release
RUN apk --no-cache add tzdata libc6-compat && \
    ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=dev /work/src/build/app /usr/local/bin/app
EXPOSE 8080
ENTRYPOINT [ "app" ]
