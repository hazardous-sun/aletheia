import ollama
from typing import Dict, Optional

class ContentAnalyzer:
    def __init__(self, model_name: str = 'deepseek-r1:1.5b'):
        """
        Initialize the ContentAnalyzer with a specific AI model.

        Args:
            model_name: Name of the Ollama model to use (default: 'deepseek-r1:1.5b')
        """
        self.model_name = model_name
        # Ensure the model is available
        ollama.pull(self.model_name)

    def analyze_content(
            self,
            original_post_content: str,
            online_news_content: str,
            user_context: Optional[str] = None
    ) -> str:
        """
        Compare original post content against news sources and provide analysis.

        Args:
            original_post_content: Content of the original post to analyze
            online_news_content: Content from reputable news sources for comparison
            user_context: Optional additional context for the analysis

        Returns:
            The AI-generated analysis as a string

        Raises:
            ValueError: If required content is missing
        """
        # Validate inputs
        if not original_post_content.strip():
            raise ValueError("Original post content cannot be empty")
        if not online_news_content.strip():
            raise ValueError("Online news content cannot be empty")

        # Build the prompt
        prompt = self._build_prompt(
            original_post_content,
            online_news_content,
            user_context
        )

        # Get and return AI response
        return self._get_ai_response(prompt)

    def _build_prompt(
            self,
            original_post: str,
            news_content: str,
            user_context: Optional[str] = None
    ) -> str:
        """Construct the analysis prompt for the AI model"""
        prompt_template = """
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

        Original post content: "{original_post}"

        Reputable news sources content: "{news_content}"
        """

        prompt_text = prompt_template.format(
            original_post=original_post,
            news_content=news_content
        )

        if user_context and user_context.strip():
            prompt_text += f"\n\nExtra user context: {user_context}"

        return prompt_text

    def _get_ai_response(self, prompt: str) -> str:
        """Get response from the AI model"""
        response = ollama.generate(
            model=self.model_name,
            prompt=prompt,
            options={
                'temperature': 0.3,  # More deterministic output
                'num_ctx': 8192      # Larger context window if needed
            }
        )
        return response['response']

# Example usage
if __name__ == "__main__":
    # Example data (in a real API, this would come from the request)
    sample_post = "The government announced new tax cuts that will benefit middle-class families."
    sample_news = """
    [Reuters] The parliament passed new tax legislation today that will reduce rates for incomes under $100,000.
    [BBC] New tax bill approved, affecting approximately 60% of taxpayers with cuts averaging $1,200 annually.
    """
    sample_context = "Focus specifically on the income brackets mentioned in the post versus the news."

    # Create analyzer and get results
    analyzer = ContentAnalyzer()
    try:
        analysis = analyzer.analyze_content(
            original_post_content=sample_post,
            online_news_content=sample_news,
            user_context=sample_context
        )
        print("Analysis Results:")
        print(analysis)
    except ValueError as e:
        print(f"Error: {str(e)}")