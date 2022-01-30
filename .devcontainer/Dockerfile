ARG VARIANT="1.17-bullseye"

FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# This Dockerfile adds a non-root user with sudo access. Use the "remoteUser"
# property in devcontainer.json to use it. On Linux, the container user's GID/UIDs
# will be updated to match your local UID/GID (when using the dockerFile property).
# See https://aka.ms/vscode-remote/containers/non-root-user for details.
RUN wget https://github.com/exercism/cli/releases/download/v3.0.13/exercism-3.0.13-linux-x86_64.tar.gz
RUN tar -xf exercism-3.0.13-linux-x86_64.tar.gz
RUN mkdir -p ~/bin
RUN mv exercism ~/bin
RUN echo 'export PATH=~/bin:$PATH' >> ~/.bash_profile
