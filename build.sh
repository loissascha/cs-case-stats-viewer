#!/bin/bash
tsup scraper.ts --format cjs --out-dir dist
pkg dist/scraper.js --targets node18-linux
