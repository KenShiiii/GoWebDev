# Persisting your application

To run our application after the terminal session has ended, we must do one of the following:

## Possible options
1. screen
1. init.d
1. upstart
1. system.d

## System.d
1. Create a configuration file
  - cd /etc/systemd/system/
  - sudo nano ```<filename>```.service

```
[Unit]
Description=Go Server
[Service]
ExecStart=/home/<username>/<exepath>
User=root
Group=root
Restart=always
[Install]
WantedBy=multi-user.target
```

1. Add the service to systemd.
  - sudo systemctl enable ```<filename>```.service
1. Activate the service.
  - sudo systemctl start ```<filename>```.service
1. Check if systemd started it.
  - sudo systemctl status ```<filename>```.service
1. Stop systemd if so desired.
  - sudo systemctl stop ```<filename>```.service
  
  