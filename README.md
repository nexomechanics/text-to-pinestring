# Introduction
A simple API service that converts structured text into pine script string format.

It's useful for cases where you want to write a long tooltip for an input or other applications like alert messages in pine script.

Pine script has a weird syntax when it comes to multiline text and it can be hard to track. Hopefully this app makes it easier to convert regular text into pine script. If you do this a lot, it should save you a lot of time.

You can also do this with AI / ChatGPT, but the input text should be perfect, otherwise the mistakes will be inherited or you will risk undesired porting. Often time, when doing with task with AI, you need a couple of tries before getting right, and that's a waste of tokens, context, and most importantly, time!

## Pine script strings limitations
Tradingview provides all uses of the escape character "\" in pine script:
- \n — newline
- \t — tab / table 
- \\ — backslash
- \" — double quote
- \' — single quote

No bold, italic, color, size, alignment, or markdown possible directly from strings.

### Tab stop rule (table alignment)
Pine script renders string in a monospace-like font. From my analysis, tab stops are fixed at every 8 characters.

If two rows have different length values before a \t, they land on different tab stops and the columns look misaligned.

The fix introduced in this app is adding padding to every cell to the next multiple of 8 with spaces before the \t. In other words, calculate the longest value per column, round up to the next multiple of 8, and pad everything to match.

https://www.tradingview.com/pine-script-docs/concepts/strings/#escape-sequences

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
