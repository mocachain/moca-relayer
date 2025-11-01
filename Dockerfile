FROM golang:1.22.4-bullseye AS builder

WORKDIR /workspace

ENV CGO_CFLAGS="-O -D__BLST_PORTABLE__"
ENV CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"

ENV GOPRIVATE=github.com/MocaFoundation

ARG GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}:@github.com/".insteadOf "https://github.com/"

COPY . .

RUN  make build

FROM golang:1.22.4-bullseye

WORKDIR /app

RUN apt-get update -y && apt-get install ca-certificates jq -y

COPY --from=builder /workspace/build/moca-relayer  /usr/bin/moca-relayer

CMD ["moca-relayer"]