FROM golang as build
WORKDIR /build
COPY . .
RUN go mod tidy && \
  CGO_ENABLED=0 go build -o ./osubot ./cmd


FROM alpine
COPY --from=build /build/osubot /bin/osubot
RUN chmod +x /bin/osubot
CMD [ "osubot" ]