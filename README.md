# CS Case Opening Stats Extractor/Analyser

You can use this script(s) to extract and analyse all your case opening stats from CS:GO and CS2.

## Requirements:

To use the extractor you either need docker or podman installed on your system. If you don't know how install either of them on your system please check their documentation.

Currently the docker/podman run commands expect you to be running Linux + Wayland. If you want to run it on different systems you'll have to adjust the docker/podman run command within the docker.sh/podman.sh files!

## Usage Extractor:

clone this repository `git clone https://github.com/loissascha/cs-case-stats-viewer.git`

cd into extractor folder `cd cs-case-stats-viewer/extractor`

start container. If you have docker installed: `./docker.sh`, if you have podman installed: `./podman.sh`

log in with your steam and wait for the data extractor to finish (it will take some time! The browser will close when it's done.)

the extracted data will be located inside the output folder

## Usage Analyser:

Download the latest binary for the analyser from the [Releases](https://github.com/loissascha/cs-case-stats-viewer/releases) page.

Make sure to put the extracted 'unlocked_container.json' file from the Extractor in the same folder as the downloaded binary (do not rename the file).

Run the analyser from a terminal with `./analyser`
