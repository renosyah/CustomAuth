package com.example.renosyahputra.grpcservicelib.authFragment;

import android.app.Activity;
import android.app.DialogFragment;
import android.content.Context;
import android.content.Intent;
import android.graphics.Color;
import android.graphics.drawable.ColorDrawable;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.view.Window;
import com.example.renosyahputra.grpcservicelib.R;
import com.example.renosyahputra.grpcservicelib.activityAuth.AuthActivity;
import com.example.renosyahputra.grpcservicelib.dialogAuth.DialogAuth;
import com.example.renosyahputra.grpcservicelib.model.UserDataModel;

import java.util.ArrayList;

public class AuthFragment extends DialogFragment {
    private View v;
    private Context ctx;
    private int RequestCode = 0;
    private DialogAuth.OnAuthListener listener;

    public static AuthFragment getInstance(){
        return  new AuthFragment();
    }

    public void setOnAuthListener(DialogAuth.OnAuthListener listener) {
        this.listener = listener;
    }

    @Nullable
    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, Bundle savedInstanceState) {
        v = inflater.inflate(R.layout.auth_fragment,container,false);

        ctx = getActivity();

        getDialog().getWindow().setBackgroundDrawable(new ColorDrawable(Color.TRANSPARENT));
        getDialog().requestWindowFeature(Window.FEATURE_NO_TITLE);

        RequestCode = 135;

        startActivityForResult(new Intent(ctx, AuthActivity.class), RequestCode);

        return v;
    }

    @Override
    public void onActivityResult(int requestCode, int resultCode, Intent data) {
        if (requestCode == RequestCode && Activity.RESULT_OK == resultCode && listener != null && data != null){

            UserDataModel userDataModel = (UserDataModel) data.getSerializableExtra("data");
            listener.onAuthSucces(userDataModel);

            ArrayList<String> errors = data.getStringArrayListExtra("errors");
            listener.onAuthFail(errors);

        }
        getDialog().dismiss();
    }
}
