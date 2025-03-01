from fastapi import FastAPI

from middleware.lib.mw_tracker import MwTracker

# from middleware import MwTracker
tracker=MwTracker()

app = FastAPI()

@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.get("/items/{item_id}")
def read_item(item_id: int, q: str = None):
    return {"item_id": item_id, "q": q}
