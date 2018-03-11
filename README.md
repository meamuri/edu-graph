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
