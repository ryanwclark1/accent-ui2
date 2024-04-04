FROM openapitools/openapi-generator-cli:latest

ENV PATH /usr/local/go/bin:$PATH

ENV GOLANG_VERSION 1.22.1

RUN set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
      g++ \
      gcc \
      libc6-dev \
      make \
      pkg-config \
    ; \
    rm -rf /var/lib/apt/lists/*; \
    url='https://dl.google.com/go/go1.22.1.linux-amd64.tar.gz'; \
    sha256='aab8e15785c997ae20f9c88422ee35d962c4562212bb0f879d052a35c8307c7f'; \
    wget -O go.tgz.asc "$url.asc"; \
    wget -O go.tgz "$url" --progress=dot:giga; \
    echo "$sha256 *go.tgz" | sha256sum -c -; \
    tar -C /usr/local -xzf go.tgz; \
    rm go.tgz; \