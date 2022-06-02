### About
Converts other image files to png.
Supported formats:
  - [x] webp
  - [ ] jpg/jpeg 
  - [ ] bmp
  - [ ] gif

### Usage

convert a single file named `target.webp`
```bash
 x2png target.webp
```
convert all the `webp` files from target folder
```bash
 x2png ./target-folder
```
remove `webp` after each convert
```bash
 x2png ./target-folder -D
```
### IO
must inputs
    - single file
        - file name
    - directory
        - dir name
    - wildcard
        - regex
    - scan

selectable
    - output directory

### Components
- Engine
- Codec
