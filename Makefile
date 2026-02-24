# Makefile for Building and Running Application on Linux

# Compiler and Linker
CC = gcc

# Directories
SRC_DIR = src
BIN_DIR = bin
OBJ_DIR = obj

# Source Files
SOURCES = $(wildcard $(SRC_DIR)/*.c)
OBJ_FILES = $(SOURCES:$(SRC_DIR)/%.c=$(OBJ_DIR)/%.o)

# Executable Name
EXECUTABLE = $(BIN_DIR)/application

# Flags
CFLAGS = -Wall -Werror

# Build Rules
all: $(EXECUTABLE)

$(EXECUTABLE): $(OBJ_FILES)
	$(CC) $(OBJ_FILES) -o $@

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c | $(OBJ_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(OBJ_DIR):
	mkdir -p $@

.PHONY: clean
clean:
	 rm -rf $(OBJ_DIR)/*.o
	 rm -f $(EXECUTABLE)
