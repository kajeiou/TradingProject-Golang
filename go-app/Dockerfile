# Utilisez l'image recommandée par la documentation de fresh
FROM golang:latest

WORKDIR /app

# Copiez les fichiers de votre application dans le conteneur
COPY . .

# Construisez votre application Go
RUN go build -o main

# Exécutez votre application avec Air pour le hot reload
CMD ["main"]
