#!/bin/bash

if [ ! -f .env ]; then
    cp .env.example .env
    echo "Created .env file from template. Please update with your values."
else
    echo ".env file already exists."
fi
