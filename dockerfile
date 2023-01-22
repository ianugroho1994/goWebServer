FROM alpine:latest

LABEL maintainer="ardiantonugroho <ardianto.nugroho1994@gmail.com>"

ARG http_port=1325

ENV TZ=Asia/Jakarta \
    PATH="/app:${PATH}"

RUN apk add --update --no-cache \
    sqlite \
    tzdata \
    libc6-compat \
    gcompat \
    ca-certificates \
    bash \
    && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ENV PORT $http_port
ENV DB_HOST mysql
ENV DB_PORT 3306
ENV DB_USER root
ENV DB_PASS H4rdtmann
EXPOSE $http_port

WORKDIR /app

COPY ./build/smartlab /app/
RUN mkdir -p /app/cache
CMD ["./smartlab"]