FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

ENV PORT 3000
ENV DATABASE_URL postgresql://postgres:postgres@db:5432/letvagas

RUN mkdir -p /build && go build -o /build/letvagas
RUN cp -r /app/templates /build/
RUN cp -r /app/static /build/

FROM alpine
COPY --from=builder /build/ /

WORKDIR /

EXPOSE 3000

CMD [ "./letvagas" ]
