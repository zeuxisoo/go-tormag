# Tormag

A command line application to convert the torrent directory or file to magnets

## Installation

Install the vendor (if not exists)

    dep ensure

Build the application

    make build

Pack the application

    make pack

Release the application (build + pack)

    make release

## Usage

For file

    tormag -f /path/to/your/torrent/file -o /path/to/store/magnett.txt

For directory

    tormag -d /path/to/your/torrent/directory -o /path/to/store/magnett.txt
