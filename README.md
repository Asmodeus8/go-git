# Go-Git

An educational implementation of core Git storage concepts in Go.

Implemented commands:
- `init` creates the `.git` object/ref layout and HEAD
- `hash-object <file>` creates a Git-compatible blob payload, SHA-1 object ID and zlib-compressed loose object
- `cat-file <sha>` reads/decompresses a loose object

```bash
go run . init
go run . hash-object README.md
go run . cat-file <sha>
```

Implemented independently as a systems-programming exercise by Adewale Babalola.
