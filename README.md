# Aletheia

> "If a lie is only printed often enough, it becomes a quasi-truth, and if such a truth is repeated often enough, it
> becomes an article of belief, a dogma, and men will die for it."
> - Isabella Jane Blagden

This project implements a "fact-checking" tool (note the quotes) developed in Go. It aims to help users avoid falling 
into the rabbit holes of fake news by providing a way to verify if an online post aligns with information from reputable 
news sources. The tool works by crawling predefined online news outlets, gathering the latest news on a specific topic, 
and analyzing the content using AI to compare it with the original post. The flow of this process is outlined in 
[`diagrams/`](/diagrams).

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

Feel free to contribute to the project, just make sure to run `git config core.hooksPath .githooks` first to activate 
the project hooks.