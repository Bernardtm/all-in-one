from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

# Allow all origins (not recommended for production)
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],       # Allows all origins
    allow_credentials=True,    # Allows cookies and other credentials
    allow_methods=["*"],       # Allows all methods
    allow_headers=["*"],       # Allows all headers
)

@app.get("/")
async def root():
    return {"message": "UP"}