#! /usr/bin/env python3

from PIL import Image, ImageDraw, ImageFont
import argparse
import sys

def main():
    parser = argparse.ArgumentParser(description='Converts ASCII file to image image (Good for converting radare2 graphs)')
    parser.add_argument('X', metavar='X Coord', default = 0, type=int, nargs='?',help='x coordinate for image')
    parser.add_argument('Y', metavar='Y Coord', default = 0, type=int, nargs='?',help='y coordinate for image')
    parser.add_argument('--ascii', metavar='ASCII File', default = "", nargs='?',help='ASCII file input')
    parser.add_argument('--img', metavar='IMG Output File', default = "",  nargs='?',help='IMG file output')
    args = parser.parse_args()

    fontfile="/usr/share/fonts/truetype/liberation/LiberationMono-Regular.ttf"
    fontSize = 18
    xMultiplier = fontSize * 0.63
    yMultiplier = fontSize * 1.07
    startX = 10
    startY = 10

    imgX = args.X
    imgY = args.Y
    asciiX=0
    asciiY=0
    fullString=""
    asciiLines=[]
    asciiLinesLength=[]

    if args.ascii!="":
        #open file to read number of lines & line length
        with open(args.ascii) as f:
            for line in f:
                asciiLines.append(line)
                asciiLinesLength.append(len(line))

        #reopen to read as whole
        with open(args.ascii) as f:
            fullString=f.read()
    else:
        fullString = sys.stdin.read()
        asciiLines = fullString.split("\n")
        for line in asciiLines:
            asciiLinesLength.append(len(line))

    #get text X & Y coords
    asciiY = len(asciiLines)
    asciiX = max(asciiLinesLength)

    #setup imgX and imgY if they are not set by user
    if imgX==0:
        imgX = int(asciiX * xMultiplier + (2 * startX))

    if imgY==0:
        imgY = int(asciiY * yMultiplier + (2 * startY))

    # make a blank image for the text
    txt = Image.new('RGB', (imgX, imgY), (255,255,255))

    # get a font
    fnt = ImageFont.truetype(fontfile, fontSize)
    # get a drawing context
    d = ImageDraw.Draw(txt)

    # draw text, full opacity
    d.text((startX,startY), fullString, font=fnt, fill=(0,0,0))

    if args.img=="":
        txt.show()
    else:
        txt.save(args.img)

if __name__ == "__main__":
    main()
