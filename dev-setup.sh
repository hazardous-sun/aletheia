#!/usr/bin/env bash

# This script is a setup that anyone willing to contribute to the project should run
# before making any changes to the repo. It sets some Git configs that are required
# in order to guarantee the quality of the project.

# Setup project's Git Hooks
git config core.hooksPath ./docs/.githooks

# Setup Git commit message template
git config commit.template ./docs/.gitmessage