# Generator

## Usage

1. **Build the program:**
   ```bash
   go build -o ./challenges/1brc/bin/generator ./challenges/1brc/generator 
   ```
   ```bash
   cd aoc/challenges/1brc/bin
   sh ./build.sh
   ```
2. **Run the program:**
   ```sh
      ./bin/generator/generator -size <number of records to create> -output <output file>
      ./bin/generator/generator-windows-amd64 -size 1_000_000_000 -output weather_data.csv
   ```