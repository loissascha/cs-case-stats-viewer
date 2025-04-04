#!/bin/bash
mkdir -p ./output

docker build -t my-puppeteer-app .
docker run --rm -it \
  --env WAYLAND_DISPLAY=$WAYLAND_DISPLAY \
  --env XDG_RUNTIME_DIR=$XDG_RUNTIME_DIR \
  -v $XDG_RUNTIME_DIR/$WAYLAND_DISPLAY:$XDG_RUNTIME_DIR/$WAYLAND_DISPLAY \
  -v "$(pwd)/output:/app/output" \
  --device /dev/dri \
  --security-opt seccomp=unconfined \
  my-puppeteer-app
