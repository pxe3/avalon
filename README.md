# Avalon Lang

Avalon Lang is a custom programming language designed specifically for note-taking. This project aims to create a domain-specific language (DSL) that compiles into structured notes, providing features like tagging, linking between notes, and custom sections.

## Project Overview

Avalon Lang is currently in early development. The project includes:

- A lexer for tokenizing Avalon Lang syntax
- A basic GUI for entering notes and viewing lexer output
- (In progress) A parser for building an Abstract Syntax Tree (AST) from tokens

## Features (Planned)

- Intuitive syntax for structured note-taking
- Support for tagging notes
- Ability to link between notes
- Custom sections within notes
- Real-time preview of parsed notes
- Export to common formats (e.g., Markdown, HTML)

## Current Status

- [x] Basic lexer implementation
- [x] Simple GUI for input and lexer output
- [ ] Parser implementation (in progress)
- [ ] AST representation of notes
- [ ] Interpreter for executing/rendering notes
- [ ] Advanced GUI with real-time preview

## Getting Started

### Prerequisites

- Go 1.16 or later
- Fyne toolkit

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/avalon-lang.git
   cd avalon-lang
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Application

To run the current version of Avalon Lang:

```
go run main.go
```

This will launch a GUI where you can enter note text and see the lexer output.
