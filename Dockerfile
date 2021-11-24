FROM golang:alpine as build

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM scratch as final
COPY --from=build /app/main .
COPY --from=build /app/app.ini .
EXPOSE 8080
CMD ["/main"]