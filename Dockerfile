# ================================
# Build
# ================================
FROM golang:1.21.0 as build

WORKDIR /build

COPY . .

RUN go get

RUN go build -ldflags="-s -w" -trimpath

WORKDIR /staging

RUN cp /build/firebase-authentication-service /staging


# ================================
# Run image
# ================================
FROM ubuntu:jammy

RUN useradd --user-group --create-home --system --skel /dev/null --home-dir /app authenticator

WORKDIR /app

COPY --from=build --chown=authenticator:authenticator /staging /app

USER authenticator:authenticator

EXPOSE 5000

ENTRYPOINT ["./firebase-authentication-service"]
