from django.contrib import admin

from .forms import UserChangeForm
from .models import User


class UserAdmin(admin.ModelAdmin):
    form = UserChangeForm

admin.site.register(User, UserAdmin)

