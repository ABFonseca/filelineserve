# File Line Serve
Simple API to serve lines from a static file


## How does the system work?
This API as a single endpoint to read a specific line index from a file that is being served (Swagger documentation in <serverURL>/swagger/index.html)

To make a request just make an HTTP request to <serverURL>/lines/lineIndex


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

If running on windows the script will not work, please do:
```
go run main.go <filePath/fileName.txt>
```
or
```
./filelineserve.exe <filePath/fileName.txt>
```

### Running the automated tests
To run the automated tests just run the test script:
```
./test.sh
```


## How will the system perform

### File size scalling
I was somewhat worried about the file size scalling, but I manually tested with a 1GB, 5GB and 10GB files and requested line index 50k, 500k, 5M and 50M and the results were:
|      | 50000    | 500000   | 5000000  | 50000000 |
|------|----------|----------|----------|----------|
| 1GB  | 0m0.078s | 0m0.172s | 0m0.767s | 0m0.770s |
| 5GB  | 0m0.095s | 0m0.187s | 0m0.815s | 0m2.684s |
| 10GB | 0m0.095s | 0m0.185s | 0m0.814s | 0m5.090s |

The file size doesn't seem to have significant impact, the index of the line to be read does seem to have a considerable and potentialy worrying impact.
There are possible solutions to deal with this

#### Possible solution
At start-up evaluate the number of lines in a file (this would also have the benefit of responding to lines beyond EOF faster).

Break File into multiple files and when fetching, depending on the index I would open the correct sub-file and only find the index from that known point


### Users scalling
It's complex to test this on a personal computer, because if I'm both sending and serving the requests there would be a bottlneck that wouldn't be the same as a life environment.
Ideally I would have this on a server and have a network sending requests to have propper testing.

That said, Go deals with concurrency very well due to it's use of threads and I don't expect this to be a bottleneck. If during testing this would prove to be a problem, and since it's a stateless read-only service, I would suggest to have multiple servers running this API with a load balancer in front of it to redirect on a round-robin to one of the APIs


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
- Doing documentation: 30 minutes
- Finding big enough files to test for performance: More than I'm willing to admit
- Doing automated testing: 45 minutes (had a silly bug on the test code and took me a bit too long to detect it)

## Critiquing my code/choices
There are a few points I don't fully agree with my code:
- For the specific problem presented this code is over-engineerd and over-architectured
  - The reason for my choice is because I wanted to showcase a propper architecture for an API
- Depending on interpretation, the chalenge requests for the return of the API call to be just the string and I return a JSON containg the string
  - Again this was to showcase a more generic API, with organized code. If needed, the change to return just the string is changing 1 line and erasing ~5 lines. 
- With more time and if 5GB or bigger files would be something we would have to deal with I would implement the solution mentioned above in the performance section
- To be able to have the code unit tested I would need to have dependency injection. for this exercise the code was too simple and hard to break into anything smaller than it already is, so in this case it would be hard to test just an isolated part, since there is only one part/responsability
  - that said, for production code i would put the DI any case to be prepared from begging when we add more funcionality, and even to be run on pipelines before going to Testing/Production environments
