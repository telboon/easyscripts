#!/usr/bin/env python3
# 
# Script to URL encode on CLI
# Author: Samuel Pua (samuel.pua@mwrinfosecurity.com)

import urllib.parse
import sys

if ("--help" in sys.argv) or ("-h" in sys.argv) or ("-?" in sys.argv):
    print("URL Decoder")
    print("Author: Samuel Pua\n")
    print("Usage:")
    print("[1] Pipe input into file")
    print()
    print("Flags:")
    print("--help : Prints this help page")
    print("--persistent: Persistent url encoding")
elif ("--persistent" in sys.argv):
    print("You're in persistent mode. Press Control-C to exit once you're done.\n")
    while True:
        try:
            before=input()
        except KeyboardInterrupt:
            print("Program exiting...")
            sys.exit(0)
        after=before
        after=after.replace("+"," ")
        after=urllib.parse.unquote(after)
        print(after)
        print()
else:
    before=sys.stdin.read()

    beforeList=before.split("\n")
    for before in beforeList:
        after=before
        after=after.replace("+"," ")
        after=urllib.parse.unquote(after)
        print(after)

