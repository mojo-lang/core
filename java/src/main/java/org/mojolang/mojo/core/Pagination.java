package org.mojolang.mojo.core;

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Pagination<T> {
    private Error error;
    private int totalCount;
    private String nextPageToken;
    private List<T> data;

    public Pagination<T> setError(Error error) {
        this.error = error;
        return this;
    }
    public Pagination<T> setError(ErrorCode code) {
        this.error = Error.newBuilder().setCode(code).build();
        return this;
    }
    public Pagination<T> setError(Integer code) {
        this.error = Error.newBuilder().setCode(ErrorCodes.build(code)).build();
        return this;
    }
    public Pagination<T> setError(Integer code, String msg) {
        this.error = Error.newBuilder().setCode(ErrorCodes.build(code)).setMessage(msg).build();
        return this;
    }

    public Pagination<T> setTotalCount(int count) {
        this.totalCount = count;
        return this;
    }

    public Pagination<T> setNextPageToken(String token) {
        this.nextPageToken = token;
        return this;
    }

    public Pagination<T> setData(List<T> data) {
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

    public static <T> Pagination<T> fail() {
        return build(ErrorCodes.BAD_REQUEST, null, 0, "");
    }

    public static <T> Pagination<T> fail(int code) {
        return build(code, null, 0, "");
    }

    public static <T> Pagination<T> fail(int code, String msg) {
        return build(code, msg, null, 0, "");
    }

    public static <T> Pagination<T> fail(ErrorException exception) {
        return build(exception, null, 0, "");
    }

    public static <T> Pagination<T> build(ErrorCode code, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<T>().setError(code).setData(data).setTotalCount(totalCount).setNextPageToken(nextPageToken);
    }

    public static <T> Pagination<T> build(Integer code, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<T>().setError (code).setData(data).setTotalCount(totalCount).setNextPageToken(nextPageToken);
    }

    public static <T> Pagination<T> build(Integer code, String msg, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<T>().setError(code, msg).setData(data).setTotalCount(totalCount).setNextPageToken(nextPageToken);
    }

    public static <T> Pagination<T> build(ErrorException exception, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<T>().setError(exception.toError()).setData(data).setTotalCount(totalCount).setNextPageToken(nextPageToken);
    }
}