package org.mojolang.mojo.core;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class Result<T> {
    private Exception exception;
    private T data;

    public static <T> Result<T> success() {
        return build(null, ErrorCodes.SUCCESS);
    }

    public static <T> Result<T> success(T data) {
        return build(data, ErrorCodes.SUCCESS);
    }

    public static <T> Result<T> fail() {
        return build(null, ErrorCodes.BAD_REQUEST);
    }

    public static <T> Result<T> fail(T data) {
        return build(data, ErrorCodes.BAD_REQUEST);
    }

    public static <T> Result<T> fail(int code) {
        return build(null, code);
    }

    public static <T> Result<T> fail(int code, T data) {
        return build(data, code);
    }

    public static <T> Result<T> fail(int code, String msg) {
        return build(null, code, msg);
    }

    public static <T> Result<T> fail(int code, String msg, T data) {
        return build(data, code, msg);
    }

    public static <T> Result<T> fail(Exception exception) {
        return fail(exception, null);
    }

    public static <T> Result<T> fail(Exception exception, T data) {
        return new Result<>(exception, data);
    }

    public static <T> Result<T> build(T data, ErrorCode code) {
        return new Result<>(new Exception(code), data);
    }

    public static <T> Result<T> build(T data, Integer code) {
        return new Result<>(new Exception(ErrorCodes.build(code)), data);
    }

    public static <T> Result<T> build(T data, Integer code, String msg) {
        return new Result<>(new Exception(ErrorCodes.build(code), msg), data);
    }
}