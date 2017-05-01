# fitgo
Go utilities to parse text in oneline manner

```bash
./parse.sh 'Lines("https://www.cloudflare.com/ips-v4").Grep(`^103`).Echo()'
curl -s https://www.cloudflare.com/ips-v4 | ./parse.sh 'Stdin().Grep(`^103`).Echo()'
```