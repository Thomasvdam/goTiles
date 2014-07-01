# goTiles

A simple Go program that attempts to find all solutions for a given tileset and board.

Simple excercise to get back in the swing of things and to mess around with testing in Go.

## Usage

The `board.txt` file serves as the primary input method, since it beats having to write a million flags and is somewhat more user friendly than having to type some long obfuscated string with the details of the board. I might add the option for a CL argument string later though.  
The first line should contain the dimensions of the board, sperated by a single space: `width height`.
The remaining lines all specify the details of a single tile type: `amount width height`.  
```
5 5
1 3 3
3 2 2
4 1 1
```
Will be interpreted as a 5x5 grid where 1 3x3, 3 2x2, and 4 1x1 tiles have to be placed.
