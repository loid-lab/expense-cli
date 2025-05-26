

# 💸 Little Expense Tracker CLI

A lightweight command-line tool to add, update, and delete your personal expenses. Built with Go and Cobra, this project helps you manage your spending efficiently from the terminal.

## 📦 Features

- Add new expenses
- Update existing expenses
- Delete expenses
- Easy-to-use CLI interface

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18 or newer)

### Installation

```bash
git clone https://github.com/loid-lab/expense-cli.git
cd expense-cli
go build -o tracker
```

### Usage

```bash
./tracker add "Lunch" 12.50
./tracker update 1 --amount 15.00
./tracker delete 1
```

> You can extend these commands with flags like `--name`, `--amount`, `--date`, etc., depending on how your CLI is set up.

## 📁 Project Structure

```
.
├── cmd/             # Cobra commands (add, update, delete)
├── main.go          # Entry point
├── go.mod           # Go module definition
└── README.md        # This file
```

## 🛠 Built With

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)

## 🧾 License

This project is licensed under the MIT License.