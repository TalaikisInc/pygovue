from os.path import dirname, abspath, join
from os import environ

from psycopg2 import extensions
from dotenv import load_dotenv


BASE_DIR = dirname(dirname(abspath(__file__)))

load_dotenv(join(BASE_DIR, '.env'))

SECRET_KEY = environ.get("SECRET_KEY")
assert isinstance(SECRET_KEY, str), "Check environments."

DEBUG = int(environ.get("DEV"))

ALLOWED_HOSTS = []

INSTALLED_APPS = [
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'tasks',
    'tasks.summarizer'
]

MIDDLEWARE = [
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
    'django.middleware.common.CommonMiddleware',
    'django.middleware.csrf.CsrfViewMiddleware',
    'django.contrib.auth.middleware.AuthenticationMiddleware',
    'django.contrib.messages.middleware.MessageMiddleware',
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
]

ROOT_URLCONF = 'backend.urls'

TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [],
        'APP_DIRS': True,
        'OPTIONS': {
            'context_processors': [
                'django.template.context_processors.debug',
                'django.template.context_processors.request',
                'django.contrib.auth.context_processors.auth',
                'django.contrib.messages.context_processors.messages',
            ],
        },
    },
]

WSGI_APPLICATION = 'backend.wsgi.application'

DATABASE_USER = environ.get("DATABASE_USER")
DATABASE_PASSWORD = environ.get("DATABASE_PASSWORD")
DATABASE_NAME = environ.get("DATABASE_NAME")
DATABASE_HOST = environ.get("DATABASE_HOST")
DATABASE_PORT = int(environ.get("DATABASE_PORT"))

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',

        'NAME': DATABASE_NAME,
        'USER': DATABASE_USER,
        'PASSWORD': DATABASE_PASSWORD,
        'HOST': DATABASE_HOST,
        'PORT': DATABASE_PORT,
        'ATOMIC_REQUESTS': True,
        'OPTIONS': {
            'isolation_level': extensions.ISOLATION_LEVEL_SERIALIZABLE,
            'sslmode': 'disable',
        },
    },
}

AUTH_PASSWORD_VALIDATORS = [
    {
        'NAME': 'django.contrib.auth.password_validation.UserAttributeSimilarityValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.CommonPasswordValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.NumericPasswordValidator',
    },
]

LANGUAGE_CODE = 'en-us'

TIME_ZONE = 'UTC'

USE_I18N = False

USE_L10N = False

USE_TZ = False

STATIC_URL = '/static/'
3