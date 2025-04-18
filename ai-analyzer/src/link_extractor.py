import ollama
from bs4 import BeautifulSoup
import json
from typing import List, Dict, Union

class LinkExtractor:
    def __init__(self, model_name: str = 'deepseek-r1:1.5b'):
        """
        Initialize the LinkExtractor with a specific AI model.

        Args:
            model_name: Name of the Ollama model to use (default: 'deepseek-r1:1.5b')
        """
        self.model_name = model_name
        # Ensure the model is available
        ollama.pull(self.model_name)

    def extract_links(self, html_content: str) -> List[Dict[str, str]]:
        """
        Extract relevant article links from HTML content using AI.

        Args:
            html_content: Raw HTML content to analyze

        Returns:
            List of dictionaries containing title-url pairs
            Example: [{"title1": "url1"}, {"title2": "url2"}]
        """
        if not html_content or not html_content.strip():
            raise ValueError("HTML content cannot be empty")

        # Build the prompt
        prompt = self._build_prompt(html_content)

        # Get AI response
        response = self._get_ai_response(prompt)

        # Parse and return the links
        return self._parse_response(response)

    def _build_prompt(self, html_content: str) -> str:
        """Construct the prompt for the AI model"""
        return f"""
        Extract all news article links and their titles from the provided HTML content.
        Return ONLY a JSON array where each element is a dictionary with a single key-value pair:
        the article title as key and the URL as value.
        
        Requirements:
        1. Include only links to actual news articles (no navigation, ads, or unrelated links)
        2. Use the actual article title text (cleaned, 3-15 words)
        3. Return complete, absolute URLs
        4. Return ONLY the JSON, no additional text or explanation
        
        Example output:
        [
            {{"Breaking News: Earthquake Hits Region": "https://example.com/earthquake"}},
            {{"Political Summit Concludes": "https://example.com/summit"}}
        ]
        
        HTML content:
        \"\"\"
        {html_content}
        \"\"\"
        """

    def _get_ai_response(self, prompt: str) -> str:
        """Get response from the AI model"""
        response = ollama.generate(
            model=self.model_name,
            prompt=prompt,
            options={
                'temperature': 0.3,  # More deterministic output
                'num_ctx': 8192     # Larger context window if needed
            }
        )
        return response['response']

    def _parse_response(self, response: str) -> List[Dict[str, str]]:
        """
        Parse the AI response and extract the JSON data.

        Args:
            response: Raw response from the AI model

        Returns:
            Parsed list of title-url dictionaries

        Raises:
            ValueError: If the response cannot be parsed as valid JSON
        """
        try:
            # Try to find JSON in the response
            start = response.find('[')
            end = response.rfind(']') + 1
            json_str = response[start:end]

            # Parse and validate
            links = json.loads(json_str)

            # Validate structure
            if not isinstance(links, list):
                raise ValueError("Expected a JSON array")

            for item in links:
                if not isinstance(item, dict) or len(item) != 1:
                    raise ValueError("Each array item should be a single key-value pair")

            return links

        except (json.JSONDecodeError, ValueError, AttributeError) as e:
            raise ValueError(f"Failed to parse AI response: {str(e)}. Original response: {response}")

# Example usage
if __name__ == "__main__":
    quit(1)