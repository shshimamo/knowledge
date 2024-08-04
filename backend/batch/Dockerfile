ARG GO_VERSION=1.22.5
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

COPY . .

RUN go build -o /go/bin/echo cmd/echo.go
RUN go build -o /go/bin/show_user cmd/show_user.go
RUN go build -o /go/bin/exec_query cmd/exec_query.go

FROM alpine:latest AS final

# Install any runtime dependencies that are needed to run your application.
# Leverage a cache mount to /var/cache/apk/ to speed up subsequent builds.
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

# Copy the executable from the "build" stage.
COPY --from=build /go/bin/echo /go/bin/
COPY --from=build /go/bin/show_user /go/bin/
COPY --from=build /go/bin/exec_query /go/bin/

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# What the container should run when it is started.
ENTRYPOINT [ "/entrypoint.sh" ]
