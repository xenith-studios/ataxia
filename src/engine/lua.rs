use rlua::{Lua, Function};

/// Run a sample command to show how to use rlua
pub fn test() {
    let lua = Lua::new();
    lua.context(|context| {
        let args: Vec<String> = vec!["test".to_string(), "123".to_string()];
        context.load(r#"
            function(uid, args)
                local res = "UID: "..uid.."\nArgs: "
                for i,v in ipairs(args) do
                res = res..v.." "
                end
                print(res)
            end
        "#).eval::<Function>()?.call::<_, ()>(("1", args))
    }).unwrap();
}
