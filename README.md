# CS Case Opening Stats Extractor/Analyser

You can use this script(s) to extract and analyse all your case opening stats from CS:GO and CS2.

## Requirements:

To use the extractor you need either docker or podman installed on your system. If you don't know how install either of them please check the documentation.

Currently the docker/podman commands expect you to be running Linux and Wayland. If you want to run it on different systems you'll have to adjust the docker/podman run command within the docker.sh/podman.sh files!

## Usage Extractor:

clone this repository `git clone https://github.com/loissascha/cs-case-stats-viewer.git`

cd into extractor folder `cd cs-case-stats-viewer/extractor`

start container. If you have docker installed: `./docker.sh`, if you have podman installed: `./podman.sh`

log in with your steam and wait for the data extractor to finish (it will take some time! The browser will close when it's done.)

the extracted data will be located inside the output folder

## Usage Analyser: 

