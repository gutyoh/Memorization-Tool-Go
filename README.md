# Memorization-Tool-Go
[Hyperskill Memorization Tool project](https://hyperskill.org/projects/159) from the Python/Flask track with Go solutions using GORM

# How to run this project with Go solutions

Since this project uses GORM you will need to initialize Go modules. I recommend creating a folder with the same name "Memorization-Tool-Go" and then cloning this repository within that folder.

After cloning the repository, the directory tree should look like this:

```
Memorization-Tool-Go # "top-level" directory
|   course-info.yaml
|   course-remote-info.yaml
|   requirements.txt
|
\---Memorization-Tool-Go
    |   lesson-info.yaml
    |   lesson-remote-info.yaml
    |
    +---stage1
    |   |   main.go
    |   |   task-info.yaml
    |   |   task-remote-info.yaml
    |   |   task.html
    |   |   tests.py
    |   |
    |   \---test
    |       |   tests.py
    |       |   __init__.py
    |       |
    |       \---__pycache__
    |               tests.cpython-310.pyc
    |               __init__.cpython-310.pyc
    |
    +---stage2
    |   |   main.go
    |   |   task-info.yaml
    |   |   task-remote-info.yaml
    |   |   task.html
    |   |   tests.py
    |   |
    |   \---test
    |       |   tests.py
    |       |   __init__.py
    |       |
    |       \---__pycache__
    |               tests.cpython-310.pyc
    |               __init__.cpython-310.pyc
    |
    +---stage3
    |   |   main.go
    |   |   task-info.yaml
    |   |   task-remote-info.yaml
    |   |   task.html
    |   |   tests.py
    |   |
    |   \---test
    |       |   tests.py
    |       |   __init__.py
    |       |
    |       \---__pycache__
    |               tests.cpython-310.pyc
    |               __init__.cpython-310.pyc
    |
    \---stage4
        |   main.go
        |   task-info.yaml
        |   task-remote-info.yaml
        |   task.html
        |   tests.py
        |
        \---test
            |   tests.py
            |   __init__.py
            |
            \---__pycache__
                    tests.cpython-310.pyc
                    __init__.cpython-310.pyc
```

Then you'll have to `cd` into the "top-level" _Memorization-Tool-Go_ directory and initialize Go modules:

```
go mod init Memorization-Tool-Go
go mod tidy
```

After initializing Go modules, below is a picture of how the project directory should look like including the _go.mod_ and _go.sum_ files:

![image](https://user-images.githubusercontent.com/8846884/215644131-411d7a10-78b6-4ef3-962a-24bb6ba9ef97.png)

The next step is to make sure your Python interpreter has the latest [hs-test-python](https://github.com/hyperskill/hs-test-python/releases/tag/v10) release installed.

Finally, you should be able to `cd` into the _stage1-stage4_ folders and within them run `python tests.py` and confirm that the test pass with Go solutions.
