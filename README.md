 ## Go Word Counter
 A simple command line utility to count the frequency of words for a given number of files or text files in a directory 
 recursively. 
 

 ## Includes
 - A synchronous version which process the files sequentially.
 - A concurrent version which forks a goroutine for each file read and then aggregating the results at the end.
 - A directory option to scan a directory and all of its sub directories to find text files and return the results.
 - A profile analysis of the two approaches is also included.

 ### Usage
 Run bin/main.exe with the file paths separated by spaces. Below flags can be specified at execution.
 - -i : will ignore the case when counting frequency
 - -asc : will print the results in an ascending order. default is false
 - -n: specify the number of elements to be printed. default is 20
 - -d: use the directory mode
 