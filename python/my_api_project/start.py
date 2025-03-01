# start.py
import uvicorn
from middleware.installer import run as middleware_run

def main():
    print("Hello, World!")
    middleware_run.run(uvicorn.run("my_api_project.main:app", host="127.0.0.1", port=8000, reload=True))
    # uvicorn.run("my_api_project.main:app", host="127.0.0.1", port=8000, reload=True)

if __name__ == "__main__":
    main()
