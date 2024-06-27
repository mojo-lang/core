package org.mojolang.mojo.core;

public class Exception extends RuntimeException{
    private ErrorCode errorCode;

    public Error toError() {
        return Error.newBuilder()
                .setCode(errorCode)
                .build();
    }

    public Exception(ErrorCode errorCode) {
        super(errorCode.getDescription());
    }

    public Exception(ErrorCode errorCode, String message) {
        super(message);
        this.errorCode = errorCode;
    }

    public int getCode() {
        return errorCode.getCode();
    }

    public int getHttpStatusCode() {
        return errorCode.getHttpStatusCode();
    }
}
