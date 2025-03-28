FROM tarantool/tarantool:2.11
COPY app.lua /opt/tarantool
CMD ["tarantool", "/opt/tarantool/app.lua"]

