package org.mojolang.mojo.core;

import lombok.AllArgsConstructor;
import lombok.Data;
import java.util.List;

@Data
@AllArgsConstructor
public class Pagination<T> {
    private Exception exception;
    private int totalCount;
    private String nextPageToken;
    private List<T> data;

    public void setException(ErrorCode code) {
        this.exception = new Exception(code);
    }
    public void setException(Integer code) {
        this.exception = new Exception(ErrorCodes.build(code));
    }
    public void setException(Integer code, String msg) {
        this.exception = new Exception(ErrorCodes.build(code), msg);
    }

    public static <T> Pagination<T> fail() {
        return build(ErrorCodes.BAD_REQUEST, null, 0, "");
    }

    public static <T> Pagination<T> fail(int code) {
        return build(code, null, 0, "");
    }

    public static <T> Pagination<T> fail(int code, String msg) {
        return build(code, msg, null, 0, "");
    }

    public static <T> Pagination<T> fail(Exception exception) {
        return build(exception, null, 0, "");
    }

    public static <T> Pagination<T> build(ErrorCode code, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<>(new Exception(code), totalCount, nextPageToken, data);
    }

    public static <T> Pagination<T> build(Integer code, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<>(new Exception(ErrorCodes.build(code)), totalCount, nextPageToken, data);
    }

    public static <T> Pagination<T> build(Integer code, String msg, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<>(new Exception(ErrorCodes.build(code), msg), totalCount, nextPageToken, data);
    }

    public static <T> Pagination<T> build(Exception exception, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<>(exception, totalCount, nextPageToken, data);
    }
}