# File Line Serve
Simple API to serve lines from a static file


## How does the system work?
....

### Requirements
For this code to run you need to have go 1.18 or later installed on your machine

### Building the code
Just run the building script
```
./build.sh
```

### Running the code
To run the API just use the run script and pass the path to the file that you want to be served
```
./run.sh <filePath/fileName.txt>
```

## How will the system perform

### File size scalling
___________


### Users scalling
________________



## What documentation, websites, papers were consulted
I wouldn't be a developer without Stack Overflow so I consulted that one.
I consulted previous API projects done by me to doublecheck on structure/orgnaization of code
I used https://loremipsum.io/generator to generate lines for the example file

## What third-party libraries or other tools does the system use
This API uses:
- Gin framework, this is one of the most common API frameworks in Go
- Swaggo which is a tool to generate the API Swag documentation from the code comments
- Netflix/go-env lib to marshal environment variables into a go struct that can be used in code
- Subosito/gotenv lib to load a .env file into environment-variable

Appart from Gin none of the tools perform a core operation or perform during API request processing, so I just went with libs I knew and used before and 
didn't really compared with other alternative libs that would do the same or similar.


## How long did I spend on this exercise (aproximate values)
- Doing the code (without automatic tests): 2 hours
- Doing documentation: 20 minutes
- Doing automated testing: ???

## Critiquing my code/choices
There are a few points I don't fully agree with my code:
- For the specific problem presented this code is over-engineerd and over-architectured
 - The reason for my choice is because I wanted to showcase a propper architecture for an API
- Depending on interpretation, the chalenge requests for the return of the API call to be just the string and I return a JSON containg the string
 - Again this was to showcase a more generic API, with organized code. If needed, the change to return just the string is changing 1 line and erasing ~5 lines. 
