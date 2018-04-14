# Go graph

Record of business transaction has next struct: 
```
{
  tid: long,
  body: string,
  timestamp: long,
  params: {
    string -> any
  }
}
```

## short description  

* Records with similar *tid* will be stored to single graph.

* Graphs should be clustered if we many unique transactions.

* Records input: syslog sever

## testing

use for testing:

```
$ go test -v ./module_name
```

## TODO:

* message broker consumer

* neo4j snapshot saver

* Computing/Recomputing edge's weight

* Clustering

## running

```
go install      // in project folder
sudo edu-graph  // sudo for listening tcp port
```

for sending tcp package: 
```
echo -n "package body" | nc localhost 514
echo -e "{\"tid\": 1, \"body\": \"hey\", \"timestamp\": 12, \"params\": {}}" | nc localhost 514
```

# Makefile manual

* launch graph: `make`
* initialization graph: `make init-graph`