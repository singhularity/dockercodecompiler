echo "Creating Docker Image"
docker build -t ${1:-'compiler_machine'} - < Dockerfile
echo "Retrieving Installed Docker Images"
docker images
