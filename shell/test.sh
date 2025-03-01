#!/bin/bash

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <number>"
    exit 1
fi

urls=("http://localhost:8080/", "http://localhost:8080/assets/img/elon-musk.jpg", "http://localhost:8080/assets/img/logo.png")

# Extract the first command-line argument
number=$1

# Check if the argument is a valid integer
#if ! [[ "$number" =~ ^[0-9]+$ ]]; then
 #   echo "Error: '$number' is not a valid integer."
  #  exit 1
#fi

# Perform addition by 1
result=$((number + 1))

# Print the result
echo "The result of adding 1 to $number is: $result"
