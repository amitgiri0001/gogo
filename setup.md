# How did I made the system
## Step 1 - Env
```
docker pull golang
```

```
docker run -itd golang
```

```
››› docker container ps                                                                                     
CONTAINER ID   IMAGE     COMMAND   CREATED          STATUS          PORTS     NAMES
433faba83c9c   golang    "bash"    37 minutes ago   Up 12 minutes             unruffled_haslett
```

## Step 2 - Editor
- Remote connection from VScode to docker container

## Setp 3 - Git
### Add SSH keys
```
ssh-keygen -t rsa -b 4096 -C "**********@gmail.com"
```

```
eval "$(ssh-agent -s)"
```

```
ssh-add ~/.ssh/id_rsa
```
Add to Github
```
cat ~/.ssh/id_rsa.pub
```

```
git clone git@github.com:amitgiri0001/gogo.git
```

## Setp 4 - Shell (ZSH)
```
apt update && apt upgrade
apt install zsh
```