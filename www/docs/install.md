# Install 

You can find the pre-compiled binaries on the release page [releases](https://github.com/vdbulcke/cert-monitor/releases)


## Getting Latest Version 


```sh
TAG=$(curl https://api.github.com/repos/vdbulcke/cert-monitor/releases/latest  |jq .tag_name -r )
VERSION=$(echo $TAG | cut -d 'v' -f 2)
```

!!! info
    You will need `jq` and `curl` in your `PATH`

!!! warning 
    If you don't have `jq` you can replace the `TAG` and `VERSION` variable manually in the command below


## MacOS 

=== "Intel"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/cert-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/cert-monitor_${VERSION}_Darwin_x86_64.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "cert-monitor_${VERSION}_Darwin_x86_64.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./cert-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install cert-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv cert-monitor /usr/local/bin/
      ```
      
    1. *Checksum:* You can validate the checksum of the downloaded packages by comparing the hash in the `checksums.txt` that comes with each release, with the computed hash of the downloaded archive: 

        1. Download `checksums.txt`:
          ```sh
          curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/checksums.txt"

          ```
        1. Compute Checksum:  
          ```sh
          sha256sum  "cert-monitor_${VERSION}_Darwin_x86_64.tar.gz"

          ```
        1. Compare against `checksums.txt`: 
          ```sh
          grep "cert-monitor_${VERSION}_Darwin_x86_64.tar.gz" checksums.txt
          ```


=== "ARM"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/cert-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/cert-monitor_${VERSION}_Darwin_amr64.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "cert-monitor_${VERSION}_Darwin_amr64.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./cert-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install cert-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv cert-monitor /usr/local/bin/
      ```
      
    1. *Checksum:* You can validate the checksum of the downloaded packages by comparing the hash in the `checksums.txt` that comes with each release, with the computed hash of the downloaded archive:

        1. Download `checksums.txt`:
          ```sh
          curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/checksums.txt"

          ```
        1. Compute Checksum:  
          ```sh
          sha256sum  "cert-monitor_${VERSION}_Darwin_amr64.tar.gz"

          ```
        1. Compare against `checksums.txt`: 
          ```sh
          grep "cert-monitor_${VERSION}_Darwin_amr64.tar.gz" checksums.txt
          ```

=== "ARM"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/cert-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/cert-monitor_${VERSION}_Darwin_all.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "cert-monitor_${VERSION}_Darwin_all.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./cert-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install cert-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv cert-monitor /usr/local/bin/
      ```
      
    1. *Checksum:* You can validate the checksum of the downloaded packages by comparing the hash in the `checksums.txt` that comes with each release, with the computed hash of the downloaded archive: 

        1. Download `checksums.txt`:
          ```sh
          curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/checksums.txt"

          ```
        1. Compute Checksum:  
          ```sh
          sha256sum  "cert-monitor_${VERSION}_Darwin_all.tar.gz"

          ```
        1. Compare against `checksums.txt`: 
          ```sh
          grep "cert-monitor_${VERSION}_Darwin_all.tar.gz" checksums.txt
          ```




## Linux 


=== "Intel"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/cert-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/cert-monitor_${VERSION}_Linux_x86_64.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "cert-monitor_${VERSION}_Linux_x86_64.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./cert-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install cert-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv cert-monitor /usr/local/bin/
      ```
      
    1. *Checksum:* You can validate the checksum of the downloaded packages by comparing the hash in the `checksums.txt` that comes with each release, with the computed hash of the downloaded archive:

        1. Download `checksums.txt`:
          ```sh
          curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/checksums.txt"

          ```
        1. Compute Checksum:  
          ```sh
          sha256sum  "cert-monitor_${VERSION}_Linux_x86_64.tar.gz"

          ```
        1. Compare against `checksums.txt`: 
          ```sh
          grep "cert-monitor_${VERSION}_Linux_x86_64.tar.gz" checksums.txt
          ```


=== "ARM"
    1. Download the binary  from the [releases](https://github.com/vdbulcke/cert-monitor/releases) page:
      ```sh
      curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/cert-monitor_${VERSION}_Linux_amr64.tar.gz"
      
      ```
    1. Extract Binary:
      ```sh
      tar xzf "cert-monitor_${VERSION}_Linux_amr64.tar.gz"
      ```
    1. Check Version: 
      ```sh
      ./cert-monitor version
      ```
    1. Install in your `PATH`: 
      ```sh
      sudo install cert-monitor /usr/local/bin/
      ```
      Or
      ```sh
      sudo mv cert-monitor /usr/local/bin/
      ```
      
    1. *Checksum:* You can validate the checksum of the downloaded packages by comparing the hash in the `checksums.txt` that comes with each release, with the computed hash of the downloaded archive:

        1. Download `checksums.txt`:
          ```sh
          curl -LO "https://github.com/vdbulcke/cert-monitor/releases/download/${TAG}/checksums.txt"

          ```
        1. Compute Checksum:  
          ```sh
          sha256sum  "cert-monitor_${VERSION}_Linux_amr64.tar.gz"

          ```
        1. Compare against `checksums.txt`: 
          ```sh
          grep "cert-monitor_${VERSION}_Linux_amr64.tar.gz" checksums.txt
          ```


## Windows 


=== "Intel"
    1. Download the binary `cert-monitor_[VERSION]_Windows_x86_64.zip`  from the [releases](https://github.com/vdbulcke/cert-monitor/releases) page
     
    1. Unzip the Binary

    1. Check Version: 
      ```sh
      ./cert-monitor.exe version
      ```

