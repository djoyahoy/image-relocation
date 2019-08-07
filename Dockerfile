FROM golang as builder

COPY . /image-relocation
WORKDIR /image-relocation
RUN go build -o irel"./cmd/irel"

FROM alpine
COPY --from=builder "/image-relocation/irel" "/irel"
