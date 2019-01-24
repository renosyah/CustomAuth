package com.example.renosyahputra.grpcservicelib.model;

import java.io.Serializable;

public class CallbackDataModel implements Serializable {
    public String IdCallback = "";
    public UserDataModel userDataModel = new UserDataModel();

    public CallbackDataModel() {
        super();
    }

    public CallbackDataModel(String idCallback, UserDataModel userDataModel) {
        IdCallback = idCallback;
        this.userDataModel = userDataModel;
    }
}
