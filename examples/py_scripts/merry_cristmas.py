#!/usr/bin/env python3

with open("xmas.bin", "wb") as f:
    # Screen setup (0x01)
    # Width: 85 (0x55), Height: 20 (0x14), Color mode: 256 colors (0x02)
    f.write(bytes([0x01, 0x03, 0x55, 0x14, 0x02]))

    # Clear screen (0x07)
    # No additional data
    f.write(bytes([0x07, 0x00]))

    # Let's draw a simple ASCII Christmas tree:
    # Center the tree roughly around x=40 for aesthetics.
    # Y coordinates start at 2 for the top star and go down line by line.
    #
    #    (y=2)                 *
    #    (y=3)                ***
    #    (y=4)               *****
    #    (y=5)              *******
    #    (y=6)             *********
    #    (y=7)                 |
    #    (y=8)                 |
    #
    # We'll assume:
    # - Star and tree use color index = 3 for the star (yellow) and 3 for the rest (just keeping it consistent).
    # - Trunk uses color index = 5 (arbitrary choice).
    #
    # Command 0x04: Render text
    # Data: [x, y, color, ASCII chars]

    # Top star at (40,2)
    # Data bytes: x=0x28, y=0x02, color=0x03, text="*"
    # Data length = 3 (coords+color) +1 (char) =4
    f.write(bytes([0x04, 0x04, 0x28, 0x02, 0x03, ord('*')]))

    # Row of 3 stars at (39,3)
    # Data: x=0x27, y=0x03, color=0x03, text="***" (3 chars)
    # Length=3+3=6
    f.write(bytes([0x04, 0x06, 0x27, 0x03, 0x03, ord('*'), ord('*'), ord('*')]))

    # Row of 5 stars at (38,4)
    # Data: x=0x26, y=0x04, color=0x03, text="*****"
    # Length=3+5=8
    f.write(bytes([0x04, 0x08, 0x26, 0x04, 0x03, ord(
        '*'), ord('*'), ord('*'), ord('*'), ord('*')]))

    # Row of 7 stars at (37,5)
    # Data: x=0x25, y=0x05, color=0x03, text="*******"
    # Length=3+7=10
    f.write(bytes([0x04, 0x0A, 0x25, 0x05, 0x03] + [ord('*')] * 7))

    # Row of 9 stars at (36,6)
    # Data: x=0x24, y=0x06, color=0x03, text="*********"
    # Length=3+9=12
    f.write(bytes([0x04, 0x0C, 0x24, 0x06, 0x03] + [ord('*')] * 9))

    # Trunk at (40,7)
    # Data: x=0x28, y=0x07, color=0x05, text="|"
    # Length=3+1=4
    f.write(bytes([0x04, 0x04, 0x28, 0x07, 0x05, ord('|')]))

    # Trunk at (40,8)
    # Same parameters, different y
    f.write(bytes([0x04, 0x04, 0x28, 0x08, 0x05, ord('|')]))

    # Now a Christmas message below the tree:
    # "Merry Christmas and Happy Holidays!"
    # Let's place it at (20,10) with color=1 (red)
    message = "Merry Christmas and Happy Holidays!"
    # Count data length: 3 (x,y,color) + len(message)=3+35=38
    f.write(bytes([0x04, 0x26, 0x14, 0x0A, 0x01]) + message.encode('ascii'))

    # Add a border around the screen for a decorative touch using draw line (0x03):
    # Draw line data format: [x1, y1, x2, y2, color, char]

    # Top border line: from (0,0) to (84,0), color=2, char='-'
    # Data length=6
    f.write(bytes([0x03, 0x06, 0x00, 0x00, 0x54, 0x00, 0x02, ord('-')]))

    # Bottom border line: from (0,19) to (84,19)
    f.write(bytes([0x03, 0x06, 0x00, 0x13, 0x54, 0x13, 0x02, ord('-')]))

    # Left border line: from (0,0) to (0,19)
    f.write(bytes([0x03, 0x06, 0x00, 0x00, 0x00, 0x13, 0x02, ord('|')]))

    # Right border line: from (84,0) to (84,19)
    f.write(bytes([0x03, 0x06, 0x54, 0x00, 0x54, 0x13, 0x02, ord('|')]))

    # End of file (0xFF)
    # No additional data
    f.write(bytes([0xFF, 0x00]))
