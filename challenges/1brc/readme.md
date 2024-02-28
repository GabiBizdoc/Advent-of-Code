# 1️⃣🐝🏎️ The One Billion Row Challenge (1BRC) - Golang Solution

## Introduction

This project presents a solution to the One Billion Row Challenge (1BRC) using the Go programming language. The
challenge involves optimizing a program to process temperature data from various weather stations, calculating minimum,
mean, and maximum temperature values per station, and outputting the results sorted alphabetically by station name.

## Problem Statement

The program reads a text file containing temperature measurements for various weather stations. Each row in the file
represents one measurement, formatted as `<station name>;<measurement>`, where the measurement value has exactly one
fractional digit. The program calculates the minimum, mean, and maximum temperature values per station, rounding to one
fractional digit. It then writes the results to an output file in the
format `{<station name>=<min>/<mean>/<max>, ...}`.

## Input Data

The input file contains temperature values for a range of weather stations.
Each row is formatted as `<station name>;<measurement>`, where:
`<station name>` is a UTF-8 string of minimum length 1 character and maximum length 100 bytes, containing
neither `;` nor `\n` characters.
`<measurement>` is a float between -99.9 (inclusive) and 99.9 (inclusive), always with one fractional digit.
There is a maximum of 10,000 unique station names.

# Generator

## Usage

To generate the required data use the provided data generators

   ```sh
   cd challenges/1brc/bin/generator
  ./generator-<platform> -size <number of records to create> -output <output file>
   ```

Supported platforms:

- **Linux (AMD64):** `./generator-linux-amd64`
- **Linux (386):** `./generator-linux-386`
- **Linux (ARM64):** `./generator-linux-arm64`
- **Windows (AMD64):** `./generator-windows-amd64.exe`
- **Windows (386):** `./generator-windows-386.exe`
- **macOS (AMD64):** `./generator-darwin-amd64`
- **macOS (ARM64):** `./generator-darwin-arm64`

1. **Build it yourself:**
   ```sh
   go build -o ./challenges/1brc/bin/generator/ ./challenges/1brc/generator 
   ```
   or
   ```sh
   cd challenges/1brc/bin
   sh ./build.sh
   ```
2. **Run the program:**
   ```sh
   cd challenges/1brc/bin
   ./bin/generator/generator -size <number of records to create> -output <output file>
   ./bin/generator/generator-windows-amd64 -size 1_000_000_000 -output weather_data.csv
   ```