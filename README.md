Objectives

You must follow the same instructions as in the first subject but the alignment can be changed.
To change the alignment of the output it must be possible to use a flag --align=<type>, in which type can be :

    center

    left

    right

    justify

    You must adapt your representation to the terminal size. If you reduce the terminal window the graphical representation should be adapted to the terminal size.

    Only text that fits the terminal size will be tested.

    The flag must have exactly the same format as above, any other formats must return the following usage message:

Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard

If there are other ascii-art optional projects implemented, the program should accept other correctly formatted [OPTION] and/or [BANNER].
Additionally, the program must still be able to run with a single [STRING] argument.
Instructions

    Your project must be written in Go.
    The code must respect the good practices.
    It is recommended to have test files for unit testing.
Allowed packages

    Only the standard Go packages are allowed.

This project will help you learn about :

    The Go file system(fs) API
    Data manipulation
    Terminal display
