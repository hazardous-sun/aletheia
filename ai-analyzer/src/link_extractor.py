import ollama
from bs4 import BeautifulSoup
import json
import re
from typing import List, Dict


class LinkExtractor:
    def __init__(self, model_name: str = 'phi3:3.8b'):
        self.model_name = model_name
        try:
            ollama.pull(self.model_name)
        except Exception as e:
            print(f"Warning: Could not pull model - {str(e)}")

    def extract_links(self, html_content: str) -> List[Dict[str, str]]:
        try:
            # Log input size for debugging
            print(f"Original HTML size: {len(html_content)}")

            # Clean HTML while preserving structure
            clean_html = self._clean_html(html_content)
            print(f"Cleaned HTML size: {len(clean_html)}")

            # Reduce size to fit model context
            # clean_html = clean_html[:12000]
            print(f"Truncated HTML size: {len"(clean_html)}")

            # Define prompt template with escaped curly braces
            prompt_template = """
        Analyze the following HTML content and extract all news article links with their titles.
        Return ONLY a valid JSON array where each element is an object with "title" and "url" properties.
        
        HTML Content:
        {html}
        
        Required Output Format:
        [
            {{"title": "TITLE_ARTICLE_1", "url": "URL_ARTICLE_1"}},
            {{"title": "TITLE_ARTICLE_2", "url": "URL_ARTICLE_2"}},
            {{"title": "TITLE_ARTICLE_3", "url": "URL_ARTICLE_3"}}
        ]
        
        Rules:
        1. Only include links that point to news articles
        2. Titles should be 3-15 words, in the original language
        3. URLs must be complete and valid (include http/https)
        4. If no news links found, return empty array []
        5. No additional text or explanations
        """

            # Safely format the prompt
            try:
                prompt = prompt_template.format(html=clean_html)
            except Exception as e:
                print(f"Error formatting prompt: {str(e)}")
                # Manual replacement fallback
                prompt = prompt_template.replace("{html}", clean_html)
                print(f"Used manual replacement for prompt formatting")

            # Generate response from the model
            response = ollama.generate(
                model=self.model_name,
                prompt=prompt,
                options={'temperature': 0.1, 'format': 'json'}
            )

            raw_response = response['response']
            print(f"Model raw response: {raw_response[:500]}...")  # Log first 500 chars

            # Extract JSON section
            json_str = self._extract_json(raw_response)
            print(f"Extracted JSON: {json_str[:500]}...")  # Log first 500 chars

            # Parse JSON
            try:
                links = json.loads(json_str)
            except json.JSONDecodeError:
                # Attempt to fix common JSON issues
                fixed_json = self._fix_json(json_str)
                links = json.loads(fixed_json)

            # Validate the links structure
            if not isinstance(links, list):
                print(f"Response is not a list: {type(links)}")
                return []

            validated_links = []
            for link in links:
                if isinstance(link, dict) and 'url' in link and 'title' in link:
                    # Ensure URL has valid protocol
                    url = str(link['url']).strip()
                    if url.startswith(('http://', 'https://')):
                        validated_links.append({
                            'title': str(link['title']).strip(),
                            'url': url
                        })
                    else:
                        print(f"Invalid URL protocol: {url}")
                else:
                    print(f"Skipping invalid link structure: {link}")

            print(f"Extracted {len(validated_links)} valid links")
            return validated_links

        except json.JSONDecodeError as e:
            print(f"JSON decode error: {str(e)}")
            return []
        except Exception as e:
            print(f"Unexpected error in extract_links: {type(e).__name__}: {str(e)}")
            return []

    def _fix_json(self, text: str) -> str:
        """Attempt to fix common JSON formatting issues"""
        # Fix 1: Remove trailing commas
        text = re.sub(r',\s*([}\]])', r'\1', text)

        # Fix 2: Add quotes around unquoted property names
        text = re.sub(r'([\{,])\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*:', r'\1"\2":', text)

        # Fix 3: Escape problematic characters
        text = text.replace('\n', '\\n').replace('\t', '\\t')

        # Fix 4: Handle single quotes
        text = re.sub(r"'([^']*)'", r'"\1"', text)

        # Fix 5: Remove control characters
        text = re.sub(r'[\x00-\x1F\x7F]', '', text)

        # Fix 6: Handle missing commas between objects
        text = re.sub(r'}\s*{', '},{', text)

        print(f"Fixed JSON: {text[:500]}...")
        return text

    def _clean_html(self, html: str) -> str:
        """Clean HTML while preserving structure"""
        try:
            soup = BeautifulSoup(html, 'html.parser')
            for element in soup(['script', 'style', 'meta', 'link', 'noscript']):
                element.decompose()
            return str(soup)
        except Exception as e:
            print(f"HTML cleaning error: {str(e)}")
            return html[:20000]

    def _extract_json(self, text: str) -> str:
        """Robust JSON extraction with bracket counting"""
        start = text.find('[')
        if start < 0:
            return '[]'

        # Count brackets to find proper closing
        open_brackets = 0
        for i in range(start, len(text)):
            if text[i] == '[':
                open_brackets += 1
            elif text[i] == ']':
                open_brackets -= 1
                if open_brackets == 0:
                    return text[start:i+1]
        return '[]'


if __name__ == "__main__":
    quit(1)