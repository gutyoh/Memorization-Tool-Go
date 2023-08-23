#!/bin/bash

PROJECT_DIR="$HOME/Hyperskill/Memorization-Tool-Go"

source "$PROJECT_DIR/memo_tool_venv/bin/activate"

for stage in stage1 stage2 stage3 stage4; do
    echo "Running tests in $stage..."

    cd "$PROJECT_DIR/$stage" || exit

    python3 tests.py

    echo "--------------------------------------"
done
