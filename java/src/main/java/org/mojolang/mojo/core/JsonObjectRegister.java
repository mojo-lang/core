package org.mojolang.mojo.core;

import com.alibaba.fastjson2.JSON;
//@Component
public class JsonObjectRegister {
    //@PostConstruct
    public static void init() {
        JSON.register(org.mojolang.mojo.core.Object.class, JsonObjectWriter.INSTANCE);
        // JSON.register(org.mojolang.mojo.core.Object.class, JsonObjectReader.INSTANCE);
    }
}
