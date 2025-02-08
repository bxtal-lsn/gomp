# gomp (go-multipass)

## Overview

gomp is a lightweight go-based cli tool that simplifies the management of Ubuntu instances in Multipass.

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
go install github.com/bxtal-lsn/gomp@latest
```

This installs the binary to $GOPATH/bin/gomp.

[2] Download Prebuilt Binaries

Visit the [GitHub Releases page](https://github.com/bxtal-lsn/gomp/releases/) and download the appropriate binary for your OS, and move it to PATH.

## Usage

1 launch a New Instance

gomp launch

Follow the prompts to specify names, memory, storage, CPUs, and cloud-init configuration.

2️⃣ Start an Instance

gomp start my-instance

3️⃣ Stop an Instance

gomp stop my-instance

4️⃣ Delete an Instance

gomp delete my-instance

To delete all instances:

gomp delete --all

5️⃣ Open a Shell in an Instance

gomp shell my-instance

6️⃣ Execute Commands in an Instance

gomp exec my-instance ls -l

7️⃣ Get Instance Info

gomp info my-instance

8️⃣ Recover a Deleted Instance

gomp recover my-instance

9️⃣ Mount a Folder into an Instance

gomp mount /host/path my-instance:/mnt/path

🔟 Unmount a Folder

gomp unmount my-instance:/mnt/path

[11] Create an Alias for a Command

gomp alias update="sudo apt update && sudo apt upgrade -y"

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
