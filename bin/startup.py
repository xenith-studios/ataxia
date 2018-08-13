#!/usr/bin/env python
import subprocess
# Check the system to make sure all required directories and files are in the
# right places
# If game data doesn't exist, print a message to run the game data seed script

# Check for exiting PID file to hint that a copy of the engine is already
# running or possibly crashed

proxy = subprocess.Popen(["bin/ataxia-proxy"])
subprocess.call(["bin/ataxia-engine"])
proxy.wait()
