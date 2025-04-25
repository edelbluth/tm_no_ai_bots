![No AI Bots Middleware with robots.txt for Traefik](./.assets/banner.png)

# No AI Bots Middleware with robots.txt

Repeatedly, bots and crawlers from AI companies generate significant load on web
applications. The `robots.txt` file is often disregarded, or bots do not consider
themselves bound by its directives. This middleware searches for configurable name
components in the `User-Agent` header and blocks access with an HTTP 403 status
code upon a match. If the `robots.txt` file is requested, the middleware responds
with an HTTP 200 status code and returns this standardized response:

```text  
User-agent: *  
Disallow: /  
```

This signals that the bot's visit is unwelcome.

More about the motivation for this plugin:
[Massenhafte Anfragen an Gitea von OpenAIs AI-Crawler »GPTbot« mit Caddy »abwehren« (German Blog Post)](https://j4n.e7h.eu/articles/2025-01-19-openai_chatgpt_gitea_vielzahl_requests)

## Configuration

The only configuration option is a list of name components under the configuration
key `botPatterns`. A good start could be this list:

```yaml
botPatterns:
  - gptbot
  - amazon
  - bytespider
  - openai
  - chatgpt
  - perplexity
  - ccbot
  - google-extended
  - omgili
  - anthropic
  - claude
  - cohere
  - meta-extern
```

### Static configuration

```yaml
experimental:
  plugins:
    TmNoAiBotsPlugin:
      moduleName: "github.com/edelbluth/tm_no_ai_bots"
      version: "v0.2.2"
```

### Dynamic Configuration

```yaml
http:
  middlewares:
    TmNoAiBotsMiddleware:
      plugin:
        TmNoAiBotsPlugin:
          botPatterns:
            - gptbot
            - amazon
            - bytespider
            - openai
            - chatgpt
            - and ... more ...
```

### Example

You can now use the `TmNoAiBotsMiddleware@file` middleware like any other
middleware. You might even use it directly in the entrypoint configuration:

```yaml
entryPoints:
  https:
    address: ":443"
    http:
      middlewares:
        - "TmNoAiBotsMiddleware@file"
      tls: {}
```
