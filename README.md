# Go Word Counter

This program works on the command line.

After building the program:

    go build main.go getWordCount.go convertToString.go

you can run it with ./main, followed by two arguments.

    ./main [dir] [min]
    
The program will read all the text files (in .odt, .doc, .docx or .pdf formats) that exist within a given directory [dir].

It will then create a new file in [dir] containing a list of all the words in each file that repeated more than [min] times.
