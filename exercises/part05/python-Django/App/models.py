from django.db import models

# Create your models here.

class HealthCheck(models.Model):
    status = models.CharField(max_length=50)
