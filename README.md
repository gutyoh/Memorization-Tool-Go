# Memorization-Tool-Go
[Hyperskill Memorization Tool project](https://hyperskill.org/projects/159) from the Python/Flask track, reimplemented in Go using GORM.

## Requirements:

Before you get started with the Memorization Tool project, ensure you have the following requirements installed in your machine:

- Go version 1.21.0
- Python 3 (for testing)
- Make (for automation)

If you're using **Windows 🪟**, you'll also need:

- [Git Bash](https://gitforwindows.org/)
- The [Make binary for Windows](http://gnuwin32.sourceforge.net/packages/make.htm)

Having these tools installed will ensure a smooth development and testing experience across different platforms.


## How to run this project with Go solutions

1. **Clone the repository and `cd` into it:**

```bash
git clone https://github.com/gutyoh/Memorization-Tool-Go.git
cd Memorization-Tool-Go
```

2. **Directory structure**: After cloning the repository, the directory tree should look like this:

```
Memorization-Tool-Go
├── go.mod
├── go.sum
├── lesson-info.yaml
├── lesson-remote-info.yaml
├── Makefile
├── README.md
├── requirements.txt
├── stage1
│   ├── main.go
│   ├── task.html
│   ├── task-info.yaml
│   ├── task-remote-info.yaml
│   ├── test
│   │   ├── __init__.py
│   │   └── tests.py
│   └── tests.py
├── stage2
│   ├── main.go
│   ├── task.html
│   ├── task-info.yaml
│   ├── task-remote-info.yaml
│   ├── test
│   │   ├── __init__.py
│   │   └── tests.py
│   └── tests.py
├── stage3
│   ├── main.go
│   ├── task.html
│   ├── task-info.yaml
│   ├── task-remote-info.yaml
│   ├── test
│   │   ├── __init__.py
│   │   └── tests.py
│   └── tests.py
└── stage4
    ├── main.go
    ├── task.html
    ├── task-info.yaml
    ├── task-remote-info.yaml
    ├── test
    │   ├── __init__.py
    │   └── tests.py
    └── tests.py
```

3. **Setup Python Virtual Environment:**

**Linux 🐧 and macOS 🍏:**

```bash
python3 -m venv memo_tool_venv
source memo_tool_venv/bin/activate
```

**Windows 🪟:**

```bash
python -m venv memo_tool_venv
memo_tool_venv\Scripts\activate
```

4. **Install the necessary Python packages:**

```bash
pip install -r requirements.txt
```

After completing the above steps, below is a picture of how the project directory should look like:

![image](https://github.com/gutyoh/Memorization-Tool-Go/assets/8846884/5b125647-1ab9-461a-ad78-b3459aadac21)

5. **From the root of the project directory, run the project using `make`:**
```bash
make
```

Executing the `make` command will automatically run tests for each stage (_stage1_ to _stage4_).

Finally, confirm that the tests pass with Go solutions:

![image](https://github.com/gutyoh/Memorization-Tool-Go/assets/8846884/32ca8b7c-3478-4490-8eed-50b7c71756ab)
