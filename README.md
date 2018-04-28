# Hot Cold Bot

A bot that learns how to reach a goal by learning from its past mistakes.

# Getting It Running

This was built using Golang 1.9.5. Simple clone the repo and run <code>go build</code> within the directory.

## Windows Users

If running on windows, you must change the <code>"clear"</code> command in main.go to <code>"cls"</code> before building.

# How It Works

Imagine a character on a 2D grid. It has a pre-defined set of directions it must move in, and gets feedback each time it makes a move. The feedback is simply if the direction the character moved in made it further from the goal, or not. If it moved further (or out of bounds), it will make note of that and try a direction it has not yet moved in yet. The character cannot see where the goal is, and can only find out where it is by adjusting its movements according to the feedback it gets.

The character must follow the following rules:
- The pre-defined steps must not be changed in the middle of a round
- The character may not go out of bounds
- The character may not move diagonally
