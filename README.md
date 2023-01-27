# Go Word Counter

This program works on the command line.

You can run it with ./main, followed by two arguments.

    ./main dir (min/word)
    
The program will read all the text files (in .odt, .doc, .docx or .pdf formats) that exist within a given directory *dir* located relative to the program. If there's another directory within the given *dir*, the program will go inside that directory, etc.

If the second argument is a number, it will then create a new file in *dir* containing a list of all the words in each file that repeated more than *min* times.

If the second argument is a word, it will create a new file in *dir* containing the number of times that particular word appears in each file.

![](https://github.com/maxhealy01/word_counter/blob/main/functionality.gif)
