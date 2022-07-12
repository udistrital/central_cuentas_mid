FROM python:3
RUN pip install awscli
WORKDIR /app/src/
COPY entrypoint.sh main conf/app.conf ./
CMD ["./entrypoint.sh"]
