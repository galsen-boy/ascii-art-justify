
<table align="center" border="0">
  <tr>
    <td><a href="https://imgbb.com/"><img src="https://i.ibb.co/tL6SpfD/zone01.png" alt="zone01" border="0" width="100px"></a></td>
    <td><a href="https://imgbb.com/"><img src="https://i.ibb.co/W2vM8Ws/golang.png" alt="golang" border="0" width="120px"></a></td>
   </tr>
</table>
<h1 align="center">ASCII-ART-JUSTIFY</h1>
*  A small Go program that align ASCII art from a string input. 



## Objectives

You must follow the same instructions as in the first subject but the alignment can be changed.
To change the alignment of the output it must be possible to use a flag --align=<type>, in which type can be :

    center

    left

    right

    justify

    You must adapt your representation to the terminal size. If you reduce the terminal window the graphical representation should be adapted to the terminal size.

    Only text that fits the terminal size will be tested.

    The flag must have exactly the same format as above, any other formats must return the following usage message:

## Installation

 - To use this program simple clone this Git repository on your local machine

```bash
  git clone https://learn.zone01dakar.sn/git/daiba/ascii-art-justify
```
-  Open Terminal and install go package
```bash
  apt install golang
```
-  Run program using the next command
```bash
  go run . [OPTION] [STRING]
```
- USAGE : 
```bash 
go run . "--align=left" "Hello"
```
```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $ 
```
## Author

- [@daiba](https://learn.zone01dakar.sn/git/daiba)
- [@nifaye](https://learn.zone01dakar.sn/git/nifaye)
- [@ymadicke](https://learn.zone01dakar.sn/git/ymadicke)

![Logo](https://go.dev/images/go-logo-white.svg)
