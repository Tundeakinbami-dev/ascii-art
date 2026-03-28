# ASCII-Art Generator (Go)

## Overview

This project is a command-line application written in **Go** that converts regular text into **ASCII art** using predefined banner templates.

Each character from the input text is mapped to a graphical representation made of ASCII characters. The program reads these representations from banner files and prints the result to the terminal.

The supported banner templates are:

- `standard`
- `shadow`
- `thinkertoy`

Each character in a banner file has a **height of 8 lines**, and the ASCII representation is printed line-by-line to construct the final output.

---

# Features

- Converts text into ASCII art
- Supports multiple banner styles
- Handles newline characters (`\n`)
- Supports ASCII characters from **32 (space)** to **126 (~)**
- Clean modular Go structure
- Includes unit tests
- Uses only **standard Go packages**

---

# Project Structure

```
ascii-art/
│
├── main.go
├── go.mod
│
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── internal/
│   └── ascii/
│        ├── banner_loader.go
│        ├── renderer.go
│        └── renderer_test.go
│
└── README.md
```

### Description

| File               | Purpose                             |
| ------------------ | ----------------------------------- |
| `main.go`          | Entry point of the program          |
| `banner_loader.go` | Loads ASCII banner templates        |
| `renderer.go`      | Generates ASCII art from input text |
| `renderer_test.go` | Unit tests                          |
| `banners/`         | ASCII template files                |

---

# Banner File Format

Banner files contain ASCII-art representations for characters.

- ASCII characters from **32 to 126**
- Each character uses **8 lines**
- Characters are separated by an empty line

Example representation:

```
......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....
```

Each dot (`.`) represents a space.

---

# Installation

Clone the repository:

```
git clone <repository-url>
cd ascii-art
```

Initialize the Go module if needed:

```
go mod tidy
```

---

# Usage

Run the program with:

```
go run . "text"
```

Example:

```
go run . "Hello"
```

Example output:

```
 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/
```

---

# Handling New Lines

The program supports newline characters using `\n`.

Example:

```
go run . "Hello\nThere"
```

This prints ASCII art on **multiple lines**.

---

# Running Tests

Unit tests can be executed using:

```
go test ./...
```

Tests verify that the ASCII renderer correctly processes input text and banner files.

---

# Requirements

- Go **1.20+**
- Only **standard Go packages** are used.

---

# Good Practices Implemented

- Modular architecture
- Error handling for file operations
- Efficient string building
- Unit testing
- Clean folder structure

---

# Author

Developed as part of a **Go programming project** focused on:

- File processing
- ASCII manipulation
- Command-line tools
- Clean Go architecture
