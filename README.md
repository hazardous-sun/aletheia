# Aletheia

> "If a lie is only printed often enough, it becomes a quasi-truth, and if such a truth is repeated often enough, it
> becomes an article of belief, a dogma, and men will die for it."
> - Isabella Jane Blagden

This project implements a "fact-checking" tool (note the quotes) developed in Go. It aims to help users avoid falling
into the rabbit holes of fake news by providing a way to verify if an online post aligns with information from reputable
news sources. The tool works by crawling predefined online news outlets, gathering the latest news on a specific topic,
and analyzing the content using AI to compare it with the original post. The flow of this process is outlined in
[`diagrams/`](/docs/diagrams).

However, as Isabella Blagden’s quote suggests, if everyone in the world were to lie about a topic (or if no one even
talked about it), this tool would be unable to fact-check effectively. Thus, it is not an infallible oracle, but rather
a practical aid in navigating information.

## Features

### [Client](client/README.md)

The client is a straightforward GUI built in Go using the Fyne framework. Users can input the URL of the news they wish
to fact-check, along with additional context or flags to specify whether the tool should analyze images or videos.

### [Server API](server-api/README.md)

The Server API, developed in Go using the Gin framework, continuously listens for incoming requests. Upon receiving a
request, it processes the provided data and initiates multiple web crawlers to search for related information across
predefined news outlets.

### [AI Analyzer](ai-analyzer/README.md)

The AI Analyzer compares the data gathered by the crawlers against the original post submitted by the client, assessing
whether the post’s content aligns with information from other sources.

## Contributing

Feel free to contribute to the project, just make sure to run the [dev-setup.sh](dev-setup.sh) script first, to set up
the correct Git configs for the project. Currently, the setup script activates the project hooks, maintained under the
[`.githooks`](docs/.githooks) directory, and sets [`.gitmessage`](docs/.gitmessage) as the commit template message.
Also, be sure to update [`CHANGELOG.md`](CHANGELOG.md). The project is currently using
[git-cliff](https://git-cliff.org/) to automatically generate the changelog, you can install it by following the steps
under the official website. After installing git-cliff, you can easily update the changelog by running the following
command:

```shell
git-cliff -o CHANGELOG.md
```

### Git Hooks

Currently, the project has 2 Git Hooks:

- [`commit-msg`](docs/.githooks/commit-msg): this hook guarantees that the commit messages start with one of the 
  keywords used by git-cliff to generate the `CHANGELOG.md` file.
- [`pre-push`](docs/.githooks/pre-push): this hook runs the project's unit tests before pushing the local changes to the
  remote repository.