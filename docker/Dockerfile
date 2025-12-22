FROM ubuntu:22.04

LABEL maintainer="archzks@gmail.com"

RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    make \
    gcc \
    g++ \
    cmake \
    gnupg \
    && rm -rf /var/lib/apt/lists/*

RUN curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
RUN curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | tee /etc/apt/sources.list.d/caddy-stable.list
RUN chmod o+r /usr/share/keyrings/caddy-stable-archive-keyring.gpg
RUN chmod o+r /etc/apt/sources.list.d/caddy-stable.list
RUN apt update
RUN apt install caddy
CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]

