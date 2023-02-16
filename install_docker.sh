
# check if docker is installed

if ! [ -x "$(command -v docker)" ]; then
    # install docker
    echo "Installing docker"
    echo "Y" | sudo apt-get update
    echo "Done updating"
    sudo apt-get install \
        ca-certificates \
        curl \
        gnupg \
        lsb-release
    echo "Done installing dependencies"

    sudo mkdir -m 0755 -p /etc/apt/keyrings
    echo "Done creating directory"
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    echo "Done curling"

    echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

    echo "Done adding docker repo"

    sudo apt-get update
    echo "Done updating"
    sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
    echo "Done installing docker"
    sudo docker run hello-world
    echo "Done running hello world"
else
echo "Docker is already installed"
fi

