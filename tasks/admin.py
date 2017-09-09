from django.contrib import admin

from .models import Post, Category


class CategoryAdmin(admin.ModelAdmin):
    list_display = ('title', 'parsed')
    list_filter = ('parsed', 'parsed')
    search_fields = ('title', 'slug')


class PostAdmin(admin.ModelAdmin):
    list_display = ('title', 'category', 'date', 'image')
    list_filter = ('category', 'date')
    search_fields = ('title', 'content')
    date_hierarchy = 'date'


admin.site.register(Category, CategoryAdmin)
admin.site.register(Post, PostAdmin)