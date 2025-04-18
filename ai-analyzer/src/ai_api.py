from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from content_analyzer import ContentAnalyzer
from link_extractor import LinkExtractor
import uvicorn

app = FastAPI()

# Add CORS middleware to allow requests from your frontend
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Request models for type validation
class LinkRequest(BaseModel):
    html_content: str

class AnalysisRequest(BaseModel):
    post_content: str
    news_content: str
    user_context: str = None

@app.post("/getLinks")
async def get_links(request: LinkRequest):
    """
    Endpoint to extract article links from HTML content
    """
    extractor = LinkExtractor()
    try:
        links = extractor.extract_links(request.html_content)
        return {
            "success": True,
            "links": links,
            "count": len(links)
        }
    except Exception as e:
        raise HTTPException(
            status_code=400,
            detail=str(e)
        )

@app.post("/analyze")
async def analyze_content(request: AnalysisRequest):
    """
    Endpoint to analyze post content against news sources
    """
    analyzer = ContentAnalyzer()
    try:
        analysis = analyzer.analyze_content(
            original_post_content=request.post_content,
            online_news_content=request.news_content,
            user_context=request.user_context
        )
        return {
            "success": True,
            "analysis": analysis
        }
    except ValueError as e:
        raise HTTPException(
            status_code=400,
            detail=str(e)
        )
    except Exception as e:
        raise HTTPException(
            status_code=500,
            detail=f"An unexpected error occurred: {str(e)}"
        )

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=7654)