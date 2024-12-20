#!/usr/bin/env python3

with open("animated_santa.bin", "wb") as f:
    # Screen setup (0x01)
    # Width: 85 (0x55), Height: 20 (0x14), Color mode: 256 (0x02)
    f.write(bytes([0x01, 0x03, 0x55, 0x14, 0x02]))

    def draw_static_scene():
        # Draw a small moon at top-left corner: (5,1), char='o', color=6
        f.write(bytes([0x04, 0x04, 0x05, 0x01, 0x06, ord('o')]))

        # Add some stars
        star_positions = [(10, 2), (20, 3), (30, 1), (50, 2), (60, 1),
                          (25, 4), (70, 2), (15, 3), (45, 4), (80, 1)]
        for (sx, sy) in star_positions:
            # color=7 (bright white)
            f.write(bytes([0x04, 0x04, sx, sy, 0x07, ord('*')]))

        # "Season's Greetings!" message at (30,6), color=2 (green)
        message = "Season's Greetings!"
        x_msg, y_msg, color_msg = 30, 6, 2
        length_msg = 3 + len(message)
        f.write(bytes([0x04, length_msg, x_msg, y_msg, color_msg]) +
                message.encode('ascii'))

        # City skyline (color=8 gray)
        skyline = "#" * 85
        x_skyline, y_skyline, color_skyline = 0, 18, 8
        length_skyline = 3 + len(skyline)
        f.write(bytes([0x04, length_skyline, x_skyline, y_skyline,
                color_skyline]) + skyline.encode('ascii'))

        x_lower, y_lower, color_lower = 0, 19, 8
        length_lower = 3 + len(skyline)
        f.write(bytes([0x04, length_lower, x_lower, y_lower,
                color_lower]) + skyline.encode('ascii'))

    santa_color = 3
    santa_lines = [
        "    __     _  __ ",
        "    | \\__ `\\O/  `--  {}    \\}    {/",
        "    \\    \\_(~)/______/=____/=____/=*",
        "     \\=======/    //\\  >\\/> || \\> ",
        "    ----`---`---  `` `` ```` `` ``",
    ]

    def draw_sled(x_start):
        # Draw the sled at (x_start, y)
        for i, line in enumerate(santa_lines):
            y = 8 + i
            data = line.encode('ascii')
            length = 3 + len(data)
            f.write(bytes([0x04, length, x_start, y, santa_color]) + data)

    def erase_sled(x_start):
        # Overwrite the sled's old position with spaces to restore background (assumed empty sky).
        for i, line in enumerate(santa_lines):
            y = 8 + i
            blank_line = ' ' * len(line)
            length = 3 + len(blank_line)
            # Use color=0 (or some background color) for the blanks
            f.write(bytes([0x04, length, x_start, y, 0x00]) + blank_line.encode('ascii'))

    def delay_millis(millis):
        for _ in range(millis):
            f.write(bytes([0x08, 0x00]))

    # Initially, clear the screen once
    f.write(bytes([0x07, 0x00]))

    # Draw the static scene once
    draw_static_scene()

    # Starting position of the sled
    old_x = 10
    draw_sled(old_x)

    # Move the sled from x=11 to x=50
    for x in range(11, 51):
        delay_millis(30)
        # Erase the sled at the old position
        erase_sled(old_x)
        # Draw the sled at the new position
        draw_sled(x)
        old_x = x

    # Render text(0x04)
    # Start: (16, 10), Color: 1 (0x01), Text: "Hello" (ASCII)
    f.write(bytes([0x04, 0x28, 0x16, 0x10, 0x01]) +
            b"Merry Christmas from terminal screen.")

    # Final delay before ending
    delay_millis(1000)


    # End of file
    f.write(bytes([0xFF, 0x00]))
