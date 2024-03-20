import hashlib
from django.db import models
from django.contrib.auth.models import BaseUserManager, AbstractBaseUser, PermissionsMixin


def upload_directory_path(instance, filename):
    return '{}/{}/{}.{}'.format(
        instance.__class__.__name__.lower(),
        str(instance.id),
        hashlib.sha224(filename.encode()).hexdigest(),
        filename.split(".")[-1:]
    )


class MyUserManager(BaseUserManager):
    def create_user(self, email, nome, password=None):
        if not email:
            raise ValueError('Users must have an email address')

        user = self.model(
            email=self.normalize_email(email),
            nome=nome
        )

        user.set_password(password)
        user.save(using=self._db)
        return user

    def create_superuser(self, nome, email, password):
        if not email:
            raise ValueError('Users must have an email address')

        user = self.model(
            email=self.normalize_email(email),
            nome=nome,
            is_superuser=True
        )

        user.set_password(password)
        user.is_admin = True
        user.is_staff = True
        user.is_active = True
        user.save(using=self._db)
        return user


class User(AbstractBaseUser, PermissionsMixin):
    email = models.EmailField(max_length=255, unique=True)
    nome = models.CharField(max_length=300)
    avatar = models.ImageField(null=True, upload_to=upload_directory_path)
    is_active = models.BooleanField(default=False)
    is_admin = models.BooleanField(default=False)
    is_staff = models.BooleanField(default=False)

    objects = MyUserManager()

    USERNAME_FIELD = 'email'
    REQUIRED_FIELDS = ['nome']

    class Meta:
        ordering = ['nome']

    def get_full_name(self):
        return self.nome or self.email

    def __str__(self):
        return f'{self.nome} - {self.email}'

