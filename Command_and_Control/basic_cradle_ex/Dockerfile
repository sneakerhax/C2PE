FROM python:3


ADD server.py /server.py
RUN apt install && apt upgrade -y
RUN python3 -m pip install flask gunicorn

# Dev
# ENTRYPOINT [ "flask", "--app", "server", "run", "--host=0.0.0.0", "-p", "80" ]

# Prod
ENTRYPOINT [ "gunicorn", "--bind", "0.0.0.0:8080", "server:app" ]