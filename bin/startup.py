#!/usr/bin/env python
import subprocess
# Check the system to make sure all required directories and files are in the
# right places
# If game data doesn't exist, print a message to run the game data seed script

# Check for exiting PID files to hint that a copy of the programs are already
# running or possibly crashed

proxy = subprocess.Popen(["./bin/ataxia-proxy", "-d"])
engine = subprocess.Popen(["./bin/ataxia-engine", "-d"])
engine.wait()
proxy.wait()
