# TweetHub-CLI

tweethub es una herramienta de línea de comandos (CLI) que te permite interactuar con Twitter de manera automatizada. Utiliza un enfoque de navegación sin cabeza (headless browsing) con la biblioteca chromedp para realizar acciones como dar "like", "retweet", "quote", "follow" y más.

## Uso

### Follow
Para seguir a un usuario en Twitter, utiliza el comando **follow**:
```bash
tweethub follow --username <nombre-de-usuario>
```

### Like
Para dar "like" a un tweet, utiliza el comando **like**:
```bash
tweethub like --url <URL-del-tweet>
```

### Quote
Para responder a un tweet, utiliza el comando **quote**:
```bash
tweethub quote --url <URL-del-tweet>
```

### Repost
Para repostear a un tweet, utiliza el comando **repost**:
```bash
tweethub repost --url <URL-del-tweet>
```

### Tweet
Para publicar un nuevo tweet, utiliza el comando **tweet**:
```bash
tweethub tweet --message "Contenido del tweet"
```

