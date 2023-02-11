# spidee

Discord webhook CLI that enables you to send messages with embeds and files, modify webhook and more.

## Usage

`spidee [global options] command [command options] [arguments...]`

Pass `--id <WEBHOOK ID>` and `--token <WEBHOOK TOKEN>` before command, or set `SPIDEE_WEBHOOK_ID` and `SPIDEE_WEBHOOK_TOKEN` environment variables (recommended). Webhook ID and token can be obtained from the URL (<https://discord.com/api/webhooks/>***id***/***token***). You can get started with:

```shell
spidee send "Hello world!"
```

There are various alternatives to this command, such as piping output, specifying flag, or even sending raw JSON payload. If you want to just send content or payload, `--content` or `--payload` can be omitted as in the example above. When pipe is available, `send` becomes the default executing command. All of these will produce the exact same result:

```shell
spidee send "Hello world!"
spidee send '{"content":"Hello world!"}'
spidee send --content "Hello world!"
spidee send --payload '{"content":"Hello world!"}'
echo "Hello world!" | spidee
echo '{"content":"Hello world!"}' | spidee
```

Argument or pipe will be parsed for JSON; if it succeeds, payload will be sent instead of content. The same logic applies for most of the available commands that accept content or payload.

If you care about the produced message and need output, pass both `--wait` and `--json` flags to command to ensure that message has been created and output JSON. Output can be parsed through a JSON parser:

```shell
spidee send --wait --json --embed-title "My fancy embed" --embed-description "Blah blah blah" --embed-color 0xfcba03 | jq
```

Other commands that support `--json` do not need the wait flag.

### Per-message username & avatar

Pass `--username` and `--avatar-url` to send message with a custom username and avatar:

```shell
spidee send --username "Husky" --avatar-url "https://example.com/husky.png" --content "Woof!"
```

### Embeds

One embed can be sent at a time only. Embed properties such as video and provider are not exposed through the CLI since they have no effect on the embed. Available flags:

```shell
$ spidee send --help
...
   --embed-title value, --et value                                      embed title
   --embed-description value, --ed value                                embed description
   --embed-url value, --eu value                                        embed url
   --embed-timestamp value, --eT value                                  embed timestamp (now|RFC3339 timestamp)
   --embed-color value, --ec value                                      embed color (hex) (default: 0)
   --embed-footer-text value, --eft value                               embed footer text
   --embed-footer-icon value, --efi value                               embed footer icon
   --embed-image-url value, --eiu value                                 embed image url
   --embed-thumbnail-url value, --etu value                             embed thumbnail url
   --embed-author-name value, --ean value                               embed author name
   --embed-author-url value, --eau value                                embed author url
   --embed-author-icon value, --eai value                               embed author icon
   --embed-field value, --ef value [ --embed-field value, --ef value ]  embed field (name,value,inline)
...
```

You can supply an RFC3339 timestamp, or get current timestamp:

```shell
spidee send --embed-content "Embed with current timestamp" --embed-timestamp now
spidee send --embed-content "Embed with some other timestamp" --embed-timestamp 2020-12-09T16:09:53+00:00
```

Embed color in hex:

```shell
spidee send --embed-content "Embed with some other timestamp" --embed-color 0xfcba03
```

Embed fields can be specified multiple times, must follow the `name,value,inline` syntax (inline is optional):

```shell
spidee send --embed-field "Some field","Some value" --embed-field "Another field","Yet another value",true --embed-field "Another field","Yet another value",true
```

### Webhook

spidee is also able to manage some parts of the webhook:

```shell
spidee self modify --username "spidee" --avatar ./avatar.png
spidee self delete
```

Token field will be redacted unless `--no-redact` is passed:

```shell
$ spidee self get --json | jq .token
null
```

## License

spidee is licensed under the MIT License. Please consult the attached LICENSE file for further details.
