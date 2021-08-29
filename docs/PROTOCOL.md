# Flame Command Runtime Protocol(FCRP)
Protocol itself is located in `protocol` folder of the project. It is concisely documented.

However, there are some things you must take into account:
- All messages are transfered via NATS queue groups. Both queue and group must be configured manually.
- They are transfered as raw JSON text.