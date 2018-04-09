# easyscripts
Bunch of scripts that make life easier on command line

Author: Samuel Pua (kahkin@gmail.com)

## Current Scripts Available
* alliteration-gen
     * Creates alliteration in the format of <adjective> <noun>. Mainly used for naming purposes
  
* ascii2img
     * Creates image out of an ASCII text file -* may be used to convert ASCII output from Radare2 to image

     * Possible to edit the text file with annotation before saving

* clearExtraLines
     * Burp Suite creates '\r\n' on its output for it's proxy and repeater. Probably useful for Windows, but not so useful for Linux
     * This creates a problem of 2 lines on Linux
     * Tool converts '\r\n' to '\n' in the clipboard buffer

* fuzzPattern
     * Creates pattern to be used for fuzzing
     * The fuzz output must be more than or equal to 4 bytes (32 bit)

* ratio
     * Converts a ratio to it's smallest common denominator

* urlencode.py
     * Encodes strings into urlencoding

* urldecode.py
     * Decodes url encoding back to ascii


## Todos
- Tab completion
> Seriously out of my league for now. Easy to way to it seems to be using external library argcomplete. But I prefer not to use external libraries
- Add no search scenario for fuzzPatternSearch
> For longer than 4 bytes
> For really cannot find
