from django import forms
from django.contrib.auth.forms import ReadOnlyPasswordHashField

from .models import User


class UserChangeForm(forms.ModelForm):
    password = ReadOnlyPasswordHashField()

    class Meta:
        model = User
        fields = ('nome', 'email', 'password', 'is_active', 'is_admin', 'is_staff')
