# Memorization-Tool-Go
[Hyperskill Memorization Tool project](https://hyperskill.org/projects/159) from the Python/Flask track with Go solutions using GORM

## Requirements:

- Go version 1.21.0
- Python 3 (for testing)

## How to run this project with Go solutions

1. **Clone the Repository and `cd` into it:**

```shell
git clone https://github.com/gutyoh/Memorization-Tool-Go.git
cd Memorization-Tool-Go
```

2. **Directory Structure**: After cloning the repository, the directory tree should look like this:

```
Memorization-Tool-Go
├── go.mod
├── go.sum
├── lesson-info.yaml
├── lesson-remote-info.yaml
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

**Linux and macOS:**

```shell
python3 -m venv memo_tool_venv
source memo_tool_venv/bin/activate
```

**Windows:**

```bash
python -m venv memo_tool_venv
memo_tool_venv\Scripts\activate
```


4. **Install the necessary Python packages:**

```shell
pip install -r requirements.txt
```


Then you'll have to `cd` into the "top-level" _Memorization-Tool-Go_ directory and initialize Go modules:

```
go mod init Memorization-Tool-Go
go mod tidy
```

After completing the above steps, below is a picture of how the project directory should look like:

![image](https://github.com/gutyoh/Memorization-Tool-Go/assets/8846884/82e6afa6-d251-4dd0-91f4-e637dd25a390)

5. **Run the Python tests:**

Finally, you should be able to `cd` into the _stage1-stage4_ folders and within them, run:

**Linux and macOS:**

```shell
python3 tests.py
```

**Windows:**

```bash
python tests.py
```

And confirm that the tests pass with Go solutions:

![image](https://github.com/gutyoh/Memorization-Tool-Go/assets/8846884/e0014c2b-3140-4875-aba4-915ad026427a)
