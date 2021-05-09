# CLOUDO
The goal of this project is to build a easy to use distributed loging system, which are also known as write-ahead-logs or transaction logs or commit logs, these kind of logs are the heart of storage engine, message queue, version controle, replication and consensus algorithms.

##Building a Log
The distibuted logging system will be build using the buttom-up approach. 
starting with the store and index file, then the segment and after  which the log. this will enable writting of and running test as each piece is been built
```bash
Store--the file we store records in.
Index--the file we store index entries in.
Segment--the abstraction that ties the store and index together.
Log--the abstraction that ties all segments together.
```


