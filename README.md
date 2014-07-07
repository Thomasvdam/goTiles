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

## To Do

* [x] Refactor the testing files. They're messy and clunky, it must be possible to streamline this.
	* [x] Stuff literal vars in functions so they're not global.
* [ ] Track tile placement.
* [ ] Add backtracking.
* [ ] Write the algorithm that places tiles.
* [ ] Make it spawn several workers to speed up the process.

## Thoughts

For testing I use the github.com/stretchr/testify/assert package, which helps a lot in reducing the amount of clutter code around assertions. I still need to figure out how to do comparisons (`<=` or `>`) to remove some more clutter.  
However, I did miss support for a test that checks whether a variable has changed. It is not hard to write with the tools that testify/assert provides, but it feels like a feature that should be supported.