#!/usr/bin/env python3
# 
# Script to URL encode on CLI
# Author: Samuel Pua (samuel.pua@mwrinfosecurity.com)

import urllib.parse
import sys

if ("--help" in sys.argv) or ("-h" in sys.argv) or ("-?" in sys.argv):
    print("URL Encoder")
    print("Author: Samuel Pua\n")
    print("Usage:")
    print("[1] Pipe input into file")
    print()
    print("Flags:")
    print("--help : Prints this help page")
    print("--ignore-newline: Ignores newline while processing as a single request")
    print("--enforce-newline: Generates url code for newline instead of treating as seperate requests")
    print("--persistent: Persistent url encoding")
elif ("--persistent" in sys.argv):
    print("You're in persistent mode. Press Control-C to exit once you're done.\n")
    while True:
        try:
            before=input()
        except KeyboardInterrupt:
            print("Program exiting...")
            sys.exit(0)
        after=urllib.parse.urlencode({'str':before})
        after=after[4:]
        print(after)
        print()
else:
    before=sys.stdin.read()
    if ("--ignore-newline" in sys.argv):
        before=before.replace("\r\n", "")
        before=before.replace("\n", "")
        beforeList=[before]
    elif ("--enforce-newline" in sys.argv):
        beforeList=[before]
    else:
        beforeList=before.split("\n")

    for before in beforeList:
        after=urllib.parse.urlencode({'str':before})
        after=after[4:]
        print(after)

