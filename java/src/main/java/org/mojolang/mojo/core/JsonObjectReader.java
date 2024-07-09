package org.mojolang.mojo.core;

import java.lang.reflect.Type;
import java.lang.reflect.ParameterizedType;

import com.alibaba.fastjson2.reader.ObjectReader;
import com.alibaba.fastjson2.JSONReader;

public class JsonObjectReader<T> implements ObjectReader<T> {
    //public static final JsonObjectReader<T> INSTANCE = new JsonObjectReader<>();

    @Override
    public T readObject(JSONReader jsonReader, Type fieldType, java.lang.Object fieldName, long features) {
        if (jsonReader.nextIfNull()) {
            return null;
        }

        final String value = jsonReader.readJSONObject().toString();
        // final String value = parser.parse(fieldName).toString();

//        if (fieldType instanceof Class && Message.class.isAssignableFrom((Class<?>) fieldType)) {
//            return (T) ProtobufUtils.toBean(value, (Class) fieldType);
//        }
//
//        if (fieldType instanceof ParameterizedType) {
//            final ParameterizedType type = (ParameterizedType) fieldType;
//
//            if (List.class.isAssignableFrom((Class<?>) type.getRawType())) {
//                final Type argument = type.getActualTypeArguments()[0];
//                if (Message.class.isAssignableFrom((Class<?>) argument)) {
//                    return (T) ProtobufUtils.toBeanList(value, (Class) argument);
//                }
//            }
//
//            if (Map.class.isAssignableFrom((Class<?>) type.getRawType())) {
//                final Type[] arguments = type.getActualTypeArguments();
//                if (arguments.length == 2) {
//                    final Type keyType = arguments[0], valueType = arguments[1];
//                    if (Message.class.isAssignableFrom((Class<?>) valueType)) {
//                        return (T) ProtobufUtils.toBeanMap(value, (Class) keyType, (Class) valueType);
//                    }
//                }
//            }
//        }

        return null;
    }
}
