# Update the VARIANT arg in devcontainer.json to pick an Go version
ARG VARIANT=1.14
FROM golang:${VARIANT}

# This Dockerfile adds a non-root user with sudo access. Update the “remoteUser” property in
# devcontainer.json to use it. More info: https://aka.ms/vscode-remote/containers/non-root-user.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Options for common setup script - SHA generated on release
ARG INSTALL_ZSH="true"
ARG UPGRADE_PACKAGES="false"
ARG COMMON_SCRIPT_SOURCE="https://raw.githubusercontent.com/microsoft/vscode-dev-containers/master/script-library/common-debian.sh"
ARG COMMON_SCRIPT_SHA="dev-mode"

# Install needed packages and setup non-root user. Use a separate RUN statement to add your own dependencies.
RUN apt-get update \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends curl ca-certificates 2>&1 \
    && curl -sSL  ${COMMON_SCRIPT_SOURCE} -o /tmp/common-setup.sh \
    && ([ "${COMMON_SCRIPT_SHA}" = "dev-mode" ] || (echo "${COMMON_SCRIPT_SHA} /tmp/common-setup.sh" | sha256sum -c -)) \
    && /bin/bash /tmp/common-setup.sh "${INSTALL_ZSH}" "${USERNAME}" "${USER_UID}" "${USER_GID}" "${UPGRADE_PACKAGES}" \
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Install Go tools
ARG GO_TOOLS_WITH_MODULES="\
    golang.org/x/tools/gopls \
    honnef.co/go/tools/... \
    golang.org/x/tools/cmd/gorename \
    golang.org/x/tools/cmd/goimports \
    golang.org/x/tools/cmd/guru \
    golang.org/x/lint/golint \
    github.com/mdempsky/gocode \
    github.com/cweill/gotests/... \
    github.com/haya14busa/goplay/cmd/goplay \
    github.com/sqs/goreturns \
    github.com/josharian/impl \
    github.com/davidrjenni/reftools/cmd/fillstruct \
    github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    github.com/ramya-rao-a/go-outline \
    github.com/acroca/go-symbols \
    github.com/godoctor/godoctor \
    github.com/rogpeppe/godef \
    github.com/zmb3/gogetdoc \
    github.com/fatih/gomodifytags \
    github.com/mgechev/revive \
    github.com/go-delve/delve/cmd/dlv"
RUN mkdir -p /tmp/gotools \
    && cd /tmp/gotools \
    && export GOPATH=/tmp/gotools \
    # Go tools w/module support
    && export GO111MODULE=on \
    && (echo "${GO_TOOLS_WITH_MODULES}" | xargs -n 1 go get -x )2>&1 \
    # gocode-gomod
    && export GO111MODULE=auto \
    && go get -x -d github.com/stamblerre/gocode 2>&1 \
    && go build -o gocode-gomod github.com/stamblerre/gocode \
    # golangci-lint
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin 2>&1 \
    # Move Go tools into path and clean up
    && mv /tmp/gotools/bin/* /usr/local/bin/ \
    && mv gocode-gomod /usr/local/bin/ \
    && rm -rf /tmp/gotools

ENV GO111MODULE=auto


RUN curl https://get.starport.network/starport@v0.13.2! | bash

RUN mkdir -p /go/src/github.com/vjdmhd/nameservice
RUN chown -R USERNAME /go/src/github.com/vjdmhd/nameservice
RUN cd /go/src/github.com/vjdmhd
RUN starport app github.com/vjdmhd/nameservice --sdk-version="launchpad"
RUN chown -R USERNAME /go/src/github.com/vjdmhd/nameservice
RUN cd ~




# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x github.com/gorilla/mux
# RUN go get -x github.com/spf13/cobra
# RUN go get -x github.com/tendermint/tendermint/abci/types
# RUN go get -x github.com/cosmos/cosmos-sdk/client
# #/context
# #RUN go get -x github.com/cosmos/cosmos-sdk/codec
# #RUN go get -x github.com/cosmos/cosmos-sdk/types
# #RUN go get -x github.com/cosmos/cosmos-sdk/types/module
# #RUN go get -x github.com/cosmos/cosmos-sdk/x/bank
# RUN go get -x github.com/vjdmhd/nameservice/x/nameservice/client/cli
# RUN go get -x github.com/vjdmhd/nameservice/x/nameservice/client/rest
# RUN go get -x github.com/vjdmhd/nameservice/x/nameservice/keeper
# RUN go get -x github.com/vjdmhd/nameservice/x/nameservice/types

# go get - u github.com/cosmos/cosmos-sdk 
# 	github.com/golang/mock 
# 	github.com/gorilla/mux
# 	github.com/spf13/afero 
# 	github.com/spf13/cobra 
# 	github.com/spf13/viper 
# 	github.com/tendermint/go-amino
# 	github.com/tendermint/tendermint
# 	github.com/tendermint/tm-db

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update \
#     && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>