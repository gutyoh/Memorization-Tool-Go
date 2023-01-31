# Memorization-Tool-Go
HS Memorization Tool project with Go solution using GORM

# How to run this project with Go solutions

Since this project uses GORM you will need to initialize Go modules. I recommend creating a folder with the same name "Memorization-Tool-Go" and then cloning this repository within that folder.

After cloning the repository, the directory tree should look like this:

```
Memorization-Tool-Go
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

Then you'll have to `cd` into the "top" _Memorization-Tool-Go_ directory and initialize Go modules:

```
go mod init Memorization-Tool-Go
go mod tidy
```

After that make sure your Python interpreter has the latest [hs-test-python](https://github.com/hyperskill/hs-test-python/releases/tag/v10) release installed.

Finally, you should be able to `cd` into the _stage1-stage4_ folders and within them run `python tests.py` and confirm that the test pass with Go solutions.
