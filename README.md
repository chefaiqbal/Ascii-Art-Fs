# ASCII Art FS

## Authors

- [tnji](https://learn.reboot01.com/git/tnji)
- [aiqbal](https://learn.reboot01.com/git/aiqbal)

## Objectives

ASCII Art FS is a program that generates graphic representations of input strings using ASCII characters. The program takes a string as input and an optional banner type, and it generates the corresponding ASCII art using the specified banner type. If no banner type is provided, it uses the "standard" banner by default.

## Usage

1. Ensure you have Go installed on your machine.
2. Clone this repository:

git clone https://learn.reboot01.com/git/aiqbal/ascii-art-fs


3. Navigate to the project directory:

cd ascii-art-fs

4. Run the program with an input string and an optional banner type:

go run . "Hello, World!" standard


The program will generate the ASCII art representation of the input string using the specified banner type and display it in the console.

If no banner type is provided, the program will use the "standard" banner by default.

## Features

- Supports input with numbers, letters, spaces, special characters, and line breaks.
- Each character has a height of 8 lines.
- Characters are separated by a new line ("\n").
- Uses separate text files for each banner type to store the ASCII representations of characters.
- Customization: You can add or modify ASCII characters by editing the corresponding banner text file.

## Available Banner Types

The program supports the following banner types:

- Standard: The default banner type.
- Shadow: A banner with a shadow effect.
- Thinkertoy: A banner with a thinkertoy-style design.

## Customization

If you want to customize the ASCII characters or add more characters, you can edit the corresponding banner text files. Each character's ASCII representation should be separated by an empty line. Ensure the width and height of each character match the constant values defined in the code.

## Example

Here is an example of generating ASCII art for the input string "Hello, World!" using the standard banner type:

```
 _    _          _   _                    __          __                 _       _   _  
| |  | |        | | | |                   \ \        / /                | |     | | | | 
| |__| |   ___  | | | |   ___              \ \  /\  / /    ___    _ __  | |   __| | | | 
|  __  |  / _ \ | | | |  / _ \              \ \/  \/ /    / _ \  | '__| | |  / _` | | | 
| |  | | |  __/ | | | | | (_) |  _           \  /\  /    | (_) | | |    | | | (_| | |_| 
|_|  |_|  \___| |_| |_|  \___/  ( )           \/  \/      \___/  |_|    |_|  \__,_| (_) 
                                |/                                                      

```