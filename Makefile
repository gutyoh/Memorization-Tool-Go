SHELL := /bin/bash
STAGES := stage1 stage2 stage3 stage4
ROOT_DIR := $(shell pwd)

# Detect OS
ifeq ($(OS),Windows_NT)
    DETECTED_OS := Windows
    PYTHON := python
else
    DETECTED_OS := $(shell uname -s)
    PYTHON := python3
endif

VENV := $(ROOT_DIR)/memo_tool_venv/bin/$(PYTHON)

all: venv $(STAGES) finished

$(STAGES): venv
	@echo -e "\033[34m========================================================\033[0m"
	@echo -e "\033[34mRunning tests for $@\033[0m"
	@echo -e "\033[34m========================================================\033[0m"
	@OUTPUT=$$(cd $@ && $(VENV) tests.py); \
	echo "$$OUTPUT"; \
	if echo "$$OUTPUT" | grep -q "#educational_plugin FAILED"; then \
		echo -e "\033[31mTests for $@ FAILED! ❌\033[0m"; \
		exit 1; \
	else \
		echo -e "\033[32mTests for $@ PASSED! ✅\033[0m"; \
	fi
	@echo

finished:
	@echo -e "\033[34m========================================================\033[0m"
	@echo -e "\033[32mAll tests finished and passed SUCCESSFULLY! ✅\033[0m"
	@echo -e "\033[34m========================================================\033[0m"

venv: memo_tool_venv/bin/activate

memo_tool_venv/bin/activate:
	$(PYTHON) -m venv memo_tool_venv
	$(VENV) -m pip install -r requirements.txt

clean:
	rm -rf memo_tool_venv
