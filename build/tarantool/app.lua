-- configure tarantool
box.cfg{
    listen = 3301
}

box.once('init', function()
    s = box.schema.space.create('sessions')
    s:format({
        {name = 'session_id', type = 'string'},
        {name = 'user_id', type = 'integer'}
    })

    s:create_index('primary', {type = 'HASH', parts = {'session_id'}})
    print("tarantool initialized")
end)