package com.example.renosyahputra.grpcservicelib.dialogAuth;

import android.annotation.SuppressLint;
import android.app.Activity;
import android.content.Context;
import android.support.annotation.NonNull;
import com.example.renosyahputra.grpcservicelib.authFragment.AuthFragment;
import com.example.renosyahputra.grpcservicelib.model.UserDataModel;

import java.util.ArrayList;

public class DialogAuth {

    private static DialogAuth _instance;
    private Context context;
    private OnAuthListener listener;

    public static DialogAuth newInstanceWithContext(Context ctx){
        _instance = new DialogAuth();
        _instance.context = ctx;
        return _instance;
    }
    public DialogAuth setOnAuthListener(OnAuthListener listener){
        _instance.listener = listener;
        return _instance;
    }

    private DialogAuth() {

    }

    @SuppressLint({"SetJavaScriptEnabled", "ClickableViewAccessibility"})
    public void show(){

        AuthFragment authFragment = AuthFragment.getInstance();
        authFragment.setOnAuthListener(listener);
        authFragment.show(((Activity)context).getFragmentManager(),"mydialog");

    }

    public interface OnAuthListener {
        void onAuthSucces(@NonNull UserDataModel user);
        void onAuthFail(@NonNull ArrayList<String> Errors);
    }
}
