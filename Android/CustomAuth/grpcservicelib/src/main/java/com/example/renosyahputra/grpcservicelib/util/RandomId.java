package com.example.renosyahputra.grpcservicelib.util;

public class RandomId {
    private static final String ALPHA_NUMERIC_STRING = "abOPQRcdklmnopefBCDEFstghijuvwzAKLGHIJqrMNSTWXYZ01234UV56789";
    public static String randomAlphaNumeric(int count) {
        StringBuilder builder = new StringBuilder();
        while (count-- != 0) {
            int character = (int)(Math.random()*ALPHA_NUMERIC_STRING.length());
            builder.append(ALPHA_NUMERIC_STRING.charAt(character));
        }
        return builder.toString();
    }
}
