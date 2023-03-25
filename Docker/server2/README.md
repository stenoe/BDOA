# Simple example for a REST API using Python and FastAPI

In this example, the server has all together three endpoints:

/ gives a simple text message
/data/ returns a csv file without changes
/pandas/ returns a csv file that is streamed out of a pandas dataframe and has the info lines deleted
