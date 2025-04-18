from content_analyzer import ContentAnalyzer
from link_extractor import LinkExtractor

def handle_article_link_request(html_content: str):
    extractor = LinkExtractor()
    try:
        links = extractor.extract_links(html_content)
        return {
            "success": True,
            "links": links,
            "count": len(links)
        }
    except Exception as e:
        return {
            "success": False,
            "error": str(e)
        }


def handle_content_analysis_request(post_content: str, news_content: str, user_context: str = None):
    analyzer = ContentAnalyzer()
    try:
        analysis = analyzer.analyze_content(
            original_post_content=post_content,
            online_news_content=news_content,
            user_context=user_context
        )
        return {
            "success": True,
            "analysis": analysis
        }
    except ValueError as e:
        return {
            "success": False,
            "error": str(e)
        }
