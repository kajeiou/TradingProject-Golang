echo "|| Cleaning... ||"; sudo docker-compose down -v;
echo "|| Running... ||"; sudo docker-compose up -d;
echo "|| Done. ||"; sudo docker ps;
echo "Run PSQL with : sh db.sh"
echo "In PSQL, use \dt to see all tables"