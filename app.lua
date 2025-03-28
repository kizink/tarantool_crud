#!/usr/bin/env tarantool

box.cfg {
    listen = 3301
}

box.once('init', function()
    -- box.schema.user.create('sample_user', 
    --     {password = 'password'}
    -- )
    -- box.schema.user.grant('sample_user', 'read,write', 'universe')
    
    box.schema.space.create('kv', {
        format = {
            {name = 'key', type = 'string'},
            {name = 'value', type = '*'}
            --{name = 'value', type = '*'}
        }
    })

    box.space.kv:create_index('primary', {
        type = 'tree',
        parts = {'key'}
    })
end)