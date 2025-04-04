#!/bin/bash
mkdir -p ./output

podman build -t my-puppeteer-app .
podman run --rm -it \
  --env WAYLAND_DISPLAY=$WAYLAND_DISPLAY \
  --env XDG_RUNTIME_DIR=$XDG_RUNTIME_DIR \
  --env-file .env \
  -v $XDG_RUNTIME_DIR/$WAYLAND_DISPLAY:$XDG_RUNTIME_DIR/$WAYLAND_DISPLAY \
  -v "$(pwd)/output:/app/output" \
  --device /dev/dri \
  --security-opt seccomp=unconfined \
  --privileged \
  my-puppeteer-app
