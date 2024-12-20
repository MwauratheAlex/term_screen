#!/usr/bin/env python3

with open("input.bin", "wb") as f:
    # Screen setup (0x01)
    # Width: 85 (0x55), Height: 20(0x14), Color mode: 256 colors (0x02)
    f.write(bytes([0x01, 0x03, 0x55, 0x14, 0x02]))

    # Draw character (0x02)
    # x: 5, y: 3, Color: 10 (0x0A), Character: 'A' (ASCII 0x41)
    f.write(bytes([0x02, 0x04, 0x05, 0x03, 0x0A, 0x41]))

    # Draw line (0x03)
    # Start: (0, 0), End: (15, 15), Color: 2 (0x02), Character: '*' (ASCII 0x2A)
    f.write(bytes([0x03, 0x06, 0x00, 0x00, 0x0F, 0x0F, 0x02, 0x2A]))

    # Render text(0x04)
    # Start: (2, 2), Color: 1 (0x01), Text: "Hello" (ASCII)
    f.write(bytes([0x04, 0x28, 0x02, 0x02, 0x01]) +
            b"Merry Christmas from terminal screen.")

    # # Cursor movement (0x05)
    # # Move cursor to (15, 10)
    f.write(bytes([0x05, 0x02, 0x0F, 0x0A]))

    # # Draw at cursor (0x06)
    # # Character: 'X' (ASCII 0x58), Color: 5 (0x05)
    f.write(bytes([0x06, 0x02, 0x58, 0x05]))

    # # Clear screen (0x07)
    # # No additional data
    # f.write(bytes([0x07, 0x00]))

    # End of file (0xFF)
    # No additional data
    f.write(bytes([0xFF, 0x00]))
