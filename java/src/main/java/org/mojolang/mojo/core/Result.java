package org.mojolang.mojo.core;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import lombok.Data;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Result<T> {
    private Error error;
    private T data;

    public Result<T> setError(Error error) {
        this.error = error;
        return this;
    }
    public Result<T> setError(ErrorCode code) {
        this.error = Error.newBuilder().setCode(code).build();
        return this;
    }
    public Result<T> setError(Integer code) {
        this.error = Error.newBuilder().setCode(ErrorCodes.build(code)).build();
        return this;
    }
    public Result<T> setError(Integer code, String msg) {
        this.error = Error.newBuilder().setCode(ErrorCodes.build(code)).setMessage(msg).build();
        return this;
    }

    public Result<T> setData(T data) {
        this.data = data;
        return this;
    }

    public int getCode() {
        if (this.error != null) {
            return this.error.getCode().getCode();
        }
        return 0;
    }

    public String getMessage() {
        if (this.error != null) {
            return this.error.getMessage();
        }
        return "";
    }

    public static <T> Result<T> success() {
        return build(ErrorCodes.SUCCESS,null);
    }

    public static <T> Result<T> success(T data) {
        return build(ErrorCodes.SUCCESS, data);
    }

    public static <T> Result<T> fail() {
        return build(ErrorCodes.BAD_REQUEST, null);
    }

    public static <T> Result<T> fail(T data) {
        return build(ErrorCodes.BAD_REQUEST, data);
    }

    public static <T> Result<T> fail(int code) {
        return build(code, null);
    }

    public static <T> Result<T> fail(int code, T data) {
        return build(code, data);
    }

    public static <T> Result<T> fail(int code, String msg) {
        return build(code, msg, null);
    }

    public static <T> Result<T> fail(int code, String msg, T data) {
        return build(code, msg, data);
    }

    public static <T> Result<T> fail(ErrorException exception) {
        return fail(exception, null);
    }

    public static <T> Result<T> fail(ErrorException exception, T data) {
        return new Result<T>(exception.toError(), data);
    }

    public static <T> Result<T> build(ErrorCode code, T data) {
        return new Result<T>().setError(code).setData(data);
    }

    public static <T> Result<T> build(Integer code, T data) {
        return new Result<T>().setError(code).setData(data);
    }

    public static <T> Result<T> build(Integer code, String msg, T data) {
        return new Result<T>().setError(code, msg).setData(data);
    }
}