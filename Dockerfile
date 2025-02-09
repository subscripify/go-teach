FROM subscripifycontreg.azurecr.io/builders/subscripifygolang:1.23.4 AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./


# only copy the needed folders
COPY main.go ./main.go

## CGO_ENABLED is set to 0 so that the binary can run on Alpine
RUN go mod download && go mod verify && go mod tidy

#have the cmd folder as ARG so that we can build the binary for the specific cmd
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/server .



FROM alpine:3.14
# try this if CGO_ENABLED suddenly stops working
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /usr/local/bin
RUN mkdir /publicpem
COPY --from=builder /usr/local/bin/server .



EXPOSE 8080
ENTRYPOINT [ "server" ]
# CMD ["serve"]