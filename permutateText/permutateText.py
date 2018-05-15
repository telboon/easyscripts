#! /usr/bin/env python3
import sys
import time

filenames=[]
fileHandles=[]
currState=[]
currRead=" "

def showState(state):
    finalStr=""
    for curr in state:
        finalStr+=curr
    return finalStr

if len(sys.argv)==1:
    print("Permutate Text")
    print("Usage:")
    print(sys.argv[0]+" textfile1 textfile2 ...")
    print("This will permutate all the combination of the strings in the nextfiles")
    sys.exit()
else:
    for i in range(1,len(sys.argv)):
        filenames.append(sys.argv[i])

    for i in range(len(filenames)):
        fileHandles.append(open(filenames[i]))
        currState.append("")

    operator=0

    while currRead!="" or operator!=-1:
        currRead=fileHandles[operator].readline()

        if currRead != "":
            if currRead[-1:]=="\n":
                currState[operator]=currRead[:-1]
            else:
                currState[operator]=currRead
        
            if operator==len(filenames)-1:
                print(showState(currState))
            elif operator<len(filenames)-1:
                operator+=1
        else:
            fileHandles[operator]=open(filenames[operator])
            operator-=1

