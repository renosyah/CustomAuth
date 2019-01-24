package com.example.renosyahputra.grpcservicelib.ModelCasting;

import auth.Auth;
import com.example.renosyahputra.grpcservicelib.model.CallbackDataModel;

public class CallbackDataModelCasting {
    public static CallbackDataModel toCallbackDataModel(Auth.callbackData callbackData){
        return new CallbackDataModel(callbackData.getIdCallback()
                ,UserDataModelCasting.toUserData(callbackData.getUser())
        );
    }

    public static Auth.callbackData toCallbackDataModelGRPC(CallbackDataModel callbackDataModel){
        return Auth.callbackData.newBuilder()
                .setIdCallback(callbackDataModel.IdCallback)
                .setUser(UserDataModelCasting.toUserDataGRPC(callbackDataModel.userDataModel))
                .build();
    }
}
