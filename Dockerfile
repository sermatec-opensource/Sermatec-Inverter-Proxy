FROM golang:1.19-alpine as builder
WORKDIR /build

RUN apk --update add --no-cache \
    	ca-certificates \
    	openssl \
    	git \
    	tzdata \
	&& update-ca-certificates

COPY ./client ./client
COPY ./common ./common
COPY ./mockserver ./mockserver
COPY ./proxyfier ./proxyfier
COPY ./webui ./webui
COPY ./main.go ./
COPY ./go.mod ./
COPY ./go.sum ./

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o SermatecProxy .
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

FROM scratch

COPY --from=builder /build/SermatecProxy /app/
COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

VOLUME /secrets

EXPOSE 80

WORKDIR /app
ENTRYPOINT [ "./SermatecProxy" ]
