import ollama
import os

# Define the environment variable names
ORIGINAL_POST_CONTENT = "ORIGINAL_POST_CONTENT"
ONLINE_NEWS_CONTENT = "ONLINE_NEWS_CONTENT"
USER_CONTEXT = "USER_CONTEXT"

class InvalidPrompt(Exception):
    def __init__(self, message, code):
        self.message = message
        self.code = code
        super().__init__(f"{message} (Code: {code})")


def main():
    # Pull the DeepSeek 1.5B model (if not already pulled)
    ollama.pull('deepseek-r1:1.5b')

    # Collect environment variables
    variables: dict[str, str] = collect_env_variables()

    # Build the prompt
    prompt_context: str = build_prompt(variables)

    # Prompt the AI model
    prompt(prompt_context)


def collect_env_variables() -> dict[str, str]:
    variables: dict[str, str] = {
        ORIGINAL_POST_CONTENT: os.environ.get(ORIGINAL_POST_CONTENT, ""),
        ONLINE_NEWS_CONTENT: os.environ.get(ONLINE_NEWS_CONTENT, ""),
        USER_CONTEXT: os.environ.get(USER_CONTEXT, "")
    }
    return variables


def build_prompt(
        variables: dict[str, str]
) -> str:
    original_post_content = get_content(variables, ORIGINAL_POST_CONTENT)
    reputable_news_content = get_content(variables, ONLINE_NEWS_CONTENT)
    user_context = get_content(variables, USER_CONTEXT)
    context: str = f"""
    You are an AI analyzer tasked with comparing the content of an original post submitted by a user against data
    gathered from reputable news sources. Your goal is to assess whether the post's content aligns with or contradicts
    the information from these sources. Your analysis must be honest, accurate, and strictly based on the data provided
    to you. Follow these guidelines:
      1. Honesty and Accuracy:
        - Only use the data provided by the crawlers from reputable news sources. Do not create, infer, or assume any
        information that is not explicitly present in the data.
        - If the data does not support a conclusion, clearly state that there is insufficient information to verify the
        post.
      2. Relevance Check:
        - Before comparing the post to the news data, analyze whether the news articles are relevant to the topic of the
        original post. If the news data does not relate to the post's topic, clearly state that no relevant information
        was found.
        - Use semantic analysis to determine if the news articles discuss the same subject, event, or claim as the
        original post.
      3. Alignment Analysis:
        - If the news data is relevant, compare the claims, facts, and context of the original post to the information
        in the news articles.
        - Identify whether the post aligns with, contradicts, or partially matches the news data.
        - Highlight specific points of agreement or disagreement, and provide evidence

    Original post content: "{original_post_content}"

    Reputable news sources content: "{reputable_news_content}"
    """

    if user_context != "":
        context = f"{context} \n\n Extra context: {user_context}"

    return context


def get_content(variables: dict[str, str], section: str) -> str:
    temp = variables.get(section, "")

    if section != USER_CONTEXT and temp == "":
        raise InvalidPrompt(f"{section} cannot be empty", 1)

    return temp


def prompt(prompt_context: str):
    # Generate a response using the model
    response = ollama.generate(
        model='deepseek-r1:1.5b',
        prompt=prompt_context
    )

    # Print the response
    print(response['response'])


if __name__ == "__main__":
    main()