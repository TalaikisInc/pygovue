#!/bin/bash

PROJECT=$1

cd /home
git clone https://github.com/xenu256/pygovue
mv /home/pygovue /home/$PROJECT
conda create --name $PROJECT python
conda activate $PROJECT
pip install -r requirements.txt
#part where db creation magic goes
python manage.py makemigrations
python manage.py migrate
python manage.py createsuperuser
cd /home/$PROJECT/api_server
go get github.com/lib/pq

cd /home/$PROJECT/frontend
npm install
