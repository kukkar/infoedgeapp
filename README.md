# infoedgeapp

1. Since there is not formal try catch mechanism in golang except for panic, recover an Error Wrapper was written which reads the output of function, write the error to stdout and returns it.
2. No external testing library has been used as golang provides internal "testing" library. All test cases have been written using it.


## Run
```
$ ./parking_lot
test
....
```
OR
```
$ ./parking_lot abc.txt
output....
```


## ALLOWED COMMANDS

```
 LOGIN <User> <password> 
 SIGNUP <Id> <email> <password> 
 LISTJOURNAL 
 CREATEJOURNAL <message>
 
```