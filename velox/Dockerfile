FROM ghcr.io/facebookincubator/velox-dev:ubuntu-22.04

ARG HTTP_PROXY
ARG HTTPS_PROXY

ENV HTTP_PROXY=${HTTP_PROXY}
ENV HTTPS_PROXY=${HTTPS_PROXY}

RUN apt install -y vim && \
    cd / && \
    rm -rf velox && \
    git clone https://github.com/facebookincubator/velox.git && \
    cd velox && \
    ./scripts/setup-ubuntu.sh && \
    make && \
    cmake --install _build/release && \
    cp -r /velox/velox/external/utf8proc /usr/local/include/velox/external/ && \
    rm -rf /velox/_build && \
    rm -rf /velox/deps-download

ENV HTTP_PROXY=""
ENV HTTPS_PROXY=""