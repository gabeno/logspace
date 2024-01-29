### what is LogSpace?

LogSpace is a commit log - an append only sequence of records ordered by time.

Record - data store in our log
Store - file we store our records in
Index - file we store index entries in
Segment - abstraction that ties an index and a store together
Log - abstraction that ties all segments together