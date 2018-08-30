from django.views.generic import TemplateView
from django.http import JsonResponse


# Create your views here.
class HealthCheck(TemplateView):
    def get(self, request, **kwargs):
        return JsonResponse({'status':'ok'},safe=False)
