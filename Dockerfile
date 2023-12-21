FROM ubuntu:latest

WORKDIR /home/app

RUN apt update -y && apt upgrade -y

# Update ca-certificates
RUN apt-get install -y ca-certificates && \
    update-ca-certificates

# Install golang
RUN apt install -y golang-go

# Install dependencies
RUN apt install -y \
        gcc \
        libc6-dev \
        libx11-dev \
        xorg-dev \
        libxtst-dev \
        xsel \
        xclip \
        libpng++-dev \
        xcb \
        libxcb-xkb-dev \
        x11-xkb-utils \
        libx11-xcb-dev \
        libxkbcommon-x11-dev \
        libxkbcommon-dev \
        mingw-w64

COPY . .

RUN go mod download && go mod verify

# Compile Windowsx64 binaries
CMD ["sh", "-c", "GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o ./bin/mouse_mover-0.1.0.windows-amd64.exe ./cmd/mouse_mover/main.go"]