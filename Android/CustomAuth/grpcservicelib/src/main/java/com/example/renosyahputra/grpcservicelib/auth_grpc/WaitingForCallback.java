package com.example.renosyahputra.grpcservicelib.auth_grpc;

import android.os.AsyncTask;
import android.support.annotation.NonNull;
import auth.Auth;
import auth.authServiceGrpc;
import com.example.renosyahputra.grpcservicelib.ModelCasting.CallbackDataModelCasting;
import com.example.renosyahputra.grpcservicelib.model.CallbackDataModel;
import com.example.renosyahputra.grpcservicelib.model.UserDataModel;
import com.example.renosyahputra.grpcservicelib.util.RandomId;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.Metadata;
import io.grpc.stub.MetadataUtils;
import io.grpc.stub.StreamObserver;

import java.util.ArrayList;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

import static com.example.renosyahputra.grpcservicelib.util.StaticVariable.*;

public class WaitingForCallback extends AsyncTask<Void,Void, CallbackDataModel> {

    private static WaitingForCallback _instance;
    private String idCallback;
    private OnWaitingForCallbackListener listener;
    private Metadata header = new Metadata();
    private ManagedChannel channel;
    private ArrayList<String> Errors = new ArrayList<>();

    CallbackDataModel callbackDataModel = new CallbackDataModel();

    public static WaitingForCallback newInstance(){
        _instance = new WaitingForCallback();
        return _instance;
    }

    public WaitingForCallback setIdCallback(String idCallback) {
        _instance.idCallback = idCallback;
        return _instance;
    }

    public WaitingForCallback setOnWaitingForCallbackListener(OnWaitingForCallbackListener listener){
        _instance.listener = listener;
        return _instance;
    }

    public void listen(){
        if (_instance != null){
            _instance.execute();
        }
    }
    public void shutDown(){
        if (_instance != null){
            _instance.cancel(true);
        }
    }


    private WaitingForCallback() {
    }


    @Override
    protected void onPreExecute() {
        super.onPreExecute();

        Metadata.Key<String> key = Metadata.Key.of(Authorization, Metadata.ASCII_STRING_MARSHALLER);
        _instance.header.put(key, RandomId.randomAlphaNumeric(10).toString());

        _instance.channel = ManagedChannelBuilder
                .forAddress(Url, Port_Grpc)
                .usePlaintext(true)
                .build();
    }

    private int timeOut = 30;
    private StreamObserver<Auth.callbackData> responseObserver;
    private StreamObserver<Auth.callbackData> requestObserver;
    private Throwable failed = null;
    private CountDownLatch finishLatch = new CountDownLatch(1);

    @Override
    protected CallbackDataModel doInBackground(Void... voids) {

        try {

            authServiceGrpc.authServiceStub stub = authServiceGrpc.newStub(_instance.channel);
            stub = MetadataUtils.attachHeaders(stub,_instance.header);

            _instance.responseObserver = new StreamObserver<Auth.callbackData>() {
                @Override
                public void onNext(Auth.callbackData value) {

                    if (value.getIdCallback().equals(_instance.idCallback)){
                        _instance.callbackDataModel = CallbackDataModelCasting.toCallbackDataModel(value);
                        _instance.requestObserver.onCompleted();
                    }
                }

                @Override
                public void onError(Throwable t) {
                    _instance.failed = t;
                    _instance.finishLatch.countDown();
                }

                @Override
                public void onCompleted() {
                    _instance.finishLatch.countDown();
                }
            };

            _instance.requestObserver = stub.waitCallback(_instance.responseObserver);

            _instance.requestObserver.onNext(CallbackDataModelCasting
                    .toCallbackDataModelGRPC(_instance.callbackDataModel)
            );

            if (!_instance.finishLatch.await(_instance.timeOut , TimeUnit.MINUTES)) {

                throw new RuntimeException(
                        "Could not finish rpc within "+timeOut+" minute, the server is likely down");
            }

            if (_instance.failed != null) {
                throw new RuntimeException(_instance.failed);
            }

        }catch (Exception e){
            _instance.Errors.add(e.getMessage());
        }

        return _instance.callbackDataModel;
    }


    @Override
    protected void onCancelled() {
        super.onCancelled();

        try {
            _instance.channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            _instance.Errors.add(e.getMessage());
        }

        if (_instance.Errors.size() > 0 && _instance.listener != null){
            _instance.listener.onFailGetCallback(_instance.Errors);
        }

    }

    @Override
    protected void onPostExecute(CallbackDataModel callbackData) {
        super.onPostExecute(callbackData);

        try {
            _instance.channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            _instance.Errors.add(e.getMessage());
        }

        if (_instance.Errors.size() > 0 && _instance.listener != null){
            _instance.listener.onFailGetCallback(_instance.Errors);
        }

        if (_instance.listener != null){
            _instance.listener.onSuccesGetCallback(callbackData.userDataModel);
        }

    }

    public interface OnWaitingForCallbackListener {
        void onSuccesGetCallback(@NonNull UserDataModel userDataModel);
        void onFailGetCallback(@NonNull ArrayList<String> errors);
    }
}
