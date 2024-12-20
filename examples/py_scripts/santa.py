#!/usr/bin/env python3


with open("santa.bin", "wb") as f:
    # Screen setup (0x01)
    # Width: 85 (0x55), Height: 20 (0x14), Color mode: 256 (0x02)
    f.write(bytes([0x01, 0x03, 0x55, 0x14, 0x02]))

    # Clear screen (0x07)
    f.write(bytes([0x07, 0x00]))

    # Draw a small moon at top-left corner: (5,1), char='o', color=6
    f.write(bytes([0x04, 0x04, 0x05, 0x01, 0x06, ord('o')]))

    # Add some stars in the night sky
    # Positions chosen somewhat arbitrarily
    star_positions = [(10, 2), (20, 3), (30, 1), (50, 2), (60, 1),
                      (25, 4), (70, 2), (15, 3), (45, 4), (80, 1)]
    for (sx, sy) in star_positions:
        # color=7 (bright white for stars)
        f.write(bytes([0x04, 0x04, sx, sy, 0x07, ord('*')]))

    # "Season's Greetings!" message at (30,6), color=2 (green)
    message = "Season's Greetings!"
    x_msg, y_msg, color_msg = 30, 6, 2
    length_msg = 3 + len(message)
    f.write(bytes([0x04, length_msg, x_msg, y_msg, color_msg]) +
            message.encode('ascii'))


    santa_color = 3
    x_start = 10

    santa_lines = [
            "    __     _  __ ",
            "    | \\__ `\\O/  `--  {}    \\}    {/",
            "    \\    \\_(~)/______/=____/=____/=*",
            "     \\=======/    //\\  >\\/> || \\> ",
            "    ----`---`---  `` `` ```` `` ``",
        ]

    # Write each line with a separate command
    # y coordinates start at 8 and increment
    for i, line in enumerate(santa_lines):
        y = 8 + i
        data = line.encode('ascii')
        length = 3 + len(data)  # x,y,color plus the text
        f.write(bytes([0x04, length, x_start, y, santa_color]) + data)

    # City skyline at the bottom (line 18 and 19)
    # Use color=8 (gray)
    skyline = "#" * 85
    x_skyline, y_skyline, color_skyline = 0, 18, 8
    length_skyline = 3 + len(skyline)
    f.write(bytes([0x04, length_skyline, x_skyline, y_skyline,
            color_skyline]) + skyline.encode('ascii'))

    # Another line of skyline at line 19
    x_lower, y_lower, color_lower = 0, 19, 8
    length_lower = 3 + len(skyline)
    f.write(bytes([0x04, length_lower, x_lower, y_lower,
            color_lower]) + skyline.encode('ascii'))

    # End of file (0xFF)
    f.write(bytes([0xFF, 0x00]))
