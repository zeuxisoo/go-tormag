# Tormag

A command line application to convert magnet and find large file from the torrent directory or file

## Usage

Disable the GO module

    export GO11MODULE=off

Install the vendor (if not exists)

    dep ensure

Install the tools

    go get github.com/jteeuwen/go-bindata/...

Build the application

    make build

Pack the application

    make pack

Release the application (build + pack)

    make release

## Convert

For file

    tormag -f /path/to/your/torrent/file -o /path/to/store/magnet.txt

For directory

    tormag -d /path/to/your/torrent/directory -o /path/to/store/magnet.txt

## Find

For file

    tormag -f /path/to/your/torrent/file -o /path/to/store/biggers.txt

For directory

    tormag -d /path/to/your/torrent/directory -o /path/to/store/biggers.txt

## Web UI

Run the command to start the web application

    tormag web
