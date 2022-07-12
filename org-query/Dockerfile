FROM 406205545357.dkr.ecr.us-east-1.amazonaws.com/mirror/golang:1.17.8 as build-image

WORKDIR /go/src
COPY go.mod go.sum main.go ./
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates
# add certificate so build deps doesn't fail on protoc-gen-grpc-gateway installation
RUN openssl s_client -showcerts -connect storage.googleapis.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > /usr/local/share/ca-certificates/googleapis.crt && \
  update-ca-certificates

RUN go build -o ../bin

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /var/task/

# Command can be overwritten by providing a different command in the template directly.
CMD ["org-query"]
