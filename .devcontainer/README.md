
# Development Toolbox

The exercise uses dedicated Apps for each Domain task, called Services.
These components use further components, for instance and the Docker Engine API, Docker Engine Compose and so on.

## Getting Started - Development Set-up

To get started with development Microsoft Visual Studio Code as an IDE and a docker container to compile and run the code base for the target system is recommended.
Here a basic Development Image is provided, mainly targeting for a GO development environment, serving as Remote container for Microsoft Visual Studio Code and providing Docker-In-Docker functionality.
In order to use the ```devbox``` as a Remote Container for Visual Studio Code, simply follow the instructions on <https://code.visualstudio.com/docs/remote/containers> for attaching an existing container after it was built and is running <https://code.visualstudio.com/docs/remote/containers#_attaching-to-running-containers>.

### Only a Few Major Steps

1. Get Visual Studio Code plus the Docker Extension and the Remote-Development Extension Pack installed
    - Download Visual Studio Code for Windows as an installer package: <https://code.visualstudio.com/>
    - Or for Linux via Package Manager: <https://code.visualstudio.com/docs/setup/linux>

2. Install Docker and Docker-Compose
    - Windows:
        - Install Docker: [Docker Download without Registration](https://docs.docker.com/docker-for-windows/release-notes/ "Docker Release Homepage")
    - Linux:
        - Install Docker per Distro: e.g. [https://docs.docker.com/engine/install/ubuntu/](https://docs.docker.com/engine/install/ubuntu/ "Install Docker Engine on Ubuntu")
    - For the Docker Compose installation one can read more on <https://docs.docker.com/compose/install/>, but it should come included in the Windows Set-up

3. Check out the Project Repo in BitBucket
    - Clone all the relevant projects code base with ```git```

    ```bash
    cd ~/somewhereToProjects/
    git clone https://github.com/FranzBranntvino/urlsGoShort.git
    ```

4. Start a Remote-Dev-Container as Remote Folder
    - Open Visual Studio Code
    - Click in the lower Left Corner on the tiny little Connect Icon
    - Select: "Remote-Containers: Open Folder in Container ..."
    - Simply select the checked out repo-directory ```urlsGoShort```
    - Sit back and wait while the development Environment gets created for you ...

### Extra Information on the Remote Plugin for Visual Studio Code

In order to get informed how the remote Plugin for ssh and containers work and how a set-up in the ```./devcontainer``` folder via the ```devcontainer.json``` is done, one can check the following links:
    <https://code.visualstudio.com/docs/remote/containers>
    <https://code.visualstudio.com/docs/remote/containers-advanced>
    <https://code.visualstudio.com/docs/remote/containers#_devcontainerjson-reference>
    <https://code.visualstudio.com/docs/editor/variables-reference>
