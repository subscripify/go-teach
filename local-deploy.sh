docker build -t dev/subscripify-new-color-server:latest .
docker tag dev/subscripify-new-color-server subscripifycontreg.azurecr.io/dev/subscripify-new-color-server:latest
docker push subscripifycontreg.azurecr.io/dev/subscripify-new-color-server:latest


