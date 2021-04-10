# weatherfetch

### A weather status update for your terminal.

weatherfetch is a terminal application meant to be run as part of your .bashrc startup, built in Go.

weatherfetch will grab your location based on the IP address, run a GET request for weather, and display the result with a condition-specific ASCII art image to welcome you.

Location data is only used for your specific request, and no information is stored in a database.
