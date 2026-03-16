# Introduction
A simple API service that converts structured text into pine script string format.

It's useful for cases where you want to write a long tooltip for an input or other applications like alert messages in pine script.

Pine script has a weird syntax when it comes to multiline text and it can be hard to track. Hopefully this app makes it easier to convert regular text into pine script. If you do this a lot, it should save you a lot of time.

You can also do this with AI / ChatGPT, but the input text should be perfect, otherwise the mistakes will be inherited or you will risk undesired porting. Often time, when doing with task with AI, you need a couple of tries before getting right, and that's a waste of tokens, context, and most importantly, time!

# API

**POST** `/convert`

Accepts a JSON array of blocks, returns a Pine Script string.

## Block types

### `line`
Plain text. Newlines in the content — whether typed (Enter key) or as `\n` — are preserved as-is in the output.
```json
{"type": "line", "content": "First line\nSecond line\n\nNew paragraph"}
```

### `table`
Tab-separated columns with a header row.
```json
{"type": "table", "headers": ["Name", "Value", "Description"], "rows": [
    ["Foo", "1", "First item"],
    ["Bar", "2", "Second item"],
    ["Baz", "3", "Third item"]
]}
```

## Full example

Request:
```json
[
    {"type": "line", "content": "Section Title\n\n1. First point.\n2. Second point.\n"},
    {"type": "table", "headers": ["Name", "Value", "Description"], "rows": [
        ["Foo", "1", "First item"],
        ["Bar", "2", "Second item"],
        ["Baz", "3", "Third item"]
    ]}
]
```

Response:
```json
{"result": "\"Section Title\n\n1. First point.\n2. Second point.\nName    \tValue   \tDescription\nFoo     \t1       \tFirst item\nBar     \t2       \tSecond item\nBaz     \t3       \tThird item\n\""}
```
