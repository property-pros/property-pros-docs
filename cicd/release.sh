#!/bin/sh

source .env
echo $CONTAINER_PORT
UNIQUE_TAG=$(date +%s)

gcloud config set project vireo-401203

gcloud builds submit --tag us-west1-docker.pkg.dev/vireo-401203/pp/ppd:$UNIQUE_TAG

# gcloud run deploy property-pros-docs --no-traffic --port $CONTAINER_PORT --tag us-west1-docker.pkg.dev/vireo-401203/pp/ppd:$UNIQUE_TAG --platform managed --region us-west1 --allow-unauthenticated
gcloud run deploy property-pros-docs --use-http2  --port 8020 --platform managed --image us-west1-docker.pkg.dev/vireo-401203/pp/ppd:$UNIQUE_TAG --region us-west1 --allow-unauthenticated
echo “deployed to gcloud run”

gcloud artifacts packages list --repository=pp --location=us-west1

# us-west1-docker.pkg.dev/vireo-401203/pp/ppd:1696559415property-pros-docs
# source .env

# # Generate a unique tag based on the current timestamp
# UNIQUE_TAG=$(date +%s)

# gcloud builds submit --tag us-west1-docker.pkg.dev/vireo-401203/pp/property-pros-docs:$UNIQUE_TAG

# gcloud config set project vireo-401203

# docker push us-west1-docker.pkg.dev/vireo-401203/pp/property-pros-docs:$UNIQUE_TAG
# docker push us-west1-docker.pkg.dev/vireo-401203/pp/property-pros-docs:latest
# # docker push us-docker.pkg.dev/vireo-401203/pp/property-pros-docs:$UNIQUE_TAG

# gcloud run deploy property-pros-docs --port "$CONTAINER_PORT" --tag us-west1-docker.pkg.dev/vireo-401203/pp/property-pros-docs:$UNIQUE_TAG --platform managed --region us-west1 --allow-unauthenticated --project vireo-401203
# echo "deployed to gcloud run"

# gcloud artifacts packages list --repository=property-pros --location=us-west1 --project=vireo-401203