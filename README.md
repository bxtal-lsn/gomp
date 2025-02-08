# GOMP

## Overview

Multipass-CLI is a lightweight Go-based command-line interface (CLI) tool that simplifies the management of Ubuntu instances in Multipass. This CLI provides an intuitive way to create, manage, and configure Multipass instances using go install or prebuilt binaries.

## Features

- Launch instances with configurable memory, storage, and CPU allocation.
 
- Support cloud-init provisioning per instance or using a default file.
 
- Start, stop, and delete instances easily.
 
- Open an interactive shell inside an instance.
 
- Execute commands inside instances without logging in.
 
- Retrieve instance information (IP, memory, disk usage, etc.).
 
- Recover deleted instances before purging.
 
- Mount/unmount host folders into instances.
 
- Set aliases for frequently used commands.

## Installation

[1] Install via Go
```bash
go install github.com/your-username/multipass-cli@latest
```

This installs the binary to $GOPATH/bin/multipass-cli.

[2] Download Prebuilt Binaries

Visit the GitHub Releases page and download the appropriate binary for your OS.

**Linux/macOS**
```bash
curl -L -o gomp https://github.com/bxtal-lsn/gomp-multipass/releases/download/vo.0.1/gomp-multipass-linux
chmod +x gomp
mv gomp /usr/local/bin/
```

**Windows**

Download gomp.exe

Move it to a folder in PATH

Run:
```ps1
gomp.exe
```

## Usage

1️⃣ Launch a New Instance

multipass-cli launch

Follow the prompts to specify names, memory, storage, CPUs, and cloud-init configuration.

2️⃣ Start an Instance

multipass-cli start my-instance

3️⃣ Stop an Instance

multipass-cli stop my-instance

4️⃣ Delete an Instance

multipass-cli delete my-instance

To delete all instances:

multipass-cli delete --all

5️⃣ Open a Shell in an Instance

multipass-cli shell my-instance

6️⃣ Execute Commands in an Instance

multipass-cli exec my-instance ls -l

7️⃣ Get Instance Info

multipass-cli info my-instance

8️⃣ Recover a Deleted Instance

multipass-cli recover my-instance

9️⃣ Mount a Folder into an Instance

multipass-cli mount /host/path my-instance:/mnt/path

🔟 Unmount a Folder

multipass-cli unmount my-instance:/mnt/path

🔹 Create an Alias for a Command

multipass-cli alias update="sudo apt update && sudo apt upgrade -y"

## Cloud-Init Example Files

### Basic Cloud-Init (cloudinit-basic.yml)

#cloud-config
package_update: true
packages:
  - htop
  - curl
  - tree

### Docker Cloud-Init (cloudinit-docker.yml)

#cloud-config
package_update: true
package_upgrade: true
packages:
  - docker.io
runcmd:
  - systemctl enable --now docker
  - usermod -aG docker ubuntu

### Development Setup

git clone https://github.com/your-username/multipass-cli.git
cd multipass-cli
go mod tidy
go build -o multipass-cli

Run the CLI locally:

go run . launch
