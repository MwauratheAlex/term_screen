from enum import Enum


with open("input.bin", "wb") as f:
    # Screen setup (0x01)
    # Width: 85 (0x55), Height: 20(0x14), Color mode: 256 colors (0x02)
    f.write(bytes([0x01, 0x03, 0x55, 0x14, 0x02]))

    # Draw character (0x02)
    # x: 5, y: 3, Color: 10 (0x0A), Character: 'A' (ASCII 0x41)
    # f.write(bytes([0x02, 0x04, 0x05, 0x03, 0x0A, 0x41]))

    # Draw line (0x03)
    # Start: (0, 0), End: (15, 15), Color: 2 (0x02), Character: '-' (ASCII 0x2D)
    # f.write(bytes([0x03, 0x06, 0x00, 0x00, 0x0F, 0x0F, 0x02, 0x2D]))

    # Render text (0x04)
    # Start: (2, 2), Color: 1 (0x01), Text: "Hello" (ASCII)
    # f.write(bytes([0x04, 0x08, 0x02, 0x02, 0x01]) + b"Hello")

    # # Cursor movement (0x05)
    # # Move cursor to (10, 10)
    # f.write(bytes([0x05, 0x02, 0x0A, 0x0A]))

    # # Draw at cursor (0x06)
    # # Character: 'X' (ASCII 0x58), Color: 5 (0x05)
    # f.write(bytes([0x06, 0x02, 0x58, 0x05]))

    # # Clear screen (0x07)
    # # No additional data
    # f.write(bytes([0x07, 0x00]))

    # End of file (0xFF)
    # No additional data
    f.write(bytes([0xFF, 0x00]))


class Commands(Enum):
    SETUP = 0x01
    DRAW_CHAR = 0x02
    DRAW_LINE = 0x03
    RENDER_TEXT = 0x04
    MOV_CURSOR = 0x05
    DRAW_AT_CURSOR = 0x06
    CLEAR_SCREEN = 0x07
    EOF = 0xFF


def makeCommand(command):
    match command:
        case Commands.SETUP:
            pass
        case Commands.DRAW_CHAR:
            pass
        case Commands.DRAW_LINE:
            pass
        case Commands.RENDER_TEXT:
            pass
        case Commands.MOV_CURSOR:
            pass
        case Commands.DRAW_AT_CURSOR:
            pass
        case Commands.CLEAR_SCREEN:
            pass
        case Commands.EOF:
            return bytes([0xFF, 0x00])
        case _:
            print("Unknown command")
