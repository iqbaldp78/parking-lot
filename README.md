# parking-lot

This program to make parking-lot

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Assumption
- this program is only temporary because the data is stored in the form of a slice, not store in DB
- there are no dependencies with libraries, that no need go get
- commands used :
1. park
2. leave
3. status
4. registration_numbers_for_cars_with_colour
5. slot_numbers_for_cars_with_colour
6. slot_number_for_registration_number
- please make a parking slot first to run the program
- static .txt file, because this program goes directly to directory ./functional_spec/fixtures/file_input.txt
- if you run the program with os.Args it will go directly to the .txt file
### Prerequisites


```
Install golang https://golang.org/doc/install (if you want to run the program via golang language)

Install golang https://www.docker.com/get-started (if you want running the program via docker)
```




## Running the program
running the program 

```
1. extract file parking-lot.zip
2. enter to parking directory with terminal 'cd parking_lot/parking_lot/bin'
3. .bin/parking_lot (for interactive shell)
4. .bin/parking_lot file_inputs.txt (for input file)
```

running the program via Dockerfile
```
docker build --rm -f "Dockerfile" -t parking_lot:latest .

docker run -d parking_lot
```