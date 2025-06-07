import ollama
from bs4 import BeautifulSoup
import json
import re
from typing import List, Dict, Union


class LinkExtractor:
    def __init__(self, model_name: str = 'deepseek-r1:1.5b'):
        self.model_name = model_name
        try:
            ollama.pull(self.model_name)
        except Exception as e:
            print(f"Warning: Could not pull model - {str(e)}")

    def extract_links(self, html_content: str) -> List[Dict[str, str]]:
        try:
            # Clean and limit the HTML content
            clean_html = self._clean_html(html_content)[:90000]  # Limit to 90k chars

            # Create a more reliable prompt template
            prompt_template = """
            Analyze the following HTML content and extract all news article links with their titles.
            Return ONLY a valid JSON array where each element is an object with "title" and "url" properties.
            
            HTML Content:
            {html}
            
            Required Output Format:
            [
                {"title": "TITLE_ARTICLE_1", "url": "URL_ARTICLE_1"},
                {"title": "TITLE_ARTICLE_1", "url": "URL_ARTICLE_2"},
                ...
                {"title": "TITLE_ARTICLE_N", "url": "URL_ARTICLE_N"},}
            ]
            
            Rules:
            1. Only include links that point to news articles
            2. Titles should be 3-15 words, in the original language
            3. URLs must be complete and valid
            4. If no news links found, return empty array []
            5. No additional text or explanations
            """

            # Safely format the prompt
            try:
                prompt = prompt_template.format(html=clean_html)
            except Exception as e:
                print(f"Error formatting prompt: {str(e)}")
                prompt = prompt_template.replace("{html}", clean_html)

            response = ollama.generate(
                model=self.model_name,
                prompt=prompt,
                options={
                    'temperature': 0.1,
                    'format': 'json'
                }
            )

            # Clean and validate the response
            response = self._collect_json_section(response['response'])

            json_str = self._extract_json(response['response'])
            links = json.loads(json_str)

            # Validate the links structure
            if not isinstance(links, list):
                raise ValueError("Response is not a list")

            validated_links = []
            for link in links:
                if isinstance(link, dict) and 'url' in link and 'title' in link:
                    validated_links.append({
                        'title': str(link['title']),
                        'url': str(link['url'])
                    })

            return validated_links

        except json.JSONDecodeError as e:
            print(f"JSON decode error. Response was: {response['response']}")
            return []
        except Exception as e:
            print(f"Link extraction error: {str(e)}")
            raise ValueError(f"Failed to extract links: {str(e)}")

    def _clean_html(self, html: str) -> str:
        """Clean HTML content before processing"""
        try:
            soup = BeautifulSoup(html, 'html.parser')
            # Remove unwanted elements
            for element in soup(['script', 'style', 'meta', 'link', 'noscript']):
                element.decompose()
            # Get clean text
            return soup.get_text(' ', strip=True)
        except Exception as e:
            print(f"HTML cleaning error: {str(e)}")
            return html[:20000]  # Fallback to simple truncation

    def _extract_json(self, text: str) -> str:
        """Extract JSON string from response text"""
        # Look for JSON array pattern
        start = text.find('[')
        end = text.rfind(']') + 1
        if start >= 0 and end > 0:
            return text[start:end]
        return '[]'  # Return empty array if no JSON found

    def _collect_json_section(self, result: str) -> str:
        start = result.find('[')
        end = result.rfind(']')
        return result[start:end]


if __name__ == "__main__":
    quit(1)