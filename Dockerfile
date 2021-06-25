FROM golang:1.16 AS builder
WORKDIR /go/src/app/
COPY . .
RUN make build
MAINTAINER massimo.lizza@skytv.it
FROM alpine
LABEL name="concourse-poc"
LABEL description="Concourse POC"
LABEL maintainer="massimo.lizza@skytv.it"
COPY --from=builder /go/src/app/jvs-cli-demo /go/src/app/jvs-cli-demo
CMD /go/src/app/jvs-cli-demo
