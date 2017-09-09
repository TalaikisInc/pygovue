#!/bin/bash

PROJECT=$1
PORT=$2

cd /home
git clone https://github.com/xenu256/pygovue
mv /home/pygovue /home/$PROJECT
conda create --name $PROJECT python
conda activate $PROJECT
pip install -r requirements.txt
python manage.py makemigrations
python manage.py migrate
python manage.py createsuperuser

cd /home/$PROJECT/api_server
go get github.com/lib/pq
go build

npm install pm2 -g
cd /home/$PROJECT/frontend
npm install
PORT=$PORT nuxt build
PORT=$PORT pm2 --name $PROJECT start npm -- start
# here goes nginx conf creation magic (TODO)