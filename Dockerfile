FROM scratch
MAINTAINER Brian Hechinger <wonko@4amlunch.net>

ADD auth-esi-poller-linux-amd64 auth-esi-poller
VOLUME /etc/chremoas

ENTRYPOINT ["/auth-esi-poller", "--configuration_file", "/etc/chremoas/auth-bot.yaml"]