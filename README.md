# Word Counter Script

Word counter script using a map reduce structure. The goal is to practice the map 
reduce pattern. The producer is the chunks of bytes that are read from the file being pushed
on to a channel. Each chunk has its own producer. The consumer is a continous reader of the channel
adding each check result to a counter.

# How to run 

To run the program, clone the repo, navigate to it, and use the following command:

```
go run main.go --filename <path to file>
```
