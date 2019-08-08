If run this docker standalone, it is needed to add the following config to the ```docker run``` command.

````--sysctl net.ipv4.ping_group_range="0   2147483647"````

  ping:
    build:
      context: ./ping
    sysctls:
      net.ipv4.ping_group_range: "0 2147483647"
    labels:
      - "faas.name=ping"
      - "faas.port=8080"
    volumes:
      - ./ping:/go/src/github.com/tangyuan2014/modularfianicefaas/ping