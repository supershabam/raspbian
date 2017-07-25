FROM ubuntu:16.04

RUN apt-get update \
 && apt-get install -y wpasupplicant \
 && rm -rf /var/lib/apt/lists/*
