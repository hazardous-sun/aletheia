# Changelog

All notable changes to this project will be documented in this file.

## [unreleased]

### 🚀 Features

- Added more logs to Config
- Added an API Connector struct that will be used by the client to communicate with the server
- Now the GUI can communicate with the server
- Added a 'dev-setup.sh' script
- Added a specific color for info messages to the logger
- Added a QueryParser structure that will be used to build the queries for each news outlet
- Added some error messages for the QueryParser
- Added a config to remotely debug server-api
- Added configs for debugging server-api in VS Code
- Added an AI link extractor
- Added an API for the AI analyzer
- Added an API for the AI Analyzer
- Removed Ollama base image
- Ollama and LLM are now being pulled during build process
- Added a compose file to run the API and AI analyzer together on the same network
- The API url can be configured as the environment variable AI_ANALYZER_URL now
- Added run.sh for triggering server-api alongside ai-analyzer
- Removed unused dependabot workflow
- The root run.sh now tries to run podman-compose down before collecting the command-line options
- Ai-analyzer image now contains pciutils
- Updated the return package of 'getLinks' endpoint
- Debugging will only be available when running the project with the '--DEBUG' command-line option
- Crawlers now run concurrently

### 🐛 Bug Fixes

- The value for Config.Port was not being passed before checking if it was a valid integer, making the unit tests always fail
- Made pre-commit hook executable

### 💼 Other

- Updated CHANGELOG.md
- Updated CHANGELOG.md
- Fixed a typo in README.md
- Updated CHANGELOG.md
- Removed unused diagram
- Updated CHANGELOG.md
- Updated CHANGELOG.md
- Updated CHANGELOG.md
- Updated server-api go version required to 1.23
- Removed demo workflow
- Updated client go version required to 1.23
- Updated CHANGELOG.md
- Trying to identify why Crawl() is not returning the expected data
- Updated .gitignore
- Merge branch 'main' of github.com:hazardous-sun/aletheia
- Adjusting the web crawlers error handling
- Go mod download
- Go mod download
- Go mod tidy
- Updated CHANGELOG.md
- Added workflow to update dependencies
- Go get -u ./...
- Go mod tidy
- *(deps)* Bump golang.org/x/net from 0.36.0 to 0.38.0 in /client
- Updated CHANGELOG.md
- Updated CHANGELOG.md

### 🚜 Refactor

- Refactored the Config struct and added a Port field to it
- Removed duplicated alias
- Adjusted the color for info messages to cyan
- Refactored the run script
- Refactored the logic for building the GUI
- Refactored the unit tests to the fit the new model for the Config struct
- Adjusted the type used for logging the value of the check fields
- Adjusted the run script
- Now the client GUI fields are initialized with a mix of environment variables and flags
- Adjusted the commit template
- Renamed 'pre-commit' hook to 'commit-msg'
- Replaced the original drawio diagrams by excalidraw diagrams
- Renamed the 'generateNewsSearchURL' to 'Parse'
- Go Github Action now only tries to build the project with Go 1.23
- The 'description label is not being used in the 'go' workflow'
- Updated crawler use case
- Updated crawler repository
- Adjusted the ANSII code for the info color
- Adjusted content analyzer
- Updated requirements.txt
- Renamed 'fact-check-server' to 'aletheia-api' and 'news-db' to 'aletheia-db'
- Reorganized the root run and compose files
- Updated the base prompt to the ai-analyzer '/getLinks' endpoint
- Renamed ai-analyzer to aletheia-ai-analyzer
- Added custom logging to conn.go
- Updated logic for communicating with aletheia-ai-analyzer
- 'badCrawler' and 'collectCandidateBody' are now methods from CrawlerRepository
- Removed redundant code
- Removed unused import
- Removed redundant code

### 📚 Documentation

- Added more error messages
- Initialized git cliff
- Renamed the module to aletheia-client
- Renamed the module to aletheia-server
- Added an Excalidraw file with the models for the project
- Added a pre-commit hook to guarantee the commit message is using one of the parsers for git-cliff
- Added a template for the Git commit message
- Updated README.md 'Contributing' section
- Added an 'other' section to the commit-msg hook
- Improved server-api README.md
- Added a description to the 'go' workflow
- Added priorities to the 'TODO' section and fixed some typos
- Added a section in the README about the remote debugging
- Added a section in the README about the remote debugging
- Moved the debugging section to server-api README.md
- Updated 'go' workflow display name
- Updated commit-msg to only warn the user that the last commit message to be pushed should start with one of git-cliff keywords
- Updated the pre-push hook to enforce the use of git-cliff keywords on the last commit message to be pushed
- Updated README for server-api
- Updated README for ai-analyzer
- Updated root level README.md

### 🧪 Testing

- Adjusted the value expected for the info message color
- The go.yml workflow now also install some missing dependencies
- Updated unity tests for error logging
- Updated unity tests for crawler model

<!-- generated by git-cliff -->
