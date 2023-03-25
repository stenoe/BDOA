# Application to check REST API server

from fastapi import FastAPI
from fastapi.responses import FileResponse
from fastapi.responses import StreamingResponse

import pandas as pd
import io


# define app
app = FastAPI()

# default endpoint, return message
@app.get("/")
async def read_root():
    return {"Hello" : "World"}

# Endpoint to download a csv
@app.get("/data/") 
async def read_file():
    file_path = "/code/data/co2_mm_mlo.csv"
    response = FileResponse(file_path, media_type="text/csv")
    response.headers["Content-Disposition"] = "attachment; filename=downlowded-file.csv"
    return response

# Edpoint for pandas streaming
@app.get("/pandas/")
async def streamFromDF():
    file_path = "/code/data/co2_mm_mlo.csv"
    df = pd.read_csv(file_path, skiprows=56)
    stream = io.StringIO()
    df.to_csv(stream, index=False)
    response = StreamingResponse(iter([stream.getvalue()]), media_type="text/csv")
    response.headers["Content-Disposition"] = "attachment; filename=exported_from_pandas.csv"
    return response

