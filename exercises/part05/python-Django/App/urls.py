# howdy/urls.py
from django.conf.urls import url
#from rest_framework.urlpatterns import format_suffix_patterns
from App import views


urlpatterns = [
    url(r'^health-check$', views.HealthCheck.as_view(),name="create"),
]

#urlpatterns = format_suffix_patterns(urlpatterns)
