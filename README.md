# Aletheia

> "If a lie is only printed often enough, it becomes a quasi-truth, and if such a truth is repeated often enough, it
> becomes an article of belief, a dogma, and men will die for it."
> - Isabella Jane Blagden

This project implements a "fact checking" (**please** note the quotes here) tool developed in Go. It was designed as an
**attempt** to help users avoid falling into rabbit holes of fake news. The idea is to provide a way for the user to
check if a post from the internet contains information that aligns with what the rest of the world is currently saying.
The idea is to implement a *generic* crawling tool that will navigate through a number of predefined online news outlets
collecting the most recent news they published about a specific topic. Then, all the content of the news outlets passes
through an AI analyzer that compares if the content of the original post aligns with what the rest of the world says.
This idea of the flow can be briefly analyzed in [`diagrams/`](/diagrams). Following the idea presented by the quote 
from Isabella Blagden, it is undeniable that, in the scenario where **everyone** in the world lies about a topic, this
tool would be unable to fact check anything, so keep in mind that this is not a perfect oracle that never fails.

## Features

### [Client](client/README.md)

The client is a simple GUI developed in Go using the Fyne framework where the user passes the URL to the news they would
like to fact-check. There, they could also inform some other inputs, such as extra context and flags informing if the 
application should also fact-check an image, or video.

### [Server API](server-api/README.md)

The Server API was also developed in Go, using the Gin framework. It continuously listens for requests and, when it 
receives one, it collects the data from the package, that should contain a URL and possibly some context. With the 
received data, it starts a number of web crawlers and searches for related data on predefined news outlets.

### [AI Analyzer](ai-analyzer/README.md)

The AI Analyzer compares the data collected by the crawlers and compares it to the original post sent by the client. It
then tries to determine if the content from the post aligns with what the rest of the world is saying.
