FROM debian:stretch-slim

WORKDIR /

COPY bin/wyq-scheduler /usr/local/bin

CMD ["wyq-scheduler"]

# docker build --no-cache -t registry.cnbita.com:5000/wuyiqiang/wyq-scheduler:v1 .