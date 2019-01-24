package com.example.renosyahputra.grpcservicelib.activityAuth;

import android.annotation.SuppressLint;
import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.v7.app.AppCompatActivity;
import android.view.KeyEvent;
import android.view.View;
import android.webkit.WebView;
import android.webkit.WebViewClient;
import com.example.renosyahputra.grpcservicelib.R;
import com.example.renosyahputra.grpcservicelib.auth_grpc.WaitingForCallback;
import com.example.renosyahputra.grpcservicelib.model.UserDataModel;
import com.example.renosyahputra.grpcservicelib.util.RandomId;

import java.util.ArrayList;

import static com.example.renosyahputra.grpcservicelib.util.StaticVariable.Port_Rest;
import static com.example.renosyahputra.grpcservicelib.util.StaticVariable.Url;

public class AuthActivity extends AppCompatActivity implements WaitingForCallback.OnWaitingForCallbackListener {

    Context context;
    WebView webViewAuth;
    WaitingForCallback waitingForCallback = WaitingForCallback.newInstance();
    String id_callback = RandomId.randomAlphaNumeric(10);

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_auth);
        initiationWidget();
    }

    @SuppressLint("SetJavaScriptEnabled")
    private void initiationWidget(){
        context = this;

        webViewAuth = findViewById(R.id.webViewAuth);
        webViewAuth.getSettings().setLoadsImagesAutomatically(true);
        webViewAuth.getSettings().setJavaScriptEnabled(true);
        webViewAuth.getSettings().setDomStorageEnabled(true);


        webViewAuth.getSettings().setSupportZoom(true);
        webViewAuth.getSettings().setBuiltInZoomControls(true);
        webViewAuth.getSettings().setDisplayZoomControls(false);

        webViewAuth.setScrollBarStyle(View.SCROLLBARS_INSIDE_OVERLAY);
        webViewAuth.setWebViewClient(new WebViewClient());
        webViewAuth.loadUrl("http://"+Url+":"+Port_Rest+"/?id_callback="+id_callback);

        waitingForCallback = WaitingForCallback.newInstance();
        waitingForCallback.setIdCallback(id_callback)
                .setOnWaitingForCallbackListener(this)
                .listen();
    }

    @Override
    public boolean onKeyDown(int keyCode, KeyEvent event) {
        if ((keyCode == KeyEvent.KEYCODE_BACK)) {
            onBackPressed();
        }
        return super.onKeyDown(keyCode, event);
    }

    @Override
    public void onBackPressed() {
        waitingForCallback.shutDown();
        super.onBackPressed();
    }

    @Override
    public void onSuccesGetCallback(@NonNull UserDataModel userDataModel) {
        Intent returnIntent = new Intent();
        returnIntent.putExtra("data",userDataModel);
        returnIntent.putExtra("errors",new ArrayList<String>());
        setResult(Activity.RESULT_OK,returnIntent);
        finish();
    }

    @Override
    public void onFailGetCallback(@NonNull ArrayList<String> errors) {
        Intent returnIntent = new Intent();
        returnIntent.putExtra("data",new UserDataModel());
        returnIntent.putExtra("errors",errors);
        setResult(Activity.RESULT_OK,returnIntent);
        finish();
    }
}
