FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# [Optional] Install Protocol Buffers
RUN VERSION="26.0" \
    && curl -OL "https://github.com/protocolbuffers/protobuf/releases/download/v${VERSION}/protoc-${VERSION}-linux-x86_64.zip" \
    && unzip protoc-${VERSION}-linux-x86_64.zip -d protoc \
    && mv protoc/bin/* /usr/local/bin/ \
    && mv protoc/include/* /usr/local/include/

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1