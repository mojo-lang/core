package org.mojolang.mojo.core;

public class Exception extends RuntimeException{
    private ErrorCode code;

    public Error toError() {
        return Error.newBuilder()
                .setCode(code)
                .build();
    }

    public Exception(ErrorCode code) {
        super(code.getDescription());
    }

    public Exception(ErrorCode code, String message) {
        super(message);
        this.code = code;
    }
}
