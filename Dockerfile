FROM --platform=$TARGETPLATFORM golang as build
WORKDIR /build
COPY . .
RUN go mod tidy && \
  CGO_ENABLED=0 go build -o ./osubot ./cmd


FROM --platform=$TARGETPLATFORM alpine
COPY --from=build /build/osubot /bin/osubot
RUN chmod +x /bin/osubot

LABEL org.opencontainers.image.source https://github.com/EnergoStalin/osubot

CMD [ "osubot" ]